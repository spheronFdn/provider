// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OrderMatching

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

// OrderMatchingMetaData contains all meta data concerning the OrderMatching contract.
var OrderMatchingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"}],\"name\":\"OrderMatched\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_uptime\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reputation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_providerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidPrice\",\"type\":\"uint256\"}],\"name\":\"matchOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matchedOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uptime\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260016003553480156013575f80fd5b503360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610da0806100615f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c8063533a450b146100595780638da5cb5b1461008b5780639012d5a0146100a9578063a85c38ef146100c5578063b7ba16f6146100fa575b5f80fd5b610073600480360381019061006e919061054f565b610116565b60405161008293929190610589565b60405180910390f35b61009361013c565b6040516100a091906105fd565b60405180910390f35b6100c360048036038101906100be9190610752565b610161565b005b6100df60048036038101906100da919061054f565b610250565b6040516100f196959493929190610861565b60405180910390f35b610114600480360381019061010f91906108ce565b610393565b005b6001602052805f5260405f205f91509050805f0154908060010154908060020154905083565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60035f8154809291906101749061094b565b9190505590505f6040518060c00160405280838152602001888152602001878152602001868152602001858152602001848152509050805f808481526020019081526020015f205f820151815f015560208201518160010190816101d89190610b8c565b5060408201518160020190816101ee9190610b8c565b50606082015181600301556080820151816004015560a082015181600501559050507f7e82078c35b6665b9d320ebeaa6c266960fad5b802c5558cf7df60c4769af95b8260405161023f9190610c5b565b60405180910390a150505050505050565b5f602052805f5260405f205f91509050805f015490806001018054610274906109bf565b80601f01602080910402602001604051908101604052809291908181526020018280546102a0906109bf565b80156102eb5780601f106102c2576101008083540402835291602001916102eb565b820191905f5260205f20905b8154815290600101906020018083116102ce57829003601f168201915b505050505090806002018054610300906109bf565b80601f016020809104026020016040519081016040528092919081815260200182805461032c906109bf565b80156103775780601f1061034e57610100808354040283529160200191610377565b820191905f5260205f20905b81548152906001019060200180831161035a57829003601f168201915b5050505050908060030154908060040154908060050154905086565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041990610ce4565b60405180910390fd5b5f805f8581526020019081526020015f2090505f815f01540361047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047190610d4c565b60405180910390fd5b5f60405180606001604052808681526020018581526020018481525090508060015f8781526020019081526020015f205f820151815f015560208201518160010155604082015181600201559050507f153b48ba698524f3ca3fca489cf36f46233829cb766e9de774b55345a8e832bd8585856040516104fc93929190610589565b60405180910390a15050505050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61052e8161051c565b8114610538575f80fd5b50565b5f8135905061054981610525565b92915050565b5f6020828403121561056457610563610514565b5b5f6105718482850161053b565b91505092915050565b6105838161051c565b82525050565b5f60608201905061059c5f83018661057a565b6105a9602083018561057a565b6105b6604083018461057a565b949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105e7826105be565b9050919050565b6105f7816105dd565b82525050565b5f6020820190506106105f8301846105ee565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6106648261061e565b810181811067ffffffffffffffff821117156106835761068261062e565b5b80604052505050565b5f61069561050b565b90506106a1828261065b565b919050565b5f67ffffffffffffffff8211156106c0576106bf61062e565b5b6106c98261061e565b9050602081019050919050565b828183375f83830152505050565b5f6106f66106f1846106a6565b61068c565b9050828152602081018484840111156107125761071161061a565b5b61071d8482856106d6565b509392505050565b5f82601f83011261073957610738610616565b5b81356107498482602086016106e4565b91505092915050565b5f805f805f60a0868803121561076b5761076a610514565b5b5f86013567ffffffffffffffff81111561078857610787610518565b5b61079488828901610725565b955050602086013567ffffffffffffffff8111156107b5576107b4610518565b5b6107c188828901610725565b94505060406107d28882890161053b565b93505060606107e38882890161053b565b92505060806107f48882890161053b565b9150509295509295909350565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61083382610801565b61083d818561080b565b935061084d81856020860161081b565b6108568161061e565b840191505092915050565b5f60c0820190506108745f83018961057a565b81810360208301526108868188610829565b9050818103604083015261089a8187610829565b90506108a9606083018661057a565b6108b6608083018561057a565b6108c360a083018461057a565b979650505050505050565b5f805f606084860312156108e5576108e4610514565b5b5f6108f28682870161053b565b93505060206109038682870161053b565b92505060406109148682870161053b565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6109558261051c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109875761098661091e565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806109d657607f821691505b6020821081036109e9576109e8610992565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610a4b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a10565b610a558683610a10565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610a90610a8b610a868461051c565b610a6d565b61051c565b9050919050565b5f819050919050565b610aa983610a76565b610abd610ab582610a97565b848454610a1c565b825550505050565b5f90565b610ad1610ac5565b610adc818484610aa0565b505050565b5b81811015610aff57610af45f82610ac9565b600181019050610ae2565b5050565b601f821115610b4457610b15816109ef565b610b1e84610a01565b81016020851015610b2d578190505b610b41610b3985610a01565b830182610ae1565b50505b505050565b5f82821c905092915050565b5f610b645f1984600802610b49565b1980831691505092915050565b5f610b7c8383610b55565b9150826002028217905092915050565b610b9582610801565b67ffffffffffffffff811115610bae57610bad61062e565b5b610bb882546109bf565b610bc3828285610b03565b5f60209050601f831160018114610bf4575f8415610be2578287015190505b610bec8582610b71565b865550610c53565b601f198416610c02866109ef565b5f5b82811015610c2957848901518255600182019150602085019450602081019050610c04565b86831015610c465784890151610c42601f891682610b55565b8355505b6001600288020188555050505b505050505050565b5f602082019050610c6e5f83018461057a565b92915050565b7f4f6e6c79206f776e65722063616e20706572666f726d207468697320616374695f8201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b5f610cce60228361080b565b9150610cd982610c74565b604082019050919050565b5f6020820190508181035f830152610cfb81610cc2565b9050919050565b7f4f7264657220646f6573206e6f742065786973740000000000000000000000005f82015250565b5f610d3660148361080b565b9150610d4182610d02565b602082019050919050565b5f6020820190508181035f830152610d6381610d2a565b905091905056fea26469706673582212202017fe6a04aae9494c83280ad69c10d799a90a362807da9f857555bc227424b264736f6c63430008190033",
}

// OrderMatchingABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderMatchingMetaData.ABI instead.
var OrderMatchingABI = OrderMatchingMetaData.ABI

// OrderMatchingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrderMatchingMetaData.Bin instead.
var OrderMatchingBin = OrderMatchingMetaData.Bin

// DeployOrderMatching deploys a new Ethereum contract, binding an instance of OrderMatching to it.
func DeployOrderMatching(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrderMatching, error) {
	parsed, err := OrderMatchingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrderMatchingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderMatching{OrderMatchingCaller: OrderMatchingCaller{contract: contract}, OrderMatchingTransactor: OrderMatchingTransactor{contract: contract}, OrderMatchingFilterer: OrderMatchingFilterer{contract: contract}}, nil
}

// OrderMatching is an auto generated Go binding around an Ethereum contract.
type OrderMatching struct {
	OrderMatchingCaller     // Read-only binding to the contract
	OrderMatchingTransactor // Write-only binding to the contract
	OrderMatchingFilterer   // Log filterer for contract events
}

// OrderMatchingCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderMatchingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderMatchingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderMatchingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderMatchingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderMatchingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderMatchingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderMatchingSession struct {
	Contract     *OrderMatching    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderMatchingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderMatchingCallerSession struct {
	Contract *OrderMatchingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OrderMatchingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderMatchingTransactorSession struct {
	Contract     *OrderMatchingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OrderMatchingRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderMatchingRaw struct {
	Contract *OrderMatching // Generic contract binding to access the raw methods on
}

// OrderMatchingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderMatchingCallerRaw struct {
	Contract *OrderMatchingCaller // Generic read-only contract binding to access the raw methods on
}

// OrderMatchingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderMatchingTransactorRaw struct {
	Contract *OrderMatchingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderMatching creates a new instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatching(address common.Address, backend bind.ContractBackend) (*OrderMatching, error) {
	contract, err := bindOrderMatching(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderMatching{OrderMatchingCaller: OrderMatchingCaller{contract: contract}, OrderMatchingTransactor: OrderMatchingTransactor{contract: contract}, OrderMatchingFilterer: OrderMatchingFilterer{contract: contract}}, nil
}

// NewOrderMatchingCaller creates a new read-only instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatchingCaller(address common.Address, caller bind.ContractCaller) (*OrderMatchingCaller, error) {
	contract, err := bindOrderMatching(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderMatchingCaller{contract: contract}, nil
}

// NewOrderMatchingTransactor creates a new write-only instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatchingTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderMatchingTransactor, error) {
	contract, err := bindOrderMatching(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderMatchingTransactor{contract: contract}, nil
}

// NewOrderMatchingFilterer creates a new log filterer instance of OrderMatching, bound to a specific deployed contract.
func NewOrderMatchingFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderMatchingFilterer, error) {
	contract, err := bindOrderMatching(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderMatchingFilterer{contract: contract}, nil
}

// bindOrderMatching binds a generic wrapper to an already deployed contract.
func bindOrderMatching(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderMatchingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderMatching *OrderMatchingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderMatching.Contract.OrderMatchingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderMatching *OrderMatchingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderMatching.Contract.OrderMatchingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderMatching *OrderMatchingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderMatching.Contract.OrderMatchingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderMatching *OrderMatchingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderMatching.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderMatching *OrderMatchingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderMatching.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderMatching *OrderMatchingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderMatching.Contract.contract.Transact(opts, method, params...)
}

// MatchedOrders is a free data retrieval call binding the contract method 0x533a450b.
//
// Solidity: function matchedOrders(uint256 ) view returns(uint256 orderId, uint256 providerId, uint256 bidPrice)
func (_OrderMatching *OrderMatchingCaller) MatchedOrders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId    *big.Int
	ProviderId *big.Int
	BidPrice   *big.Int
}, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "matchedOrders", arg0)

	outstruct := new(struct {
		OrderId    *big.Int
		ProviderId *big.Int
		BidPrice   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProviderId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BidPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MatchedOrders is a free data retrieval call binding the contract method 0x533a450b.
//
// Solidity: function matchedOrders(uint256 ) view returns(uint256 orderId, uint256 providerId, uint256 bidPrice)
func (_OrderMatching *OrderMatchingSession) MatchedOrders(arg0 *big.Int) (struct {
	OrderId    *big.Int
	ProviderId *big.Int
	BidPrice   *big.Int
}, error) {
	return _OrderMatching.Contract.MatchedOrders(&_OrderMatching.CallOpts, arg0)
}

// MatchedOrders is a free data retrieval call binding the contract method 0x533a450b.
//
// Solidity: function matchedOrders(uint256 ) view returns(uint256 orderId, uint256 providerId, uint256 bidPrice)
func (_OrderMatching *OrderMatchingCallerSession) MatchedOrders(arg0 *big.Int) (struct {
	OrderId    *big.Int
	ProviderId *big.Int
	BidPrice   *big.Int
}, error) {
	return _OrderMatching.Contract.MatchedOrders(&_OrderMatching.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 orderId, string region, string uptime, uint256 reputation, uint256 slashes, uint256 maxPrice)
func (_OrderMatching *OrderMatchingCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId    *big.Int
	Region     string
	Uptime     string
	Reputation *big.Int
	Slashes    *big.Int
	MaxPrice   *big.Int
}, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		OrderId    *big.Int
		Region     string
		Uptime     string
		Reputation *big.Int
		Slashes    *big.Int
		MaxPrice   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Region = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Uptime = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Reputation = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Slashes = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 orderId, string region, string uptime, uint256 reputation, uint256 slashes, uint256 maxPrice)
func (_OrderMatching *OrderMatchingSession) Orders(arg0 *big.Int) (struct {
	OrderId    *big.Int
	Region     string
	Uptime     string
	Reputation *big.Int
	Slashes    *big.Int
	MaxPrice   *big.Int
}, error) {
	return _OrderMatching.Contract.Orders(&_OrderMatching.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(uint256 orderId, string region, string uptime, uint256 reputation, uint256 slashes, uint256 maxPrice)
func (_OrderMatching *OrderMatchingCallerSession) Orders(arg0 *big.Int) (struct {
	OrderId    *big.Int
	Region     string
	Uptime     string
	Reputation *big.Int
	Slashes    *big.Int
	MaxPrice   *big.Int
}, error) {
	return _OrderMatching.Contract.Orders(&_OrderMatching.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderMatching *OrderMatchingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderMatching.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderMatching *OrderMatchingSession) Owner() (common.Address, error) {
	return _OrderMatching.Contract.Owner(&_OrderMatching.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrderMatching *OrderMatchingCallerSession) Owner() (common.Address, error) {
	return _OrderMatching.Contract.Owner(&_OrderMatching.CallOpts)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x9012d5a0.
//
// Solidity: function createOrder(string _region, string _uptime, uint256 _reputation, uint256 _slashes, uint256 _maxPrice) returns()
func (_OrderMatching *OrderMatchingTransactor) CreateOrder(opts *bind.TransactOpts, _region string, _uptime string, _reputation *big.Int, _slashes *big.Int, _maxPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "createOrder", _region, _uptime, _reputation, _slashes, _maxPrice)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x9012d5a0.
//
// Solidity: function createOrder(string _region, string _uptime, uint256 _reputation, uint256 _slashes, uint256 _maxPrice) returns()
func (_OrderMatching *OrderMatchingSession) CreateOrder(_region string, _uptime string, _reputation *big.Int, _slashes *big.Int, _maxPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.CreateOrder(&_OrderMatching.TransactOpts, _region, _uptime, _reputation, _slashes, _maxPrice)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x9012d5a0.
//
// Solidity: function createOrder(string _region, string _uptime, uint256 _reputation, uint256 _slashes, uint256 _maxPrice) returns()
func (_OrderMatching *OrderMatchingTransactorSession) CreateOrder(_region string, _uptime string, _reputation *big.Int, _slashes *big.Int, _maxPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.CreateOrder(&_OrderMatching.TransactOpts, _region, _uptime, _reputation, _slashes, _maxPrice)
}

// MatchOrder is a paid mutator transaction binding the contract method 0xb7ba16f6.
//
// Solidity: function matchOrder(uint256 _orderId, uint256 _providerId, uint256 _bidPrice) returns()
func (_OrderMatching *OrderMatchingTransactor) MatchOrder(opts *bind.TransactOpts, _orderId *big.Int, _providerId *big.Int, _bidPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.contract.Transact(opts, "matchOrder", _orderId, _providerId, _bidPrice)
}

// MatchOrder is a paid mutator transaction binding the contract method 0xb7ba16f6.
//
// Solidity: function matchOrder(uint256 _orderId, uint256 _providerId, uint256 _bidPrice) returns()
func (_OrderMatching *OrderMatchingSession) MatchOrder(_orderId *big.Int, _providerId *big.Int, _bidPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.MatchOrder(&_OrderMatching.TransactOpts, _orderId, _providerId, _bidPrice)
}

// MatchOrder is a paid mutator transaction binding the contract method 0xb7ba16f6.
//
// Solidity: function matchOrder(uint256 _orderId, uint256 _providerId, uint256 _bidPrice) returns()
func (_OrderMatching *OrderMatchingTransactorSession) MatchOrder(_orderId *big.Int, _providerId *big.Int, _bidPrice *big.Int) (*types.Transaction, error) {
	return _OrderMatching.Contract.MatchOrder(&_OrderMatching.TransactOpts, _orderId, _providerId, _bidPrice)
}

// OrderMatchingOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the OrderMatching contract.
type OrderMatchingOrderCreatedIterator struct {
	Event *OrderMatchingOrderCreated // Event containing the contract specifics and raw log

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
func (it *OrderMatchingOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingOrderCreated)
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
		it.Event = new(OrderMatchingOrderCreated)
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
func (it *OrderMatchingOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingOrderCreated represents a OrderCreated event raised by the OrderMatching contract.
type OrderMatchingOrderCreated struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x7e82078c35b6665b9d320ebeaa6c266960fad5b802c5558cf7df60c4769af95b.
//
// Solidity: event OrderCreated(uint256 orderId)
func (_OrderMatching *OrderMatchingFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*OrderMatchingOrderCreatedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingOrderCreatedIterator{contract: _OrderMatching.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x7e82078c35b6665b9d320ebeaa6c266960fad5b802c5558cf7df60c4769af95b.
//
// Solidity: event OrderCreated(uint256 orderId)
func (_OrderMatching *OrderMatchingFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *OrderMatchingOrderCreated) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingOrderCreated)
				if err := _OrderMatching.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x7e82078c35b6665b9d320ebeaa6c266960fad5b802c5558cf7df60c4769af95b.
//
// Solidity: event OrderCreated(uint256 orderId)
func (_OrderMatching *OrderMatchingFilterer) ParseOrderCreated(log types.Log) (*OrderMatchingOrderCreated, error) {
	event := new(OrderMatchingOrderCreated)
	if err := _OrderMatching.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderMatchingOrderMatchedIterator is returned from FilterOrderMatched and is used to iterate over the raw logs and unpacked data for OrderMatched events raised by the OrderMatching contract.
type OrderMatchingOrderMatchedIterator struct {
	Event *OrderMatchingOrderMatched // Event containing the contract specifics and raw log

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
func (it *OrderMatchingOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderMatchingOrderMatched)
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
		it.Event = new(OrderMatchingOrderMatched)
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
func (it *OrderMatchingOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderMatchingOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderMatchingOrderMatched represents a OrderMatched event raised by the OrderMatching contract.
type OrderMatchingOrderMatched struct {
	OrderId    *big.Int
	ProviderId *big.Int
	BidPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderMatched is a free log retrieval operation binding the contract event 0x153b48ba698524f3ca3fca489cf36f46233829cb766e9de774b55345a8e832bd.
//
// Solidity: event OrderMatched(uint256 orderId, uint256 providerId, uint256 bidPrice)
func (_OrderMatching *OrderMatchingFilterer) FilterOrderMatched(opts *bind.FilterOpts) (*OrderMatchingOrderMatchedIterator, error) {

	logs, sub, err := _OrderMatching.contract.FilterLogs(opts, "OrderMatched")
	if err != nil {
		return nil, err
	}
	return &OrderMatchingOrderMatchedIterator{contract: _OrderMatching.contract, event: "OrderMatched", logs: logs, sub: sub}, nil
}

// WatchOrderMatched is a free log subscription operation binding the contract event 0x153b48ba698524f3ca3fca489cf36f46233829cb766e9de774b55345a8e832bd.
//
// Solidity: event OrderMatched(uint256 orderId, uint256 providerId, uint256 bidPrice)
func (_OrderMatching *OrderMatchingFilterer) WatchOrderMatched(opts *bind.WatchOpts, sink chan<- *OrderMatchingOrderMatched) (event.Subscription, error) {

	logs, sub, err := _OrderMatching.contract.WatchLogs(opts, "OrderMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderMatchingOrderMatched)
				if err := _OrderMatching.contract.UnpackLog(event, "OrderMatched", log); err != nil {
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

// ParseOrderMatched is a log parse operation binding the contract event 0x153b48ba698524f3ca3fca489cf36f46233829cb766e9de774b55345a8e832bd.
//
// Solidity: event OrderMatched(uint256 orderId, uint256 providerId, uint256 bidPrice)
func (_OrderMatching *OrderMatchingFilterer) ParseOrderMatched(log types.Log) (*OrderMatchingOrderMatched, error) {
	event := new(OrderMatchingOrderMatched)
	if err := _OrderMatching.contract.UnpackLog(event, "OrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
