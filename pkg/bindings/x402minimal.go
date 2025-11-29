// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// X402MinimalPaymentIntent is an auto generated low-level Go binding around an user-defined struct.
type X402MinimalPaymentIntent struct {
	Payer    common.Address
	Merchant common.Address
	Token    common.Address
	Amount   *big.Int
	Nonce    *big.Int
	Deadline *big.Int
}

// X402MinimalMetaData contains all meta data concerning the X402Minimal contract.
var X402MinimalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structX402Minimal.PaymentIntent\",\"name\":\"intent\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610f4e8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80631647795e1461003857806365530d0d14610068575b5f5ffd5b610052600480360381019061004d9190610801565b610084565b60405161005f9190610859565b60405180910390f35b610082600480360381019061007d91906108f5565b6100ad565b005b5f602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b8260a001354211156100f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100eb906109ac565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1683602001602081019061011e91906109ca565b73ffffffffffffffffffffffffffffffffffffffff161415801561017f57505f73ffffffffffffffffffffffffffffffffffffffff16835f01602081019061016691906109ca565b73ffffffffffffffffffffffffffffffffffffffff1614155b6101be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b590610a3f565b60405180910390fd5b5f5f845f0160208101906101d291906109ca565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f846080013581526020019081526020015f205f9054906101000a900460ff161561026b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026290610aa7565b60405180910390fd5b5f61027d61027885610546565b6105c0565b90505f61028b8285856105ef565b9050845f01602081019061029f91906109ca565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461030c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030390610b0f565b60405180910390fd5b60015f5f875f01602081019061032291906109ca565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f876080013581526020019081526020015f205f6101000a81548160ff0219169083151502179055505f85604001602081019061039891906109ca565b73ffffffffffffffffffffffffffffffffffffffff166323b872dd875f0160208101906103c591906109ca565b8860200160208101906103d891906109ca565b89606001356040518463ffffffff1660e01b81526004016103fb93929190610b4b565b6020604051808303815f875af1158015610417573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061043b9190610baa565b90508061047d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047490610c1f565b60405180910390fd5b85604001602081019061049091906109ca565b73ffffffffffffffffffffffffffffffffffffffff168660200160208101906104b991906109ca565b73ffffffffffffffffffffffffffffffffffffffff16875f0160208101906104e191906109ca565b73ffffffffffffffffffffffffffffffffffffffff167fd32d5970dffb7a78dddc9bf38e51bd1ff7ae0ab4e0d3fe026dfe5c1614b57df789606001358a608001358b60a0013560405161053693929190610c3d565b60405180910390a4505050505050565b5f815f01602081019061055991906109ca565b82602001602081019061056c91906109ca565b83604001602081019061057f91906109ca565b846060013585608001358660a001356040516020016105a396959493929190610c72565b604051602081830303815290604052805190602001209050919050565b5f816040516020016105d29190610d4e565b604051602081830303815290604052805190602001209050919050565b5f60418383905014610636576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062d90610dbd565b60405180910390fd5b5f5f5f853592506020860135915060408601355f1a9050601b8160ff1614806106625750601c8160ff16145b6106a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069890610e25565b60405180910390fd5b5f6001888386866040515f81526020016040526040516106c49493929190610e6d565b6020604051602081039080840390855afa1580156106e4573d5f5f3e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361075e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075590610efa565b60405180910390fd5b809450505050509392505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61079d82610774565b9050919050565b6107ad81610793565b81146107b7575f5ffd5b50565b5f813590506107c8816107a4565b92915050565b5f819050919050565b6107e0816107ce565b81146107ea575f5ffd5b50565b5f813590506107fb816107d7565b92915050565b5f5f604083850312156108175761081661076c565b5b5f610824858286016107ba565b9250506020610835858286016107ed565b9150509250929050565b5f8115159050919050565b6108538161083f565b82525050565b5f60208201905061086c5f83018461084a565b92915050565b5f5ffd5b5f60c0828403121561088b5761088a610872565b5b81905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126108b5576108b4610894565b5b8235905067ffffffffffffffff8111156108d2576108d1610898565b5b6020830191508360018202830111156108ee576108ed61089c565b5b9250929050565b5f5f5f60e0848603121561090c5761090b61076c565b5b5f61091986828701610876565b93505060c084013567ffffffffffffffff81111561093a57610939610770565b5b610946868287016108a0565b92509250509250925092565b5f82825260208201905092915050565b7f65787069726564000000000000000000000000000000000000000000000000005f82015250565b5f610996600783610952565b91506109a182610962565b602082019050919050565b5f6020820190508181035f8301526109c38161098a565b9050919050565b5f602082840312156109df576109de61076c565b5b5f6109ec848285016107ba565b91505092915050565b7f7a65726f000000000000000000000000000000000000000000000000000000005f82015250565b5f610a29600483610952565b9150610a34826109f5565b602082019050919050565b5f6020820190508181035f830152610a5681610a1d565b9050919050565b7f6e6f6e63652d75736564000000000000000000000000000000000000000000005f82015250565b5f610a91600a83610952565b9150610a9c82610a5d565b602082019050919050565b5f6020820190508181035f830152610abe81610a85565b9050919050565b7f6261642d736967000000000000000000000000000000000000000000000000005f82015250565b5f610af9600783610952565b9150610b0482610ac5565b602082019050919050565b5f6020820190508181035f830152610b2681610aed565b9050919050565b610b3681610793565b82525050565b610b45816107ce565b82525050565b5f606082019050610b5e5f830186610b2d565b610b6b6020830185610b2d565b610b786040830184610b3c565b949350505050565b610b898161083f565b8114610b93575f5ffd5b50565b5f81519050610ba481610b80565b92915050565b5f60208284031215610bbf57610bbe61076c565b5b5f610bcc84828501610b96565b91505092915050565b7f7472616e736665722d6661696c000000000000000000000000000000000000005f82015250565b5f610c09600d83610952565b9150610c1482610bd5565b602082019050919050565b5f6020820190508181035f830152610c3681610bfd565b9050919050565b5f606082019050610c505f830186610b3c565b610c5d6020830185610b3c565b610c6a6040830184610b3c565b949350505050565b5f60c082019050610c855f830189610b2d565b610c926020830188610b2d565b610c9f6040830187610b2d565b610cac6060830186610b3c565b610cb96080830185610b3c565b610cc660a0830184610b3c565b979650505050505050565b5f81905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f610d0f601c83610cd1565b9150610d1a82610cdb565b601c82019050919050565b5f819050919050565b5f819050919050565b610d48610d4382610d25565b610d2e565b82525050565b5f610d5882610d03565b9150610d648284610d37565b60208201915081905092915050565b7f7369672d6c656e000000000000000000000000000000000000000000000000005f82015250565b5f610da7600783610952565b9150610db282610d73565b602082019050919050565b5f6020820190508181035f830152610dd481610d9b565b9050919050565b7f6261642d760000000000000000000000000000000000000000000000000000005f82015250565b5f610e0f600583610952565b9150610e1a82610ddb565b602082019050919050565b5f6020820190508181035f830152610e3c81610e03565b9050919050565b610e4c81610d25565b82525050565b5f60ff82169050919050565b610e6781610e52565b82525050565b5f608082019050610e805f830187610e43565b610e8d6020830186610e5e565b610e9a6040830185610e43565b610ea76060830184610e43565b95945050505050565b7f65637265636f7665722d7a65726f0000000000000000000000000000000000005f82015250565b5f610ee4600e83610952565b9150610eef82610eb0565b602082019050919050565b5f6020820190508181035f830152610f1181610ed8565b905091905056fea264697066735822122037e34f67772b0b3cd3c2bb1bf6bb4b71b2f6c4526f4178006056c7c08f0f34ed64736f6c634300081e0033",
}

