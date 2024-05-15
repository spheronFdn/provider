// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenRegistry

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

// TokenRegistryMetaData contains all meta data concerning the TokenRegistry contract.
var TokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ticker\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenAddress\",\"type\":\"address\"}],\"name\":\"TokenUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ticker\",\"type\":\"string\"}],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ticker\",\"type\":\"string\"}],\"name\":\"isTokenRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ticker\",\"type\":\"string\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ticker\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_newTokenAddress\",\"type\":\"address\"}],\"name\":\"updateTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110938061005c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063935b13f611610059578063935b13f614610110578063b0b22c2c14610140578063c40912361461015c578063f2fde38b1461018c57610086565b80631f4559221461008a57806322b9e3b6146100a65780632efc7a94146100d65780638da5cb5b146100f2575b5f80fd5b6100a4600480360381019061009f9190610b8a565b6101a8565b005b6100c060048036038101906100bb9190610b8a565b61035a565b6040516100cd9190610beb565b60405180910390f35b6100f060048036038101906100eb9190610c5e565b6103ce565b005b6100fa61060d565b6040516101079190610cc7565b60405180910390f35b61012a60048036038101906101259190610b8a565b610632565b6040516101379190610cc7565b60405180910390f35b61015a60048036038101906101559190610c5e565b610679565b005b61017660048036038101906101719190610b8a565b6108b8565b6040516101839190610cc7565b60405180910390f35b6101a660048036038101906101a19190610ce0565b6108fd565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610237576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022e90610d8b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff165f8260405161025d9190610dfb565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036102e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d890610e5b565b60405180910390fd5b5f816040516102f09190610dfb565b90815260200160405180910390205f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690557fcdd02d37d444f381ccea186dcdaf556989dcdeeebafc5cd307fe5e7609cd88c38160405161034f9190610eb1565b60405180910390a150565b5f8073ffffffffffffffffffffffffffffffffffffffff165f836040516103819190610dfb565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045490610d8b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff165f836040516104839190610dfb565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610507576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104fe90610e5b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610575576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056c90610f1b565b60405180910390fd5b805f836040516105859190610dfb565b90815260200160405180910390205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fab9b2a4b5369b2ede85df0177e1308b7a6e3575b7b9965d79f3a32f97510e5978282604051610601929190610f39565b60405180910390a15050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f818051602081018201805184825260208301602085012081835280955050505050505f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610708576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ff90610d8b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff165f8360405161072e9190610dfb565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a990610fb1565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610820576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081790610f1b565b60405180910390fd5b805f836040516108309190610dfb565b90815260200160405180910390205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f6d97a31531bdb3c43d920bc996ec8b9fd7733f5f55a81cd74700f88631fc603282826040516108ac929190610f39565b60405180910390a15050565b5f80826040516108c89190610dfb565b90815260200160405180910390205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461098c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098390610d8b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f19061103f565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a9c82610a56565b810181811067ffffffffffffffff82111715610abb57610aba610a66565b5b80604052505050565b5f610acd610a3d565b9050610ad98282610a93565b919050565b5f67ffffffffffffffff821115610af857610af7610a66565b5b610b0182610a56565b9050602081019050919050565b828183375f83830152505050565b5f610b2e610b2984610ade565b610ac4565b905082815260208101848484011115610b4a57610b49610a52565b5b610b55848285610b0e565b509392505050565b5f82601f830112610b7157610b70610a4e565b5b8135610b81848260208601610b1c565b91505092915050565b5f60208284031215610b9f57610b9e610a46565b5b5f82013567ffffffffffffffff811115610bbc57610bbb610a4a565b5b610bc884828501610b5d565b91505092915050565b5f8115159050919050565b610be581610bd1565b82525050565b5f602082019050610bfe5f830184610bdc565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610c2d82610c04565b9050919050565b610c3d81610c23565b8114610c47575f80fd5b50565b5f81359050610c5881610c34565b92915050565b5f8060408385031215610c7457610c73610a46565b5b5f83013567ffffffffffffffff811115610c9157610c90610a4a565b5b610c9d85828601610b5d565b9250506020610cae85828601610c4a565b9150509250929050565b610cc181610c23565b82525050565b5f602082019050610cda5f830184610cb8565b92915050565b5f60208284031215610cf557610cf4610a46565b5b5f610d0284828501610c4a565b91505092915050565b5f82825260208201905092915050565b7f4f6e6c7920746865206f776e65722063616e20706572666f726d2074686973205f8201527f616374696f6e0000000000000000000000000000000000000000000000000000602082015250565b5f610d75602683610d0b565b9150610d8082610d1b565b604082019050919050565b5f6020820190508181035f830152610da281610d69565b9050919050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f610dd582610da9565b610ddf8185610db3565b9350610def818560208601610dbd565b80840191505092915050565b5f610e068284610dcb565b915081905092915050565b7f546f6b656e206e6f7420726567697374657265642e00000000000000000000005f82015250565b5f610e45601583610d0b565b9150610e5082610e11565b602082019050919050565b5f6020820190508181035f830152610e7281610e39565b9050919050565b5f610e8382610da9565b610e8d8185610d0b565b9350610e9d818560208601610dbd565b610ea681610a56565b840191505092915050565b5f6020820190508181035f830152610ec98184610e79565b905092915050565b7f496e76616c696420546f6b656e204164647265737300000000000000000000005f82015250565b5f610f05601583610d0b565b9150610f1082610ed1565b602082019050919050565b5f6020820190508181035f830152610f3281610ef9565b9050919050565b5f6040820190508181035f830152610f518185610e79565b9050610f606020830184610cb8565b9392505050565b7f546f6b656e20616c726561647920726567697374657265642e000000000000005f82015250565b5f610f9b601983610d0b565b9150610fa682610f67565b602082019050919050565b5f6020820190508181035f830152610fc881610f8f565b9050919050565b7f4e6577206f776e65722063616e6e6f7420626520746865207a65726f206164645f8201527f726573732e000000000000000000000000000000000000000000000000000000602082015250565b5f611029602583610d0b565b915061103482610fcf565b604082019050919050565b5f6020820190508181035f8301526110568161101d565b905091905056fea264697066735822122024f92a0444002cdcd3f6d4c27a00067effd6dbb9de9472854d4c786f60d1b73c64736f6c63430008190033",
}

// TokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRegistryMetaData.ABI instead.
var TokenRegistryABI = TokenRegistryMetaData.ABI

// TokenRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenRegistryMetaData.Bin instead.
var TokenRegistryBin = TokenRegistryMetaData.Bin

// DeployTokenRegistry deploys a new Ethereum contract, binding an instance of TokenRegistry to it.
func DeployTokenRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenRegistry, error) {
	parsed, err := TokenRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// TokenRegistry is an auto generated Go binding around an Ethereum contract.
type TokenRegistry struct {
	TokenRegistryCaller     // Read-only binding to the contract
	TokenRegistryTransactor // Write-only binding to the contract
	TokenRegistryFilterer   // Log filterer for contract events
}

// TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRegistrySession struct {
	Contract     *TokenRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRegistryCallerSession struct {
	Contract *TokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRegistryTransactorSession struct {
	Contract     *TokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRegistryRaw struct {
	Contract *TokenRegistry // Generic contract binding to access the raw methods on
}

// TokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRegistryCallerRaw struct {
	Contract *TokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRegistryTransactorRaw struct {
	Contract *TokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRegistry creates a new instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistry(address common.Address, backend bind.ContractBackend) (*TokenRegistry, error) {
	contract, err := bindTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// NewTokenRegistryCaller creates a new read-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenRegistryCaller, error) {
	contract, err := bindTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryCaller{contract: contract}, nil
}

// NewTokenRegistryTransactor creates a new write-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegistryTransactor, error) {
	contract, err := bindTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTransactor{contract: contract}, nil
}

// NewTokenRegistryFilterer creates a new log filterer instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRegistryFilterer, error) {
	contract, err := bindTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryFilterer{contract: contract}, nil
}

// bindTokenRegistry binds a generic wrapper to an already deployed contract.
func bindTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.TokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0xc4091236.
//
// Solidity: function getTokenAddress(string _ticker) view returns(address)
func (_TokenRegistry *TokenRegistryCaller) GetTokenAddress(opts *bind.CallOpts, _ticker string) (common.Address, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getTokenAddress", _ticker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0xc4091236.
//
// Solidity: function getTokenAddress(string _ticker) view returns(address)
func (_TokenRegistry *TokenRegistrySession) GetTokenAddress(_ticker string) (common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddress(&_TokenRegistry.CallOpts, _ticker)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0xc4091236.
//
// Solidity: function getTokenAddress(string _ticker) view returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenAddress(_ticker string) (common.Address, error) {
	return _TokenRegistry.Contract.GetTokenAddress(&_TokenRegistry.CallOpts, _ticker)
}

// IsTokenRegistered is a free data retrieval call binding the contract method 0x22b9e3b6.
//
// Solidity: function isTokenRegistered(string _ticker) view returns(bool)
func (_TokenRegistry *TokenRegistryCaller) IsTokenRegistered(opts *bind.CallOpts, _ticker string) (bool, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "isTokenRegistered", _ticker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenRegistered is a free data retrieval call binding the contract method 0x22b9e3b6.
//
// Solidity: function isTokenRegistered(string _ticker) view returns(bool)
func (_TokenRegistry *TokenRegistrySession) IsTokenRegistered(_ticker string) (bool, error) {
	return _TokenRegistry.Contract.IsTokenRegistered(&_TokenRegistry.CallOpts, _ticker)
}

// IsTokenRegistered is a free data retrieval call binding the contract method 0x22b9e3b6.
//
// Solidity: function isTokenRegistered(string _ticker) view returns(bool)
func (_TokenRegistry *TokenRegistryCallerSession) IsTokenRegistered(_ticker string) (bool, error) {
	return _TokenRegistry.Contract.IsTokenRegistered(&_TokenRegistry.CallOpts, _ticker)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRegistry *TokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRegistry *TokenRegistrySession) Owner() (common.Address, error) {
	return _TokenRegistry.Contract.Owner(&_TokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) Owner() (common.Address, error) {
	return _TokenRegistry.Contract.Owner(&_TokenRegistry.CallOpts)
}

// TokenAddresses is a free data retrieval call binding the contract method 0x935b13f6.
//
// Solidity: function tokenAddresses(string ) view returns(address)
func (_TokenRegistry *TokenRegistryCaller) TokenAddresses(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddresses is a free data retrieval call binding the contract method 0x935b13f6.
//
// Solidity: function tokenAddresses(string ) view returns(address)
func (_TokenRegistry *TokenRegistrySession) TokenAddresses(arg0 string) (common.Address, error) {
	return _TokenRegistry.Contract.TokenAddresses(&_TokenRegistry.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0x935b13f6.
//
// Solidity: function tokenAddresses(string ) view returns(address)
func (_TokenRegistry *TokenRegistryCallerSession) TokenAddresses(arg0 string) (common.Address, error) {
	return _TokenRegistry.Contract.TokenAddresses(&_TokenRegistry.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xb0b22c2c.
//
// Solidity: function addToken(string _ticker, address _tokenAddress) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddToken(opts *bind.TransactOpts, _ticker string, _tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addToken", _ticker, _tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0xb0b22c2c.
//
// Solidity: function addToken(string _ticker, address _tokenAddress) returns()
func (_TokenRegistry *TokenRegistrySession) AddToken(_ticker string, _tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddToken(&_TokenRegistry.TransactOpts, _ticker, _tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0xb0b22c2c.
//
// Solidity: function addToken(string _ticker, address _tokenAddress) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddToken(_ticker string, _tokenAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddToken(&_TokenRegistry.TransactOpts, _ticker, _tokenAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x1f455922.
//
// Solidity: function removeToken(string _ticker) returns()
func (_TokenRegistry *TokenRegistryTransactor) RemoveToken(opts *bind.TransactOpts, _ticker string) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "removeToken", _ticker)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x1f455922.
//
// Solidity: function removeToken(string _ticker) returns()
func (_TokenRegistry *TokenRegistrySession) RemoveToken(_ticker string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveToken(&_TokenRegistry.TransactOpts, _ticker)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x1f455922.
//
// Solidity: function removeToken(string _ticker) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) RemoveToken(_ticker string) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveToken(&_TokenRegistry.TransactOpts, _ticker)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_TokenRegistry *TokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_TokenRegistry *TokenRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferOwnership(&_TokenRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferOwnership(&_TokenRegistry.TransactOpts, _newOwner)
}

// UpdateTokenAddress is a paid mutator transaction binding the contract method 0x2efc7a94.
//
// Solidity: function updateTokenAddress(string _ticker, address _newTokenAddress) returns()
func (_TokenRegistry *TokenRegistryTransactor) UpdateTokenAddress(opts *bind.TransactOpts, _ticker string, _newTokenAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "updateTokenAddress", _ticker, _newTokenAddress)
}

// UpdateTokenAddress is a paid mutator transaction binding the contract method 0x2efc7a94.
//
// Solidity: function updateTokenAddress(string _ticker, address _newTokenAddress) returns()
func (_TokenRegistry *TokenRegistrySession) UpdateTokenAddress(_ticker string, _newTokenAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.UpdateTokenAddress(&_TokenRegistry.TransactOpts, _ticker, _newTokenAddress)
}

// UpdateTokenAddress is a paid mutator transaction binding the contract method 0x2efc7a94.
//
// Solidity: function updateTokenAddress(string _ticker, address _newTokenAddress) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) UpdateTokenAddress(_ticker string, _newTokenAddress common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.UpdateTokenAddress(&_TokenRegistry.TransactOpts, _ticker, _newTokenAddress)
}

// TokenRegistryTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the TokenRegistry contract.
type TokenRegistryTokenAddedIterator struct {
	Event *TokenRegistryTokenAdded // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenAdded)
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
		it.Event = new(TokenRegistryTokenAdded)
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
func (it *TokenRegistryTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenAdded represents a TokenAdded event raised by the TokenRegistry contract.
type TokenRegistryTokenAdded struct {
	Ticker       string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x6d97a31531bdb3c43d920bc996ec8b9fd7733f5f55a81cd74700f88631fc6032.
//
// Solidity: event TokenAdded(string ticker, address tokenAddress)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*TokenRegistryTokenAddedIterator, error) {

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenAddedIterator{contract: _TokenRegistry.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x6d97a31531bdb3c43d920bc996ec8b9fd7733f5f55a81cd74700f88631fc6032.
//
// Solidity: event TokenAdded(string ticker, address tokenAddress)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenAdded) (event.Subscription, error) {

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenAdded)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x6d97a31531bdb3c43d920bc996ec8b9fd7733f5f55a81cd74700f88631fc6032.
//
// Solidity: event TokenAdded(string ticker, address tokenAddress)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenAdded(log types.Log) (*TokenRegistryTokenAdded, error) {
	event := new(TokenRegistryTokenAdded)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegistryTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the TokenRegistry contract.
type TokenRegistryTokenRemovedIterator struct {
	Event *TokenRegistryTokenRemoved // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenRemoved)
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
		it.Event = new(TokenRegistryTokenRemoved)
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
func (it *TokenRegistryTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenRemoved represents a TokenRemoved event raised by the TokenRegistry contract.
type TokenRegistryTokenRemoved struct {
	Ticker string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0xcdd02d37d444f381ccea186dcdaf556989dcdeeebafc5cd307fe5e7609cd88c3.
//
// Solidity: event TokenRemoved(string ticker)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*TokenRegistryTokenRemovedIterator, error) {

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenRemovedIterator{contract: _TokenRegistry.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0xcdd02d37d444f381ccea186dcdaf556989dcdeeebafc5cd307fe5e7609cd88c3.
//
// Solidity: event TokenRemoved(string ticker)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenRemoved)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0xcdd02d37d444f381ccea186dcdaf556989dcdeeebafc5cd307fe5e7609cd88c3.
//
// Solidity: event TokenRemoved(string ticker)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenRemoved(log types.Log) (*TokenRegistryTokenRemoved, error) {
	event := new(TokenRegistryTokenRemoved)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRegistryTokenUpdatedIterator is returned from FilterTokenUpdated and is used to iterate over the raw logs and unpacked data for TokenUpdated events raised by the TokenRegistry contract.
type TokenRegistryTokenUpdatedIterator struct {
	Event *TokenRegistryTokenUpdated // Event containing the contract specifics and raw log

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
func (it *TokenRegistryTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryTokenUpdated)
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
		it.Event = new(TokenRegistryTokenUpdated)
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
func (it *TokenRegistryTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryTokenUpdated represents a TokenUpdated event raised by the TokenRegistry contract.
type TokenRegistryTokenUpdated struct {
	Ticker          string
	NewTokenAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenUpdated is a free log retrieval operation binding the contract event 0xab9b2a4b5369b2ede85df0177e1308b7a6e3575b7b9965d79f3a32f97510e597.
//
// Solidity: event TokenUpdated(string ticker, address newTokenAddress)
func (_TokenRegistry *TokenRegistryFilterer) FilterTokenUpdated(opts *bind.FilterOpts) (*TokenRegistryTokenUpdatedIterator, error) {

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "TokenUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTokenUpdatedIterator{contract: _TokenRegistry.contract, event: "TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenUpdated is a free log subscription operation binding the contract event 0xab9b2a4b5369b2ede85df0177e1308b7a6e3575b7b9965d79f3a32f97510e597.
//
// Solidity: event TokenUpdated(string ticker, address newTokenAddress)
func (_TokenRegistry *TokenRegistryFilterer) WatchTokenUpdated(opts *bind.WatchOpts, sink chan<- *TokenRegistryTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "TokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryTokenUpdated)
				if err := _TokenRegistry.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
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

// ParseTokenUpdated is a log parse operation binding the contract event 0xab9b2a4b5369b2ede85df0177e1308b7a6e3575b7b9965d79f3a32f97510e597.
//
// Solidity: event TokenUpdated(string ticker, address newTokenAddress)
func (_TokenRegistry *TokenRegistryFilterer) ParseTokenUpdated(log types.Log) (*TokenRegistryTokenUpdated, error) {
	event := new(TokenRegistryTokenUpdated)
	if err := _TokenRegistry.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
