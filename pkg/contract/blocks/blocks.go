// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocks

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BlocksABI is the input ABI used to generate the binding from.
const BlocksABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"BlockNumberLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xdd631fe0\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_submissionBytes\",\"type\":\"bytes\"}],\"name\":\"submit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xef7fa71b\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xe7f43c68\"},{\"constant\":true,\"inputs\":[{\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getHeaderHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x78802ef1\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x5ae6256d\"}]"

// Blocks is an auto generated Go binding around an Ethereum contract.
type Blocks struct {
	BlocksCaller     // Read-only binding to the contract
	BlocksTransactor // Write-only binding to the contract
	BlocksFilterer   // Log filterer for contract events
}

// BlocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlocksSession struct {
	Contract     *Blocks           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlocksCallerSession struct {
	Contract *BlocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BlocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlocksTransactorSession struct {
	Contract     *BlocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlocksRaw struct {
	Contract *Blocks // Generic contract binding to access the raw methods on
}

// BlocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlocksCallerRaw struct {
	Contract *BlocksCaller // Generic read-only contract binding to access the raw methods on
}

// BlocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlocksTransactorRaw struct {
	Contract *BlocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlocks creates a new instance of Blocks, bound to a specific deployed contract.
func NewBlocks(address common.Address, backend bind.ContractBackend) (*Blocks, error) {
	contract, err := bindBlocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blocks{BlocksCaller: BlocksCaller{contract: contract}, BlocksTransactor: BlocksTransactor{contract: contract}, BlocksFilterer: BlocksFilterer{contract: contract}}, nil
}

// NewBlocksCaller creates a new read-only instance of Blocks, bound to a specific deployed contract.
func NewBlocksCaller(address common.Address, caller bind.ContractCaller) (*BlocksCaller, error) {
	contract, err := bindBlocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlocksCaller{contract: contract}, nil
}

// NewBlocksTransactor creates a new write-only instance of Blocks, bound to a specific deployed contract.
func NewBlocksTransactor(address common.Address, transactor bind.ContractTransactor) (*BlocksTransactor, error) {
	contract, err := bindBlocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlocksTransactor{contract: contract}, nil
}

// NewBlocksFilterer creates a new log filterer instance of Blocks, bound to a specific deployed contract.
func NewBlocksFilterer(address common.Address, filterer bind.ContractFilterer) (*BlocksFilterer, error) {
	contract, err := bindBlocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlocksFilterer{contract: contract}, nil
}

// bindBlocks binds a generic wrapper to an already deployed contract.
func bindBlocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlocksABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blocks *BlocksRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Blocks.Contract.BlocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blocks *BlocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blocks.Contract.BlocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blocks *BlocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blocks.Contract.BlocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blocks *BlocksCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Blocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blocks *BlocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blocks *BlocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blocks.Contract.contract.Transact(opts, method, params...)
}

// BlockNumberLength is a free data retrieval call binding the contract method 0xdd631fe0.
//
// Solidity: function BlockNumberLength() constant returns(uint256)
func (_Blocks *BlocksCaller) BlockNumberLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Blocks.contract.Call(opts, out, "BlockNumberLength")
	return *ret0, err
}

// BlockNumberLength is a free data retrieval call binding the contract method 0xdd631fe0.
//
// Solidity: function BlockNumberLength() constant returns(uint256)
func (_Blocks *BlocksSession) BlockNumberLength() (*big.Int, error) {
	return _Blocks.Contract.BlockNumberLength(&_Blocks.CallOpts)
}

// BlockNumberLength is a free data retrieval call binding the contract method 0xdd631fe0.
//
// Solidity: function BlockNumberLength() constant returns(uint256)
func (_Blocks *BlocksCallerSession) BlockNumberLength() (*big.Int, error) {
	return _Blocks.Contract.BlockNumberLength(&_Blocks.CallOpts)
}

// GetHeaderHash is a free data retrieval call binding the contract method 0x78802ef1.
//
// Solidity: function getHeaderHash(uint64 height) constant returns(bytes32)
func (_Blocks *BlocksCaller) GetHeaderHash(opts *bind.CallOpts, height uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Blocks.contract.Call(opts, out, "getHeaderHash", height)
	return *ret0, err
}

// GetHeaderHash is a free data retrieval call binding the contract method 0x78802ef1.
//
// Solidity: function getHeaderHash(uint64 height) constant returns(bytes32)
func (_Blocks *BlocksSession) GetHeaderHash(height uint64) ([32]byte, error) {
	return _Blocks.Contract.GetHeaderHash(&_Blocks.CallOpts, height)
}

// GetHeaderHash is a free data retrieval call binding the contract method 0x78802ef1.
//
// Solidity: function getHeaderHash(uint64 height) constant returns(bytes32)
func (_Blocks *BlocksCallerSession) GetHeaderHash(height uint64) ([32]byte, error) {
	return _Blocks.Contract.GetHeaderHash(&_Blocks.CallOpts, height)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() constant returns(uint64)
func (_Blocks *BlocksCaller) GetLastBlockNumber(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Blocks.contract.Call(opts, out, "getLastBlockNumber")
	return *ret0, err
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() constant returns(uint64)
func (_Blocks *BlocksSession) GetLastBlockNumber() (uint64, error) {
	return _Blocks.Contract.GetLastBlockNumber(&_Blocks.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() constant returns(uint64)
func (_Blocks *BlocksCallerSession) GetLastBlockNumber() (uint64, error) {
	return _Blocks.Contract.GetLastBlockNumber(&_Blocks.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() constant returns(address)
func (_Blocks *BlocksCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Blocks.contract.Call(opts, out, "getOperator")
	return *ret0, err
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() constant returns(address)
func (_Blocks *BlocksSession) GetOperator() (common.Address, error) {
	return _Blocks.Contract.GetOperator(&_Blocks.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() constant returns(address)
func (_Blocks *BlocksCallerSession) GetOperator() (common.Address, error) {
	return _Blocks.Contract.GetOperator(&_Blocks.CallOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xef7fa71b.
//
// Solidity: function submit(bytes _submissionBytes) returns()
func (_Blocks *BlocksTransactor) Submit(opts *bind.TransactOpts, _submissionBytes []byte) (*types.Transaction, error) {
	return _Blocks.contract.Transact(opts, "submit", _submissionBytes)
}

// Submit is a paid mutator transaction binding the contract method 0xef7fa71b.
//
// Solidity: function submit(bytes _submissionBytes) returns()
func (_Blocks *BlocksSession) Submit(_submissionBytes []byte) (*types.Transaction, error) {
	return _Blocks.Contract.Submit(&_Blocks.TransactOpts, _submissionBytes)
}

// Submit is a paid mutator transaction binding the contract method 0xef7fa71b.
//
// Solidity: function submit(bytes _submissionBytes) returns()
func (_Blocks *BlocksTransactorSession) Submit(_submissionBytes []byte) (*types.Transaction, error) {
	return _Blocks.Contract.Submit(&_Blocks.TransactOpts, _submissionBytes)
}
