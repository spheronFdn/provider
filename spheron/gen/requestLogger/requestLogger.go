// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package requestLogger

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

// RequestLoggerMetaData contains all meta data concerning the RequestLogger contract.
var RequestLoggerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"request\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RequestStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_request\",\"type\":\"string\"}],\"name\":\"storeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5061081b8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80635badbe4c1461004357806381d12c5814610061578063e7770d6a14610092575b5f80fd5b61004b6100ae565b6040516100589190610257565b60405180910390f35b61007b600480360381019061007691906102a2565b6100b4565b60405161008992919061033d565b60405180910390f35b6100ac60048036038101906100a791906103cc565b610159565b005b60015481565b5f602052805f5260405f205f91509050805f0180546100d290610444565b80601f01602080910402602001604051908101604052809291908181526020018280546100fe90610444565b80156101495780601f1061012057610100808354040283529160200191610149565b820191905f5260205f20905b81548152906001019060200180831161012c57829003601f168201915b5050505050908060010154905082565b6001805f82825461016a91906104a1565b92505081905550604051806040016040528083838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001428152505f8060015481526020019081526020015f205f820151815f0190816101ee919061069e565b50602082015181600101559050507f62022416a8040e1717324cb761fed1a180353f407eea93f9161547c19e3a7b0460015483834260405161023394939291906107a7565b60405180910390a15050565b5f819050919050565b6102518161023f565b82525050565b5f60208201905061026a5f830184610248565b92915050565b5f80fd5b5f80fd5b6102818161023f565b811461028b575f80fd5b50565b5f8135905061029c81610278565b92915050565b5f602082840312156102b7576102b6610270565b5b5f6102c48482850161028e565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61030f826102cd565b61031981856102d7565b93506103298185602086016102e7565b610332816102f5565b840191505092915050565b5f6040820190508181035f8301526103558185610305565b90506103646020830184610248565b9392505050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261038c5761038b61036b565b5b8235905067ffffffffffffffff8111156103a9576103a861036f565b5b6020830191508360018202830111156103c5576103c4610373565b5b9250929050565b5f80602083850312156103e2576103e1610270565b5b5f83013567ffffffffffffffff8111156103ff576103fe610274565b5b61040b85828601610377565b92509250509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061045b57607f821691505b60208210810361046e5761046d610417565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6104ab8261023f565b91506104b68361023f565b92508282019050808211156104ce576104cd610474565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261055d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610522565b6105678683610522565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6105a261059d6105988461023f565b61057f565b61023f565b9050919050565b5f819050919050565b6105bb83610588565b6105cf6105c7826105a9565b84845461052e565b825550505050565b5f90565b6105e36105d7565b6105ee8184846105b2565b505050565b5b81811015610611576106065f826105db565b6001810190506105f4565b5050565b601f8211156106565761062781610501565b61063084610513565b8101602085101561063f578190505b61065361064b85610513565b8301826105f3565b50505b505050565b5f82821c905092915050565b5f6106765f198460080261065b565b1980831691505092915050565b5f61068e8383610667565b9150826002028217905092915050565b6106a7826102cd565b67ffffffffffffffff8111156106c0576106bf6104d4565b5b6106ca8254610444565b6106d5828285610615565b5f60209050601f831160018114610706575f84156106f4578287015190505b6106fe8582610683565b865550610765565b601f19841661071486610501565b5f5b8281101561073b57848901518255600182019150602085019450602081019050610716565b868310156107585784890151610754601f891682610667565b8355505b6001600288020188555050505b505050505050565b828183375f83830152505050565b5f61078683856102d7565b935061079383858461076d565b61079c836102f5565b840190509392505050565b5f6060820190506107ba5f830187610248565b81810360208301526107cd81858761077b565b90506107dc6040830184610248565b9594505050505056fea2646970667358221220ede9b8593463ffd29f4303263cad13fed1f9e74e413bc30887b7f298db57caeb64736f6c63430008190033",
}

// RequestLoggerABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestLoggerMetaData.ABI instead.
var RequestLoggerABI = RequestLoggerMetaData.ABI

// RequestLoggerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestLoggerMetaData.Bin instead.
var RequestLoggerBin = RequestLoggerMetaData.Bin

// DeployRequestLogger deploys a new Ethereum contract, binding an instance of RequestLogger to it.
func DeployRequestLogger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestLogger, error) {
	parsed, err := RequestLoggerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestLoggerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestLogger{RequestLoggerCaller: RequestLoggerCaller{contract: contract}, RequestLoggerTransactor: RequestLoggerTransactor{contract: contract}, RequestLoggerFilterer: RequestLoggerFilterer{contract: contract}}, nil
}

// RequestLogger is an auto generated Go binding around an Ethereum contract.
type RequestLogger struct {
	RequestLoggerCaller     // Read-only binding to the contract
	RequestLoggerTransactor // Write-only binding to the contract
	RequestLoggerFilterer   // Log filterer for contract events
}

// RequestLoggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestLoggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLoggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestLoggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLoggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestLoggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLoggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestLoggerSession struct {
	Contract     *RequestLogger    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestLoggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestLoggerCallerSession struct {
	Contract *RequestLoggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RequestLoggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestLoggerTransactorSession struct {
	Contract     *RequestLoggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RequestLoggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestLoggerRaw struct {
	Contract *RequestLogger // Generic contract binding to access the raw methods on
}

// RequestLoggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestLoggerCallerRaw struct {
	Contract *RequestLoggerCaller // Generic read-only contract binding to access the raw methods on
}

// RequestLoggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestLoggerTransactorRaw struct {
	Contract *RequestLoggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestLogger creates a new instance of RequestLogger, bound to a specific deployed contract.
func NewRequestLogger(address common.Address, backend bind.ContractBackend) (*RequestLogger, error) {
	contract, err := bindRequestLogger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestLogger{RequestLoggerCaller: RequestLoggerCaller{contract: contract}, RequestLoggerTransactor: RequestLoggerTransactor{contract: contract}, RequestLoggerFilterer: RequestLoggerFilterer{contract: contract}}, nil
}

// NewRequestLoggerCaller creates a new read-only instance of RequestLogger, bound to a specific deployed contract.
func NewRequestLoggerCaller(address common.Address, caller bind.ContractCaller) (*RequestLoggerCaller, error) {
	contract, err := bindRequestLogger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLoggerCaller{contract: contract}, nil
}

// NewRequestLoggerTransactor creates a new write-only instance of RequestLogger, bound to a specific deployed contract.
func NewRequestLoggerTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestLoggerTransactor, error) {
	contract, err := bindRequestLogger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLoggerTransactor{contract: contract}, nil
}

// NewRequestLoggerFilterer creates a new log filterer instance of RequestLogger, bound to a specific deployed contract.
func NewRequestLoggerFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestLoggerFilterer, error) {
	contract, err := bindRequestLogger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestLoggerFilterer{contract: contract}, nil
}

