// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

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
)

// CountersCounter is an auto generated low-level Go binding around an user-defined struct.
type CountersCounter struct {
	Value *big.Int
}

// ProxyExchangeOrder is an auto generated low-level Go binding around an user-defined struct.
type ProxyExchangeOrder struct {
	PoolID   *big.Int
	StartAt  uint64
	Duration uint16
	Customer common.Address
	Provider common.Address
}

// ProxyExchangePool is an auto generated low-level Go binding around an user-defined struct.
type ProxyExchangePool struct {
	Provider       common.Address
	Proxies        []*big.Int
	PricePerMinute *big.Int
	ValidBeforeAt  uint64
	Vote           CountersCounter
}

// ProxyExchangeProxy is an auto generated low-level Go binding around an user-defined struct.
type ProxyExchangeProxy struct {
	Url      string
	Location [2]byte
}

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proxies\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"internalType\":\"structCounters.Counter\",\"name\":\"vote\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structProxyExchange.Pool\",\"name\":\"pool\",\"type\":\"tuple\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"duration\",\"type\":\"uint16\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"duration\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"internalType\":\"structProxyExchange.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commissionRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"}],\"name\":\"downVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"orderOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"duration\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"internalType\":\"structProxyExchange.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"orderOfUserAndIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"duration\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"internalType\":\"structProxyExchange.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"poolOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proxies\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"internalType\":\"structCounters.Counter\",\"name\":\"vote\",\"type\":\"tuple\"}],\"internalType\":\"structProxyExchange.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolOfIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proxies\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"internalType\":\"structCounters.Counter\",\"name\":\"vote\",\"type\":\"tuple\"}],\"internalType\":\"structProxyExchange.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"poolOfProviderAndIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proxies\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"internalType\":\"structCounters.Counter\",\"name\":\"vote\",\"type\":\"tuple\"}],\"internalType\":\"structProxyExchange.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"proxyOf\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes2\",\"name\":\"location\",\"type\":\"bytes2\"}],\"internalType\":\"structProxyExchange.Proxy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes2\",\"name\":\"location\",\"type\":\"bytes2\"}],\"internalType\":\"structProxyExchange.Proxy[]\",\"name\":\"proxies\",\"type\":\"tuple[]\"}],\"name\":\"publish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proxies\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"pricePerMinute\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"internalType\":\"structCounters.Counter\",\"name\":\"vote\",\"type\":\"tuple\"}],\"internalType\":\"structProxyExchange.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"}],\"name\":\"recallVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"setCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"validBeforeAt\",\"type\":\"uint64\"}],\"name\":\"setPoolValidBeforeAt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"totalOrderOfUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"totalPoolOfProvider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderID\",\"type\":\"uint256\"}],\"name\":\"upVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint8)
func (_Exchange *ExchangeCaller) CommissionRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "commissionRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint8)
func (_Exchange *ExchangeSession) CommissionRate() (uint8, error) {
	return _Exchange.Contract.CommissionRate(&_Exchange.CallOpts)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint8)
func (_Exchange *ExchangeCallerSession) CommissionRate() (uint8, error) {
	return _Exchange.Contract.CommissionRate(&_Exchange.CallOpts)
}

// IsVoted is a free data retrieval call binding the contract method 0xbb660d81.
//
// Solidity: function isVoted(uint256 orderID) view returns(bool)
func (_Exchange *ExchangeCaller) IsVoted(opts *bind.CallOpts, orderID *big.Int) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isVoted", orderID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoted is a free data retrieval call binding the contract method 0xbb660d81.
//
// Solidity: function isVoted(uint256 orderID) view returns(bool)
func (_Exchange *ExchangeSession) IsVoted(orderID *big.Int) (bool, error) {
	return _Exchange.Contract.IsVoted(&_Exchange.CallOpts, orderID)
}

// IsVoted is a free data retrieval call binding the contract method 0xbb660d81.
//
// Solidity: function isVoted(uint256 orderID) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsVoted(orderID *big.Int) (bool, error) {
	return _Exchange.Contract.IsVoted(&_Exchange.CallOpts, orderID)
}