// X402MinimalABI is the input ABI used to generate the binding from.
// Deprecated: Use X402MinimalMetaData.ABI instead.
var X402MinimalABI = X402MinimalMetaData.ABI

// X402MinimalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use X402MinimalMetaData.Bin instead.
var X402MinimalBin = X402MinimalMetaData.Bin

// DeployX402Minimal deploys a new Ethereum contract, binding an instance of X402Minimal to it.
func DeployX402Minimal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *X402Minimal, error) {
	parsed, err := X402MinimalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(X402MinimalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &X402Minimal{X402MinimalCaller: X402MinimalCaller{contract: contract}, X402MinimalTransactor: X402MinimalTransactor{contract: contract}, X402MinimalFilterer: X402MinimalFilterer{contract: contract}}, nil
}

// X402Minimal is an auto generated Go binding around an Ethereum contract.
type X402Minimal struct {
	X402MinimalCaller     // Read-only binding to the contract
	X402MinimalTransactor // Write-only binding to the contract
	X402MinimalFilterer   // Log filterer for contract events
}

// X402MinimalCaller is an auto generated read-only Go binding around an Ethereum contract.
type X402MinimalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X402MinimalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type X402MinimalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X402MinimalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type X402MinimalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X402MinimalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type X402MinimalSession struct {
	Contract     *X402Minimal      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// X402MinimalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type X402MinimalCallerSession struct {
	Contract *X402MinimalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// X402MinimalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type X402MinimalTransactorSession struct {
	Contract     *X402MinimalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// X402MinimalRaw is an auto generated low-level Go binding around an Ethereum contract.
type X402MinimalRaw struct {
	Contract *X402Minimal // Generic contract binding to access the raw methods on
}

// X402MinimalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type X402MinimalCallerRaw struct {
	Contract *X402MinimalCaller // Generic read-only contract binding to access the raw methods on
}

// X402MinimalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type X402MinimalTransactorRaw struct {
	Contract *X402MinimalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewX402Minimal creates a new instance of X402Minimal, bound to a specific deployed contract.
func NewX402Minimal(address common.Address, backend bind.ContractBackend) (*X402Minimal, error) {
	contract, err := bindX402Minimal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &X402Minimal{X402MinimalCaller: X402MinimalCaller{contract: contract}, X402MinimalTransactor: X402MinimalTransactor{contract: contract}, X402MinimalFilterer: X402MinimalFilterer{contract: contract}}, nil
}

// NewX402MinimalCaller creates a new read-only instance of X402Minimal, bound to a specific deployed contract.
func NewX402MinimalCaller(address common.Address, caller bind.ContractCaller) (*X402MinimalCaller, error) {
	contract, err := bindX402Minimal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &X402MinimalCaller{contract: contract}, nil
}

// NewX402MinimalTransactor creates a new write-only instance of X402Minimal, bound to a specific deployed contract.
func NewX402MinimalTransactor(address common.Address, transactor bind.ContractTransactor) (*X402MinimalTransactor, error) {
	contract, err := bindX402Minimal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &X402MinimalTransactor{contract: contract}, nil
}

// NewX402MinimalFilterer creates a new log filterer instance of X402Minimal, bound to a specific deployed contract.
func NewX402MinimalFilterer(address common.Address, filterer bind.ContractFilterer) (*X402MinimalFilterer, error) {
	contract, err := bindX402Minimal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &X402MinimalFilterer{contract: contract}, nil
}

// bindX402Minimal binds a generic wrapper to an already deployed contract.
func bindX402Minimal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := X402MinimalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_X402Minimal *X402MinimalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _X402Minimal.Contract.X402MinimalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_X402Minimal *X402MinimalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X402Minimal.Contract.X402MinimalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_X402Minimal *X402MinimalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _X402Minimal.Contract.X402MinimalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_X402Minimal *X402MinimalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _X402Minimal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_X402Minimal *X402MinimalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X402Minimal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_X402Minimal *X402MinimalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _X402Minimal.Contract.contract.Transact(opts, method, params...)
}

