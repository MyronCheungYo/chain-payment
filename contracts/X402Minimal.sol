// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title X402Minimal - 基于支付意图的清算合约
/// @notice 支持离线签名、nonce 防重放、ERC20 代扣并记录 Paid 事件
contract X402Minimal {
    struct PaymentIntent {
        address payer;
        address merchant;
        address token;
        uint256 amount;
        uint256 nonce;
        uint256 deadline;
    }

    /// @notice 支付完成事件，便于链下监听
    event Paid(
        address indexed payer,
        address indexed merchant,
        address indexed token,
        uint256 amount,
        uint256 nonce,
        uint256 deadline
    );

    /// @dev payer => nonce => 是否已用，防重放
    mapping(address => mapping(uint256 => bool)) public nonceUsed;

    /// @notice 由 relayer 代付 gas 调用，完成支付扣款
    /// @param intent 支付意图结构体（链下生成并由 payer 签名）
    /// @param signature payer 对 _toEthSignedMessageHash(_hashIntent(intent)) 的签名
    function pay(PaymentIntent calldata intent, bytes calldata signature) external {
        require(block.timestamp <= intent.deadline, "expired");
        require(intent.merchant != address(0) && intent.payer != address(0), "zero");
        require(!nonceUsed[intent.payer][intent.nonce], "nonce-used");

        // 先验证签名，保证是 payer 自己授权
        bytes32 digest = _toEthSignedMessageHash(_hashIntent(intent));
        address signer = _recover(digest, signature);
        require(signer == intent.payer, "bad-sig");

        nonceUsed[intent.payer][intent.nonce] = true;
        // 通过 ERC20 代扣：payer -> merchant
        bool ok = IERC20(intent.token).transferFrom(intent.payer, intent.merchant, intent.amount);
        require(ok, "transfer-fail");

        emit Paid(intent.payer, intent.merchant, intent.token, intent.amount, intent.nonce, intent.deadline);
    }

    /// @dev 将 intent 编码后 keccak256，字段顺序需与前端保持一致
    function _hashIntent(PaymentIntent calldata intent) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                intent.payer,
                intent.merchant,
                intent.token,
                intent.amount,
                intent.nonce,
                intent.deadline
            )
        );
    }

    /// @dev EIP-191 前缀，兼容 eth_sign/personal_sign
    function _toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }

    /// @dev 简易 ECDSA 恢复，不依赖外部库
    function _recover(bytes32 digest, bytes calldata signature) internal pure returns (address) {
        require(signature.length == 65, "sig-len");
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(signature.offset)
            s := calldataload(add(signature.offset, 32))
            v := byte(0, calldataload(add(signature.offset, 64)))
        }
        require(v == 27 || v == 28, "bad-v");
        address signer = ecrecover(digest, v, r, s);
        require(signer != address(0), "ecrecover-zero");
        return signer;
    }
}

interface IERC20 {
    function transferFrom(address from, address to, uint256 value) external returns (bool);
}
