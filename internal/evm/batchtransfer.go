// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
)

import (
	"errors"
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
	Bin: "0x6080604052348015600e575f5ffd5b506105758061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c80631239ec8c1461002d575b5f5ffd5b610047600480360381019061004291906102c8565b610049565b005b818190508484905014610091576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610088906103b3565b60405180910390fd5b5f5f90505b848490508110156101a8578573ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8686848181106100d0576100cf6103d1565b5b90506020020160208101906100e591906103fe565b8585858181106100f8576100f76103d1565b5b905060200201356040518363ffffffff1660e01b815260040161011c929190610450565b6020604051808303815f875af1158015610138573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061015c91906104ac565b61019b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019290610521565b60405180910390fd5b8080600101915050610096565b505050505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101e1826101b8565b9050919050565b6101f1816101d7565b81146101fb575f5ffd5b50565b5f8135905061020c816101e8565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261023357610232610212565b5b8235905067ffffffffffffffff8111156102505761024f610216565b5b60208301915083602082028301111561026c5761026b61021a565b5b9250929050565b5f5f83601f84011261028857610287610212565b5b8235905067ffffffffffffffff8111156102a5576102a4610216565b5b6020830191508360208202830111156102c1576102c061021a565b5b9250929050565b5f5f5f5f5f606086880312156102e1576102e06101b0565b5b5f6102ee888289016101fe565b955050602086013567ffffffffffffffff81111561030f5761030e6101b4565b5b61031b8882890161021e565b9450945050604086013567ffffffffffffffff81111561033e5761033d6101b4565b5b61034a88828901610273565b92509250509295509295909350565b5f82825260208201905092915050565b7f4172726179206c656e6774687320646f206e6f74206d617463680000000000005f82015250565b5f61039d601a83610359565b91506103a882610369565b602082019050919050565b5f6020820190508181035f8301526103ca81610391565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215610413576104126101b0565b5b5f610420848285016101fe565b91505092915050565b610432816101d7565b82525050565b5f819050919050565b61044a81610438565b82525050565b5f6040820190506104635f830185610429565b6104706020830184610441565b9392505050565b5f8115159050919050565b61048b81610477565b8114610495575f5ffd5b50565b5f815190506104a681610482565b92915050565b5f602082840312156104c1576104c06101b0565b5b5f6104ce84828501610498565b91505092915050565b7f5472616e73666572206661696c656400000000000000000000000000000000005f82015250565b5f61050b600f83610359565b9150610516826104d7565b602082019050919050565b5f6020820190508181035f830152610538816104ff565b905091905056fea2646970667358221220f47cef9fec6fcd8070a539cfb3dbc9b4a4f283289e4b60cc716ca2285f60818a64736f6c634300081c0033",
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