// NonceUsed is a free data retrieval call binding the contract method 0x1647795e.
//
// Solidity: function nonceUsed(address , uint256 ) view returns(bool)
func (_X402Minimal *X402MinimalCaller) NonceUsed(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _X402Minimal.contract.Call(opts, &out, "nonceUsed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceUsed is a free data retrieval call binding the contract method 0x1647795e.
//
// Solidity: function nonceUsed(address , uint256 ) view returns(bool)
func (_X402Minimal *X402MinimalSession) NonceUsed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _X402Minimal.Contract.NonceUsed(&_X402Minimal.CallOpts, arg0, arg1)
}

// NonceUsed is a free data retrieval call binding the contract method 0x1647795e.
//
// Solidity: function nonceUsed(address , uint256 ) view returns(bool)
func (_X402Minimal *X402MinimalCallerSession) NonceUsed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _X402Minimal.Contract.NonceUsed(&_X402Minimal.CallOpts, arg0, arg1)
}

// Pay is a paid mutator transaction binding the contract method 0x65530d0d.
//
// Solidity: function pay((address,address,address,uint256,uint256,uint256) intent, bytes signature) returns()
func (_X402Minimal *X402MinimalTransactor) Pay(opts *bind.TransactOpts, intent X402MinimalPaymentIntent, signature []byte) (*types.Transaction, error) {
	return _X402Minimal.contract.Transact(opts, "pay", intent, signature)
}