// OrderOf is a free data retrieval call binding the contract method 0xc7e85f94.
//
// Solidity: function orderOf(uint256 id) view returns((uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeCaller) OrderOf(opts *bind.CallOpts, id *big.Int) (ProxyExchangeOrder, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "orderOf", id)

	if err != nil {
		return *new(ProxyExchangeOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(ProxyExchangeOrder)).(*ProxyExchangeOrder)

	return out0, err

}

// OrderOf is a free data retrieval call binding the contract method 0xc7e85f94.
//
// Solidity: function orderOf(uint256 id) view returns((uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeSession) OrderOf(id *big.Int) (ProxyExchangeOrder, error) {
	return _Exchange.Contract.OrderOf(&_Exchange.CallOpts, id)
}

// OrderOf is a free data retrieval call binding the contract method 0xc7e85f94.
//
// Solidity: function orderOf(uint256 id) view returns((uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeCallerSession) OrderOf(id *big.Int) (ProxyExchangeOrder, error) {
	return _Exchange.Contract.OrderOf(&_Exchange.CallOpts, id)
}

// OrderOfUserAndIndex is a free data retrieval call binding the contract method 0xb953f4e6.
//
// Solidity: function orderOfUserAndIndex(address user, uint256 index) view returns(uint256, (uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeCaller) OrderOfUserAndIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (*big.Int, ProxyExchangeOrder, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "orderOfUserAndIndex", user, index)

	if err != nil {
		return *new(*big.Int), *new(ProxyExchangeOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(ProxyExchangeOrder)).(*ProxyExchangeOrder)

	return out0, out1, err

}

// OrderOfUserAndIndex is a free data retrieval call binding the contract method 0xb953f4e6.
//
// Solidity: function orderOfUserAndIndex(address user, uint256 index) view returns(uint256, (uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeSession) OrderOfUserAndIndex(user common.Address, index *big.Int) (*big.Int, ProxyExchangeOrder, error) {
	return _Exchange.Contract.OrderOfUserAndIndex(&_Exchange.CallOpts, user, index)
}

// OrderOfUserAndIndex is a free data retrieval call binding the contract method 0xb953f4e6.
//
// Solidity: function orderOfUserAndIndex(address user, uint256 index) view returns(uint256, (uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeCallerSession) OrderOfUserAndIndex(user common.Address, index *big.Int) (*big.Int, ProxyExchangeOrder, error) {
	return _Exchange.Contract.OrderOfUserAndIndex(&_Exchange.CallOpts, user, index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// PoolOf is a free data retrieval call binding the contract method 0x83966021.
//
// Solidity: function poolOf(uint256 id) view returns((address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeCaller) PoolOf(opts *bind.CallOpts, id *big.Int) (ProxyExchangePool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "poolOf", id)

	if err != nil {
		return *new(ProxyExchangePool), err
	}

	out0 := *abi.ConvertType(out[0], new(ProxyExchangePool)).(*ProxyExchangePool)

	return out0, err

}

// PoolOf is a free data retrieval call binding the contract method 0x83966021.
//
// Solidity: function poolOf(uint256 id) view returns((address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeSession) PoolOf(id *big.Int) (ProxyExchangePool, error) {
	return _Exchange.Contract.PoolOf(&_Exchange.CallOpts, id)
}

// PoolOf is a free data retrieval call binding the contract method 0x83966021.
//
// Solidity: function poolOf(uint256 id) view returns((address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeCallerSession) PoolOf(id *big.Int) (ProxyExchangePool, error) {
	return _Exchange.Contract.PoolOf(&_Exchange.CallOpts, id)
}

// PoolOfIndex is a free data retrieval call binding the contract method 0xa4596bb1.
//
// Solidity: function poolOfIndex(uint256 index) view returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeCaller) PoolOfIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, ProxyExchangePool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "poolOfIndex", index)

	if err != nil {
		return *new(*big.Int), *new(ProxyExchangePool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(ProxyExchangePool)).(*ProxyExchangePool)

	return out0, out1, err

}