// bindRequestLogger binds a generic wrapper to an already deployed contract.
func bindRequestLogger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequestLoggerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLogger *RequestLoggerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLogger.Contract.RequestLoggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLogger *RequestLoggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLogger.Contract.RequestLoggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLogger *RequestLoggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLogger.Contract.RequestLoggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLogger *RequestLoggerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLogger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLogger *RequestLoggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLogger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLogger *RequestLoggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLogger.Contract.contract.Transact(opts, method, params...)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_RequestLogger *RequestLoggerCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RequestLogger.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_RequestLogger *RequestLoggerSession) RequestCount() (*big.Int, error) {
	return _RequestLogger.Contract.RequestCount(&_RequestLogger.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_RequestLogger *RequestLoggerCallerSession) RequestCount() (*big.Int, error) {
	return _RequestLogger.Contract.RequestCount(&_RequestLogger.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string data, uint256 timestamp)
func (_RequestLogger *RequestLoggerCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Data      string
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _RequestLogger.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Data      string
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string data, uint256 timestamp)
func (_RequestLogger *RequestLoggerSession) Requests(arg0 *big.Int) (struct {
	Data      string
	Timestamp *big.Int
}, error) {
	return _RequestLogger.Contract.Requests(&_RequestLogger.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string data, uint256 timestamp)
func (_RequestLogger *RequestLoggerCallerSession) Requests(arg0 *big.Int) (struct {
	Data      string
	Timestamp *big.Int
}, error) {
	return _RequestLogger.Contract.Requests(&_RequestLogger.CallOpts, arg0)
}

// StoreRequest is a paid mutator transaction binding the contract method 0xe7770d6a.
//
// Solidity: function storeRequest(string _request) returns()
func (_RequestLogger *RequestLoggerTransactor) StoreRequest(opts *bind.TransactOpts, _request string) (*types.Transaction, error) {
	return _RequestLogger.contract.Transact(opts, "storeRequest", _request)
}

// StoreRequest is a paid mutator transaction binding the contract method 0xe7770d6a.
//
// Solidity: function storeRequest(string _request) returns()
func (_RequestLogger *RequestLoggerSession) StoreRequest(_request string) (*types.Transaction, error) {
	return _RequestLogger.Contract.StoreRequest(&_RequestLogger.TransactOpts, _request)
}

// StoreRequest is a paid mutator transaction binding the contract method 0xe7770d6a.
//
// Solidity: function storeRequest(string _request) returns()
func (_RequestLogger *RequestLoggerTransactorSession) StoreRequest(_request string) (*types.Transaction, error) {
	return _RequestLogger.Contract.StoreRequest(&_RequestLogger.TransactOpts, _request)
}

// RequestLoggerRequestStoredIterator is returned from FilterRequestStored and is used to iterate over the raw logs and unpacked data for RequestStored events raised by the RequestLogger contract.
type RequestLoggerRequestStoredIterator struct {
	Event *RequestLoggerRequestStored // Event containing the contract specifics and raw log

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
func (it *RequestLoggerRequestStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestLoggerRequestStored)
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
		it.Event = new(RequestLoggerRequestStored)
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
func (it *RequestLoggerRequestStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestLoggerRequestStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestLoggerRequestStored represents a RequestStored event raised by the RequestLogger contract.
type RequestLoggerRequestStored struct {
	RequestId *big.Int
	Request   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestStored is a free log retrieval operation binding the contract event 0x62022416a8040e1717324cb761fed1a180353f407eea93f9161547c19e3a7b04.
//
// Solidity: event RequestStored(uint256 requestId, string request, uint256 timestamp)
func (_RequestLogger *RequestLoggerFilterer) FilterRequestStored(opts *bind.FilterOpts) (*RequestLoggerRequestStoredIterator, error) {

	logs, sub, err := _RequestLogger.contract.FilterLogs(opts, "RequestStored")
	if err != nil {
		return nil, err
	}
	return &RequestLoggerRequestStoredIterator{contract: _RequestLogger.contract, event: "RequestStored", logs: logs, sub: sub}, nil
}

// WatchRequestStored is a free log subscription operation binding the contract event 0x62022416a8040e1717324cb761fed1a180353f407eea93f9161547c19e3a7b04.
//
// Solidity: event RequestStored(uint256 requestId, string request, uint256 timestamp)
func (_RequestLogger *RequestLoggerFilterer) WatchRequestStored(opts *bind.WatchOpts, sink chan<- *RequestLoggerRequestStored) (event.Subscription, error) {

	logs, sub, err := _RequestLogger.contract.WatchLogs(opts, "RequestStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestLoggerRequestStored)
				if err := _RequestLogger.contract.UnpackLog(event, "RequestStored", log); err != nil {
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

// ParseRequestStored is a log parse operation binding the contract event 0x62022416a8040e1717324cb761fed1a180353f407eea93f9161547c19e3a7b04.
//
// Solidity: event RequestStored(uint256 requestId, string request, uint256 timestamp)
func (_RequestLogger *RequestLoggerFilterer) ParseRequestStored(log types.Log) (*RequestLoggerRequestStored, error) {
	event := new(RequestLoggerRequestStored)
	if err := _RequestLogger.contract.UnpackLog(event, "RequestStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