// Pay is a paid mutator transaction binding the contract method 0x65530d0d.
//
// Solidity: function pay((address,address,address,uint256,uint256,uint256) intent, bytes signature) returns()
func (_X402Minimal *X402MinimalSession) Pay(intent X402MinimalPaymentIntent, signature []byte) (*types.Transaction, error) {
	return _X402Minimal.Contract.Pay(&_X402Minimal.TransactOpts, intent, signature)
}

// Pay is a paid mutator transaction binding the contract method 0x65530d0d.
//
// Solidity: function pay((address,address,address,uint256,uint256,uint256) intent, bytes signature) returns()
func (_X402Minimal *X402MinimalTransactorSession) Pay(intent X402MinimalPaymentIntent, signature []byte) (*types.Transaction, error) {
	return _X402Minimal.Contract.Pay(&_X402Minimal.TransactOpts, intent, signature)
}

// X402MinimalPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the X402Minimal contract.
type X402MinimalPaidIterator struct {
	Event *X402MinimalPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *X402MinimalPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X402MinimalPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(X402MinimalPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *X402MinimalPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X402MinimalPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X402MinimalPaid represents a Paid event raised by the X402Minimal contract.
type X402MinimalPaid struct {
	Payer    common.Address
	Merchant common.Address
	Token    common.Address
	Amount   *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0xd32d5970dffb7a78dddc9bf38e51bd1ff7ae0ab4e0d3fe026dfe5c1614b57df7.
//
// Solidity: event Paid(address indexed payer, address indexed merchant, address indexed token, uint256 amount, uint256 nonce, uint256 deadline)
func (_X402Minimal *X402MinimalFilterer) FilterPaid(opts *bind.FilterOpts, payer []common.Address, merchant []common.Address, token []common.Address) (*X402MinimalPaidIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _X402Minimal.contract.FilterLogs(opts, "Paid", payerRule, merchantRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &X402MinimalPaidIterator{contract: _X402Minimal.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0xd32d5970dffb7a78dddc9bf38e51bd1ff7ae0ab4e0d3fe026dfe5c1614b57df7.
//
// Solidity: event Paid(address indexed payer, address indexed merchant, address indexed token, uint256 amount, uint256 nonce, uint256 deadline)
func (_X402Minimal *X402MinimalFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *X402MinimalPaid, payer []common.Address, merchant []common.Address, token []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _X402Minimal.contract.WatchLogs(opts, "Paid", payerRule, merchantRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X402MinimalPaid)
				if err := _X402Minimal.contract.UnpackLog(event, "Paid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaid is a log parse operation binding the contract event 0xd32d5970dffb7a78dddc9bf38e51bd1ff7ae0ab4e0d3fe026dfe5c1614b57df7.
//
// Solidity: event Paid(address indexed payer, address indexed merchant, address indexed token, uint256 amount, uint256 nonce, uint256 deadline)
func (_X402Minimal *X402MinimalFilterer) ParsePaid(log types.Log) (*X402MinimalPaid, error) {
	event := new(X402MinimalPaid)
	if err := _X402Minimal.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