// PoolOfIndex is a free data retrieval call binding the contract method 0xa4596bb1.
//
// Solidity: function poolOfIndex(uint256 index) view returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeSession) PoolOfIndex(index *big.Int) (*big.Int, ProxyExchangePool, error) {
	return _Exchange.Contract.PoolOfIndex(&_Exchange.CallOpts, index)
}

// PoolOfIndex is a free data retrieval call binding the contract method 0xa4596bb1.
//
// Solidity: function poolOfIndex(uint256 index) view returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeCallerSession) PoolOfIndex(index *big.Int) (*big.Int, ProxyExchangePool, error) {
	return _Exchange.Contract.PoolOfIndex(&_Exchange.CallOpts, index)
}

// PoolOfProviderAndIndex is a free data retrieval call binding the contract method 0xd79f3cb1.
//
// Solidity: function poolOfProviderAndIndex(address provider, uint256 index) view returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeCaller) PoolOfProviderAndIndex(opts *bind.CallOpts, provider common.Address, index *big.Int) (*big.Int, ProxyExchangePool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "poolOfProviderAndIndex", provider, index)

	if err != nil {
		return *new(*big.Int), *new(ProxyExchangePool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(ProxyExchangePool)).(*ProxyExchangePool)

	return out0, out1, err

}

// PoolOfProviderAndIndex is a free data retrieval call binding the contract method 0xd79f3cb1.
//
// Solidity: function poolOfProviderAndIndex(address provider, uint256 index) view returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeSession) PoolOfProviderAndIndex(provider common.Address, index *big.Int) (*big.Int, ProxyExchangePool, error) {
	return _Exchange.Contract.PoolOfProviderAndIndex(&_Exchange.CallOpts, provider, index)
}

// PoolOfProviderAndIndex is a free data retrieval call binding the contract method 0xd79f3cb1.
//
// Solidity: function poolOfProviderAndIndex(address provider, uint256 index) view returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeCallerSession) PoolOfProviderAndIndex(provider common.Address, index *big.Int) (*big.Int, ProxyExchangePool, error) {
	return _Exchange.Contract.PoolOfProviderAndIndex(&_Exchange.CallOpts, provider, index)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Exchange *ExchangeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Exchange *ExchangeSession) ProxiableUUID() ([32]byte, error) {
	return _Exchange.Contract.ProxiableUUID(&_Exchange.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Exchange.Contract.ProxiableUUID(&_Exchange.CallOpts)
}

// ProxyOf is a free data retrieval call binding the contract method 0xb8867485.
//
// Solidity: function proxyOf(uint256 id) view returns((string,bytes2))
func (_Exchange *ExchangeCaller) ProxyOf(opts *bind.CallOpts, id *big.Int) (ProxyExchangeProxy, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "proxyOf", id)

	if err != nil {
		return *new(ProxyExchangeProxy), err
	}

	out0 := *abi.ConvertType(out[0], new(ProxyExchangeProxy)).(*ProxyExchangeProxy)

	return out0, err

}

// ProxyOf is a free data retrieval call binding the contract method 0xb8867485.
//
// Solidity: function proxyOf(uint256 id) view returns((string,bytes2))
func (_Exchange *ExchangeSession) ProxyOf(id *big.Int) (ProxyExchangeProxy, error) {
	return _Exchange.Contract.ProxyOf(&_Exchange.CallOpts, id)
}

// ProxyOf is a free data retrieval call binding the contract method 0xb8867485.
//
// Solidity: function proxyOf(uint256 id) view returns((string,bytes2))
func (_Exchange *ExchangeCallerSession) ProxyOf(id *big.Int) (ProxyExchangeProxy, error) {
	return _Exchange.Contract.ProxyOf(&_Exchange.CallOpts, id)
}

// TotalOrder is a free data retrieval call binding the contract method 0x216e2a97.
//
// Solidity: function totalOrder() view returns(uint256)
func (_Exchange *ExchangeCaller) TotalOrder(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "totalOrder")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOrder is a free data retrieval call binding the contract method 0x216e2a97.
//
// Solidity: function totalOrder() view returns(uint256)
func (_Exchange *ExchangeSession) TotalOrder() (*big.Int, error) {
	return _Exchange.Contract.TotalOrder(&_Exchange.CallOpts)
}

