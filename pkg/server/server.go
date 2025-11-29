package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"backend/pkg/bindings"
)

type Config struct {
	RPCURL     string
	Merchant   string
	Contract   string
	PrivateKey string
	Token      string
	ListenAddr string
}

type IntentResponse struct {
	Intent      ClientIntent `json:"intent"`
	MessageHash string       `json:"messageHash"` // 对此 hash 做 eth_sign/personal_sign
}

type ClientIntent struct {
	Payer    string `json:"payer"`
	Merchant string `json:"merchant"`
	Token    string `json:"token"`
	Amount   string `json:"amount"`
	Nonce    string `json:"nonce"`
	Deadline int64  `json:"deadline"`
}

type PayRequest struct {
	Intent    ClientIntent `json:"intent"`
	Signature string       `json:"signature"` // 0x + 65 字节，payer 对 messageHash 的签名
}

type ApproveRequest struct {
	Token   string `json:"token"`
	Spender string `json:"spender"`
	Amount  string `json:"amount"`
}

type TransferRequest struct {
	Token  string `json:"token"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type Server struct {
	router     *gin.Engine
	client     *ethclient.Client
	contract   *bindings.X402Minimal
	merchant   common.Address
	token      common.Address
	chainID    *big.Int
	relayerKey *ecdsa.PrivateKey
}

func LoadConfigFromEnv() Config {
	cfg := Config{
		RPCURL:     mustEnv("RPC_URL"),
		Merchant:   mustEnv("MERCHANT_ADDRESS"),
		Contract:   mustEnv("CONTRACT_ADDRESS"),
		PrivateKey: mustEnv("PRIVATE_KEY"), // relayer Key，用于代付 gas
		Token:      os.Getenv("TOKEN_ADDRESS"),
		ListenAddr: os.Getenv("LISTEN_ADDR"),
	}
	if cfg.Token == "" {
		cfg.Token = cfg.Merchant // 演示用：缺省时回退为商户地址占位
	}
	if cfg.ListenAddr == "" {
		cfg.ListenAddr = ":8080"
	} else if !strings.Contains(cfg.ListenAddr, ":") {
		cfg.ListenAddr = ":" + cfg.ListenAddr
	}
	return cfg
}

func New(cfg Config) (*Server, error) {
	client, err := ethClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	relayerKey, err := crypto.HexToECDSA(trim0x(cfg.PrivateKey))
	if err != nil {
		return nil, err
	}

	contract, err := bindings.NewX402Minimal(common.HexToAddress(cfg.Contract), client)
	if err != nil {
		return nil, err
	}

	s := &Server{
		router:     gin.Default(),
		client:     client,
		contract:   contract,
		merchant:   common.HexToAddress(cfg.Merchant),
		token:      common.HexToAddress(cfg.Token),
		chainID:    chainID,
		relayerKey: relayerKey,
	}

	s.registerRoutes()
	return s, nil
}

func (s *Server) Run(addr string) error {
	r := s.router

	// 访问路径里 /index/xxx 映射到本地 ./index 目录
	r.Static("/index", "./cmd/server/index")
	r.GET("/", func(c *gin.Context) {
		c.File("./cmd/server/index/merchant.html")
	})

	return r.Run(addr)
}

func (s *Server) registerRoutes() {
	s.router.GET("/api/intent", s.handleIntent)
	s.router.GET("/api/balance", s.handleBalance)
	s.router.POST("/api/approve", s.handleApprove)
	s.router.POST("/api/transfer", s.handleTransfer)
	s.router.POST("/api/pay", s.handlePay)
}

// 生成一次性的支付意图，返回 messageHash 供前端签名
func (s *Server) handleIntent(c *gin.Context) {
	payer := c.DefaultQuery("payer", s.merchant.Hex())
	amount := c.DefaultQuery("amount", "1000000000000000000") // default 1 token (18 decimals)
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	deadline := time.Now().Add(5 * time.Minute).Unix()

	intent := ClientIntent{
		Payer:    payer,
		Merchant: s.merchant.Hex(),
		Token:    s.token.Hex(),
		Amount:   amount,
		Nonce:    nonce,
		Deadline: deadline,
	}
	msgHash := calcMessageHash(intent)
	c.JSON(http.StatusOK, IntentResponse{
		Intent:      intent,
		MessageHash: "0x" + hex.EncodeToString(msgHash[:]),
	})
}

// 前端携带签名来请求代付，服务端作为 relayer 发起 pay()
func (s *Server) handlePay(c *gin.Context) {
	var req PayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	intent := toBindingIntent(req.Intent)
	sig, err := hex.DecodeString(trim0x(req.Signature))
	if err != nil || len(sig) != 65 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	auth, err := s.newTransactor(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx, err := s.contract.Pay(auth, intent, sig)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex()})
}

// 查询账户在指定 ERC20 上的余额（默认为配置中的 token）
func (s *Server) handleBalance(c *gin.Context) {
	addressHex := c.Query("address")
	if !common.IsHexAddress(addressHex) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid address"})
		return
	}
	tokenHex := c.DefaultQuery("token", s.token.Hex())
	if !common.IsHexAddress(tokenHex) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		return
	}

	account := common.HexToAddress(addressHex)
	token := common.HexToAddress(tokenHex)

	balance, err := s.queryERC20Balance(c.Request.Context(), token, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": addressHex,
		"token":   tokenHex,
		"balance": balance.String(),
	})
}

// 代为调用 ERC20 approve
func (s *Server) handleApprove(c *gin.Context) {
	var req ApproveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !common.IsHexAddress(req.Token) || !common.IsHexAddress(req.Spender) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token or spender"})
		return
	}
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
		return
	}

	contract, err := s.boundERC20(common.HexToAddress(req.Token))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth, err := s.newTransactor(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx, err := contract.Transact(auth, "approve", common.HexToAddress(req.Spender), amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex()})
}

// 代为调用 ERC20 transfer
func (s *Server) handleTransfer(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !common.IsHexAddress(req.Token) || !common.IsHexAddress(req.To) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token or recipient"})
		return
	}
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount"})
		return
	}

	contract, err := s.boundERC20(common.HexToAddress(req.Token))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth, err := s.newTransactor(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx, err := contract.Transact(auth, "transfer", common.HexToAddress(req.To), amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"txHash": tx.Hash().Hex()})
}

func (s *Server) newTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(s.relayerKey, s.chainID)
	if err != nil {
		return nil, err
	}
	auth.Context = ctx
	auth.GasLimit = 800000
	return auth, nil
}

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing env %s", key)
	}
	return val
}

func ethClient(url string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return ethclient.DialContext(ctx, url)
}

func toBindingIntent(ci ClientIntent) bindings.X402MinimalPaymentIntent {
	amount := new(big.Int)
	if _, ok := amount.SetString(ci.Amount, 10); !ok {
		amount = big.NewInt(0)
	}
	nonce := new(big.Int)
	if _, ok := nonce.SetString(ci.Nonce, 10); !ok {
		nonce = big.NewInt(0)
	}
	return bindings.X402MinimalPaymentIntent{
		Payer:    common.HexToAddress(ci.Payer),
		Merchant: common.HexToAddress(ci.Merchant),
		Token:    common.HexToAddress(ci.Token),
		Amount:   amount,
		Nonce:    nonce,
		Deadline: big.NewInt(ci.Deadline),
	}
}

func calcMessageHash(ci ClientIntent) [32]byte {
	intent := toBindingIntent(ci)
	inner := crypto.Keccak256Hash(
		paddedAddress(intent.Payer),
		paddedAddress(intent.Merchant),
		paddedAddress(intent.Token),
		paddedBig(intent.Amount),
		paddedBig(intent.Nonce),
		paddedBig(intent.Deadline),
	)
	prefix := []byte("\x19Ethereum Signed Message:\n32")
	return crypto.Keccak256Hash(append(prefix, inner.Bytes()...))
}

func paddedAddress(addr common.Address) []byte {
	return common.LeftPadBytes(addr.Bytes(), 32)
}

func paddedBig(v *big.Int) []byte {
	return common.LeftPadBytes(v.Bytes(), 32)
}

func trim0x(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}

func (s *Server) queryERC20Balance(ctx context.Context, token, account common.Address) (*big.Int, error) {
	contract, err := s.boundERC20(token)
	if err != nil {
		return nil, err
	}
	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "balanceOf", account); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, errors.New("empty balance result")
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

func (s *Server) boundERC20(token common.Address) (*bind.BoundContract, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(token, parsedABI, s.client, s.client, s.client), nil
}

const erc20ABI = `[{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}]`
