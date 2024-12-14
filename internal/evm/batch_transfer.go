// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm

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

// BatchtransferMetaData contains all meta data concerning the Batchtransfer contract.
var BatchtransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103208061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c80631239ec8c1461002d575b5f5ffd5b61004061003b366004610217565b610042565b005b8281146100965760405162461bcd60e51b815260206004820152601a60248201527f4172726179206c656e6774687320646f206e6f74206d6174636800000000000060448201526064015b60405180910390fd5b5f5b838110156101ac57856001600160a01b03166323b872dd338787858181106100c2576100c2610297565b90506020020160208101906100d791906102ab565b8686868181106100e9576100e9610297565b6040516001600160e01b031960e088901b1681526001600160a01b039586166004820152949093166024850152506020909102013560448201526064016020604051808303815f875af1158015610142573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061016691906102cb565b6101a45760405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b604482015260640161008d565b600101610098565b505050505050565b80356001600160a01b03811681146101ca575f5ffd5b919050565b5f5f83601f8401126101df575f5ffd5b50813567ffffffffffffffff8111156101f6575f5ffd5b6020830191508360208260051b8501011115610210575f5ffd5b9250929050565b5f5f5f5f5f6060868803121561022b575f5ffd5b610234866101b4565b9450602086013567ffffffffffffffff81111561024f575f5ffd5b61025b888289016101cf565b909550935050604086013567ffffffffffffffff81111561027a575f5ffd5b610286888289016101cf565b969995985093965092949392505050565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156102bb575f5ffd5b6102c4826101b4565b9392505050565b5f602082840312156102db575f5ffd5b815180151581146102c4575f5ffdfea2646970667358221220a7e55c5b7925ceed192c78287cda3791966e4da8c7738a0fde284efc006fe39664736f6c634300081c0033",
}

// BatchtransferABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchtransferMetaData.ABI instead.
var BatchtransferABI = BatchtransferMetaData.ABI

// BatchtransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchtransferMetaData.Bin instead.
var BatchtransferBin = BatchtransferMetaData.Bin

// DeployBatchtransfer deploys a new Ethereum contract, binding an instance of Batchtransfer to it.
func DeployBatchtransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Batchtransfer, error) {
	parsed, err := BatchtransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchtransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Batchtransfer{BatchtransferCaller: BatchtransferCaller{contract: contract}, BatchtransferTransactor: BatchtransferTransactor{contract: contract}, BatchtransferFilterer: BatchtransferFilterer{contract: contract}}, nil
}

// Batchtransfer is an auto generated Go binding around an Ethereum contract.
type Batchtransfer struct {
	BatchtransferCaller     // Read-only binding to the contract
	BatchtransferTransactor // Write-only binding to the contract
	BatchtransferFilterer   // Log filterer for contract events
}

// BatchtransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchtransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchtransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchtransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchtransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchtransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchtransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchtransferSession struct {
	Contract     *Batchtransfer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchtransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchtransferCallerSession struct {
	Contract *BatchtransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchtransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchtransferTransactorSession struct {
	Contract     *BatchtransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchtransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchtransferRaw struct {
	Contract *Batchtransfer // Generic contract binding to access the raw methods on
}

// BatchtransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchtransferCallerRaw struct {
	Contract *BatchtransferCaller // Generic read-only contract binding to access the raw methods on
}

// BatchtransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchtransferTransactorRaw struct {
	Contract *BatchtransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchtransfer creates a new instance of Batchtransfer, bound to a specific deployed contract.
func NewBatchtransfer(address common.Address, backend bind.ContractBackend) (*Batchtransfer, error) {
	contract, err := bindBatchtransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Batchtransfer{BatchtransferCaller: BatchtransferCaller{contract: contract}, BatchtransferTransactor: BatchtransferTransactor{contract: contract}, BatchtransferFilterer: BatchtransferFilterer{contract: contract}}, nil
}

// NewBatchtransferCaller creates a new read-only instance of Batchtransfer, bound to a specific deployed contract.
func NewBatchtransferCaller(address common.Address, caller bind.ContractCaller) (*BatchtransferCaller, error) {
	contract, err := bindBatchtransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchtransferCaller{contract: contract}, nil
}

// NewBatchtransferTransactor creates a new write-only instance of Batchtransfer, bound to a specific deployed contract.
func NewBatchtransferTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchtransferTransactor, error) {
	contract, err := bindBatchtransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchtransferTransactor{contract: contract}, nil
}

// NewBatchtransferFilterer creates a new log filterer instance of Batchtransfer, bound to a specific deployed contract.
func NewBatchtransferFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchtransferFilterer, error) {
	contract, err := bindBatchtransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchtransferFilterer{contract: contract}, nil
}

// bindBatchtransfer binds a generic wrapper to an already deployed contract.
func bindBatchtransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchtransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batchtransfer *BatchtransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batchtransfer.Contract.BatchtransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batchtransfer *BatchtransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtransfer.Contract.BatchtransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batchtransfer *BatchtransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batchtransfer.Contract.BatchtransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batchtransfer *BatchtransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batchtransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batchtransfer *BatchtransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batchtransfer *BatchtransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batchtransfer.Contract.contract.Transact(opts, method, params...)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x1239ec8c.
//
// Solidity: function batchTransfer(address token, address[] recipients, uint256[] amounts) returns()
func (_Batchtransfer *BatchtransferTransactor) BatchTransfer(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Batchtransfer.contract.Transact(opts, "batchTransfer", token, recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x1239ec8c.
//
// Solidity: function batchTransfer(address token, address[] recipients, uint256[] amounts) returns()
func (_Batchtransfer *BatchtransferSession) BatchTransfer(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Batchtransfer.Contract.BatchTransfer(&_Batchtransfer.TransactOpts, token, recipients, amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x1239ec8c.
//
// Solidity: function batchTransfer(address token, address[] recipients, uint256[] amounts) returns()
func (_Batchtransfer *BatchtransferTransactorSession) BatchTransfer(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Batchtransfer.Contract.BatchTransfer(&_Batchtransfer.TransactOpts, token, recipients, amounts)
}