// TotalOrder is a free data retrieval call binding the contract method 0x216e2a97.
//
// Solidity: function totalOrder() view returns(uint256)
func (_Exchange *ExchangeCallerSession) TotalOrder() (*big.Int, error) {
	return _Exchange.Contract.TotalOrder(&_Exchange.CallOpts)
}

// TotalOrderOfUser is a free data retrieval call binding the contract method 0x1f57813b.
//
// Solidity: function totalOrderOfUser(address user) view returns(uint256)
func (_Exchange *ExchangeCaller) TotalOrderOfUser(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "totalOrderOfUser", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOrderOfUser is a free data retrieval call binding the contract method 0x1f57813b.
//
// Solidity: function totalOrderOfUser(address user) view returns(uint256)
func (_Exchange *ExchangeSession) TotalOrderOfUser(user common.Address) (*big.Int, error) {
	return _Exchange.Contract.TotalOrderOfUser(&_Exchange.CallOpts, user)
}

// TotalOrderOfUser is a free data retrieval call binding the contract method 0x1f57813b.
//
// Solidity: function totalOrderOfUser(address user) view returns(uint256)
func (_Exchange *ExchangeCallerSession) TotalOrderOfUser(user common.Address) (*big.Int, error) {
	return _Exchange.Contract.TotalOrderOfUser(&_Exchange.CallOpts, user)
}

// TotalPool is a free data retrieval call binding the contract method 0xecfb49a3.
//
// Solidity: function totalPool() view returns(uint256)
func (_Exchange *ExchangeCaller) TotalPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "totalPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPool is a free data retrieval call binding the contract method 0xecfb49a3.
//
// Solidity: function totalPool() view returns(uint256)
func (_Exchange *ExchangeSession) TotalPool() (*big.Int, error) {
	return _Exchange.Contract.TotalPool(&_Exchange.CallOpts)
}

// TotalPool is a free data retrieval call binding the contract method 0xecfb49a3.
//
// Solidity: function totalPool() view returns(uint256)
func (_Exchange *ExchangeCallerSession) TotalPool() (*big.Int, error) {
	return _Exchange.Contract.TotalPool(&_Exchange.CallOpts)
}

// TotalPoolOfProvider is a free data retrieval call binding the contract method 0x63301f61.
//
// Solidity: function totalPoolOfProvider(address provider) view returns(uint256)
func (_Exchange *ExchangeCaller) TotalPoolOfProvider(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "totalPoolOfProvider", provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPoolOfProvider is a free data retrieval call binding the contract method 0x63301f61.
//
// Solidity: function totalPoolOfProvider(address provider) view returns(uint256)
func (_Exchange *ExchangeSession) TotalPoolOfProvider(provider common.Address) (*big.Int, error) {
	return _Exchange.Contract.TotalPoolOfProvider(&_Exchange.CallOpts, provider)
}

// TotalPoolOfProvider is a free data retrieval call binding the contract method 0x63301f61.
//
// Solidity: function totalPoolOfProvider(address provider) view returns(uint256)
func (_Exchange *ExchangeCallerSession) TotalPoolOfProvider(provider common.Address) (*big.Int, error) {
	return _Exchange.Contract.TotalPoolOfProvider(&_Exchange.CallOpts, provider)
}

// Buy is a paid mutator transaction binding the contract method 0xef3f57cc.
//
// Solidity: function buy(uint256 poolID, uint64 startAt, uint16 duration) payable returns(uint256, (uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeTransactor) Buy(opts *bind.TransactOpts, poolID *big.Int, startAt uint64, duration uint16) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "buy", poolID, startAt, duration)
}

// Buy is a paid mutator transaction binding the contract method 0xef3f57cc.
//
// Solidity: function buy(uint256 poolID, uint64 startAt, uint16 duration) payable returns(uint256, (uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeSession) Buy(poolID *big.Int, startAt uint64, duration uint16) (*types.Transaction, error) {
	return _Exchange.Contract.Buy(&_Exchange.TransactOpts, poolID, startAt, duration)
}

