package evm

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type ERC20s interface {
	Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
	Decimals(opts *bind.CallOpts) (uint8, error)
}

// Airdrop struct to manage token distribution
type Airdrop struct {
	Client       *ethclient.Client
	PrivateKey   *ecdsa.PrivateKey
	TokenAddress common.Address
	ChainID      *big.Int
	Token        ERC20s
}

// AirdropError represents a custom error type for airdrop-specific errors
type AirdropError struct {
	Code    string
	Message string
	Err     error
}