// Buy is a paid mutator transaction binding the contract method 0xef3f57cc.
//
// Solidity: function buy(uint256 poolID, uint64 startAt, uint16 duration) payable returns(uint256, (uint256,uint64,uint16,address,address))
func (_Exchange *ExchangeTransactorSession) Buy(poolID *big.Int, startAt uint64, duration uint16) (*types.Transaction, error) {
	return _Exchange.Contract.Buy(&_Exchange.TransactOpts, poolID, startAt, duration)
}

// DownVote is a paid mutator transaction binding the contract method 0x1ae9392d.
//
// Solidity: function downVote(uint256 orderID) returns()
func (_Exchange *ExchangeTransactor) DownVote(opts *bind.TransactOpts, orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "downVote", orderID)
}

// DownVote is a paid mutator transaction binding the contract method 0x1ae9392d.
//
// Solidity: function downVote(uint256 orderID) returns()
func (_Exchange *ExchangeSession) DownVote(orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DownVote(&_Exchange.TransactOpts, orderID)
}

// DownVote is a paid mutator transaction binding the contract method 0x1ae9392d.
//
// Solidity: function downVote(uint256 orderID) returns()
func (_Exchange *ExchangeTransactorSession) DownVote(orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DownVote(&_Exchange.TransactOpts, orderID)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Exchange *ExchangeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Exchange *ExchangeSession) Initialize() (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Exchange *ExchangeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts)
}

// Publish is a paid mutator transaction binding the contract method 0x3c1fe94c.
//
// Solidity: function publish(uint256 pricePerMinute, uint64 validBeforeAt, (string,bytes2)[] proxies) returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeTransactor) Publish(opts *bind.TransactOpts, pricePerMinute *big.Int, validBeforeAt uint64, proxies []ProxyExchangeProxy) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "publish", pricePerMinute, validBeforeAt, proxies)
}

// Publish is a paid mutator transaction binding the contract method 0x3c1fe94c.
//
// Solidity: function publish(uint256 pricePerMinute, uint64 validBeforeAt, (string,bytes2)[] proxies) returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeSession) Publish(pricePerMinute *big.Int, validBeforeAt uint64, proxies []ProxyExchangeProxy) (*types.Transaction, error) {
	return _Exchange.Contract.Publish(&_Exchange.TransactOpts, pricePerMinute, validBeforeAt, proxies)
}

// Publish is a paid mutator transaction binding the contract method 0x3c1fe94c.
//
// Solidity: function publish(uint256 pricePerMinute, uint64 validBeforeAt, (string,bytes2)[] proxies) returns(uint256, (address,uint256[],uint256,uint64,(uint256)))
func (_Exchange *ExchangeTransactorSession) Publish(pricePerMinute *big.Int, validBeforeAt uint64, proxies []ProxyExchangeProxy) (*types.Transaction, error) {
	return _Exchange.Contract.Publish(&_Exchange.TransactOpts, pricePerMinute, validBeforeAt, proxies)
}

// RecallVote is a paid mutator transaction binding the contract method 0x89076f47.
//
// Solidity: function recallVote(uint256 orderID) returns()
func (_Exchange *ExchangeTransactor) RecallVote(opts *bind.TransactOpts, orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "recallVote", orderID)
}

// RecallVote is a paid mutator transaction binding the contract method 0x89076f47.
//
// Solidity: function recallVote(uint256 orderID) returns()
func (_Exchange *ExchangeSession) RecallVote(orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RecallVote(&_Exchange.TransactOpts, orderID)
}

// RecallVote is a paid mutator transaction binding the contract method 0x89076f47.
//
// Solidity: function recallVote(uint256 orderID) returns()
func (_Exchange *ExchangeTransactorSession) RecallVote(orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.RecallVote(&_Exchange.TransactOpts, orderID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x12dee489.
//
// Solidity: function setCommissionRate(uint8 v) returns()
func (_Exchange *ExchangeTransactor) SetCommissionRate(opts *bind.TransactOpts, v uint8) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setCommissionRate", v)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x12dee489.
//
// Solidity: function setCommissionRate(uint8 v) returns()
func (_Exchange *ExchangeSession) SetCommissionRate(v uint8) (*types.Transaction, error) {
	return _Exchange.Contract.SetCommissionRate(&_Exchange.TransactOpts, v)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x12dee489.
//
// Solidity: function setCommissionRate(uint8 v) returns()
func (_Exchange *ExchangeTransactorSession) SetCommissionRate(v uint8) (*types.Transaction, error) {
	return _Exchange.Contract.SetCommissionRate(&_Exchange.TransactOpts, v)
}

// SetPoolValidBeforeAt is a paid mutator transaction binding the contract method 0x406aabb9.
//
// Solidity: function setPoolValidBeforeAt(uint256 poolID, uint64 validBeforeAt) returns()
func (_Exchange *ExchangeTransactor) SetPoolValidBeforeAt(opts *bind.TransactOpts, poolID *big.Int, validBeforeAt uint64) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setPoolValidBeforeAt", poolID, validBeforeAt)
}

// SetPoolValidBeforeAt is a paid mutator transaction binding the contract method 0x406aabb9.
//
// Solidity: function setPoolValidBeforeAt(uint256 poolID, uint64 validBeforeAt) returns()
func (_Exchange *ExchangeSession) SetPoolValidBeforeAt(poolID *big.Int, validBeforeAt uint64) (*types.Transaction, error) {
	return _Exchange.Contract.SetPoolValidBeforeAt(&_Exchange.TransactOpts, poolID, validBeforeAt)
}

// SetPoolValidBeforeAt is a paid mutator transaction binding the contract method 0x406aabb9.
//
// Solidity: function setPoolValidBeforeAt(uint256 poolID, uint64 validBeforeAt) returns()
func (_Exchange *ExchangeTransactorSession) SetPoolValidBeforeAt(poolID *big.Int, validBeforeAt uint64) (*types.Transaction, error) {
	return _Exchange.Contract.SetPoolValidBeforeAt(&_Exchange.TransactOpts, poolID, validBeforeAt)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// UpVote is a paid mutator transaction binding the contract method 0xec859db5.
//
// Solidity: function upVote(uint256 orderID) returns()
func (_Exchange *ExchangeTransactor) UpVote(opts *bind.TransactOpts, orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "upVote", orderID)
}

// UpVote is a paid mutator transaction binding the contract method 0xec859db5.
//
// Solidity: function upVote(uint256 orderID) returns()
func (_Exchange *ExchangeSession) UpVote(orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.UpVote(&_Exchange.TransactOpts, orderID)
}

// UpVote is a paid mutator transaction binding the contract method 0xec859db5.
//
// Solidity: function upVote(uint256 orderID) returns()
func (_Exchange *ExchangeTransactorSession) UpVote(orderID *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.UpVote(&_Exchange.TransactOpts, orderID)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Exchange *ExchangeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Exchange *ExchangeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeTo(&_Exchange.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Exchange *ExchangeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeTo(&_Exchange.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Exchange *ExchangeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Exchange *ExchangeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeToAndCall(&_Exchange.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Exchange *ExchangeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Exchange.Contract.UpgradeToAndCall(&_Exchange.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Exchange *ExchangeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Exchange *ExchangeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Withdraw(&_Exchange.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Exchange *ExchangeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Withdraw(&_Exchange.TransactOpts, amount)
}

// ExchangeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Exchange contract.
type ExchangeAdminChangedIterator struct {
	Event *ExchangeAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExchangeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeAdminChanged)
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
		it.Event = new(ExchangeAdminChanged)
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
func (it *ExchangeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeAdminChanged represents a AdminChanged event raised by the Exchange contract.
type ExchangeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Exchange *ExchangeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ExchangeAdminChangedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ExchangeAdminChangedIterator{contract: _Exchange.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Exchange *ExchangeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ExchangeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeAdminChanged)
				if err := _Exchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Exchange *ExchangeFilterer) ParseAdminChanged(log types.Log) (*ExchangeAdminChanged, error) {
	event := new(ExchangeAdminChanged)
	if err := _Exchange.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Exchange contract.
type ExchangeBeaconUpgradedIterator struct {
	Event *ExchangeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ExchangeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeBeaconUpgraded)
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
		it.Event = new(ExchangeBeaconUpgraded)
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
func (it *ExchangeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeBeaconUpgraded represents a BeaconUpgraded event raised by the Exchange contract.
type ExchangeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Exchange *ExchangeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ExchangeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeBeaconUpgradedIterator{contract: _Exchange.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Exchange *ExchangeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ExchangeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeBeaconUpgraded)
				if err := _Exchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Exchange *ExchangeFilterer) ParseBeaconUpgraded(log types.Log) (*ExchangeBeaconUpgraded, error) {
	event := new(ExchangeBeaconUpgraded)
	if err := _Exchange.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Exchange contract.
type ExchangeInitializedIterator struct {
	Event *ExchangeInitialized // Event containing the contract specifics and raw log

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
func (it *ExchangeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeInitialized)
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
		it.Event = new(ExchangeInitialized)
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
func (it *ExchangeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeInitialized represents a Initialized event raised by the Exchange contract.
type ExchangeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exchange *ExchangeFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExchangeInitializedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExchangeInitializedIterator{contract: _Exchange.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exchange *ExchangeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExchangeInitialized) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeInitialized)
				if err := _Exchange.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Exchange *ExchangeFilterer) ParseInitialized(log types.Log) (*ExchangeInitialized, error) {
	event := new(ExchangeInitialized)
	if err := _Exchange.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*ExchangeOwnershipTransferred, error) {
	event := new(ExchangeOwnershipTransferred)
	if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the Exchange contract.
type ExchangePublishIterator struct {
	Event *ExchangePublish // Event containing the contract specifics and raw log

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
func (it *ExchangePublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePublish)
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
		it.Event = new(ExchangePublish)
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
func (it *ExchangePublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePublish represents a Publish event raised by the Exchange contract.
type ExchangePublish struct {
	PoolID *big.Int
	Pool   ProxyExchangePool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x1a731ad491ca0ec444ee0041b1d50711d1bda92cea8244d0e6c3061c6261c0b6.
//
// Solidity: event Publish(uint256 poolID, (address,uint256[],uint256,uint64,(uint256)) pool)
func (_Exchange *ExchangeFilterer) FilterPublish(opts *bind.FilterOpts) (*ExchangePublishIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return &ExchangePublishIterator{contract: _Exchange.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x1a731ad491ca0ec444ee0041b1d50711d1bda92cea8244d0e6c3061c6261c0b6.
//
// Solidity: event Publish(uint256 poolID, (address,uint256[],uint256,uint64,(uint256)) pool)
func (_Exchange *ExchangeFilterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *ExchangePublish) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Publish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePublish)
				if err := _Exchange.contract.UnpackLog(event, "Publish", log); err != nil {
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

// ParsePublish is a log parse operation binding the contract event 0x1a731ad491ca0ec444ee0041b1d50711d1bda92cea8244d0e6c3061c6261c0b6.
//
// Solidity: event Publish(uint256 poolID, (address,uint256[],uint256,uint64,(uint256)) pool)
func (_Exchange *ExchangeFilterer) ParsePublish(log types.Log) (*ExchangePublish, error) {
	event := new(ExchangePublish)
	if err := _Exchange.contract.UnpackLog(event, "Publish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Exchange contract.
type ExchangeUpgradedIterator struct {
	Event *ExchangeUpgraded // Event containing the contract specifics and raw log

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
func (it *ExchangeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeUpgraded)
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
		it.Event = new(ExchangeUpgraded)
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
func (it *ExchangeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeUpgraded represents a Upgraded event raised by the Exchange contract.
type ExchangeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Exchange *ExchangeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ExchangeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeUpgradedIterator{contract: _Exchange.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Exchange *ExchangeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ExchangeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeUpgraded)
				if err := _Exchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Exchange *ExchangeFilterer) ParseUpgraded(log types.Log) (*ExchangeUpgraded, error) {
	event := new(ExchangeUpgraded)
	if err := _Exchange.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
