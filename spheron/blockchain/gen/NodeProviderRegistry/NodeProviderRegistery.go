// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NodeProviderRegistry

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

// NodeProviderRegistryMetaData contains all meta data concerning the NodeProviderRegistry contract.
var NodeProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"NodeProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"NodeProviderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"NodeProviderUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_paymentsAccepted\",\"type\":\"string[]\"}],\"name\":\"addNodeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeProvider\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"removeNodeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setNodeProviderStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractITokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260018055348015610013575f80fd5b50604051611624380380611624833981810160405281019061003591906100d9565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610104565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a88261007f565b9050919050565b6100b88161009e565b81146100c2575f80fd5b50565b5f815190506100d3816100af565b92915050565b5f602082840312156100ee576100ed61007b565b5b5f6100fb848285016100c5565b91505092915050565b611513806101115f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c8063c40a2c2e11610059578063c40a2c2e146100ec578063eadbc58914610108578063f642bcda1461013b578063f7c09189146101575761007b565b80638ad537631461007f578063915e55111461009b5780639d23c4c7146100ce575b5f80fd5b610099600480360381019061009491906109b8565b610175565b005b6100b560048036038101906100b091906109f6565b610230565b6040516100c59493929190610be2565b60405180910390f35b6100d6610442565b6040516100e39190610c8e565b60405180910390f35b61010660048036038101906101019190610edf565b610467565b005b610122600480360381019061011d91906109f6565b61066f565b6040516101329493929190610f76565b60405180910390f35b610155600480360381019061015091906109f6565b61074b565b005b61015f610849565b60405161016c9190610fc0565b60405180910390f35b5f805f8481526020019081526020015f205f0154036101c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c090611023565b60405180910390fd5b805f808481526020019081526020015f206004015f6101000a81548160ff0219169083151502179055507f8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce8282604051610224929190611041565b60405180910390a15050565b60605f60605f805f808781526020019081526020015f205f01540361028a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028190611023565b60405180910390fd5b5f805f8781526020019081526020015f20905080600101816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682600301836004015f9054906101000a900460ff168380546102e690611095565b80601f016020809104026020016040519081016040528092919081815260200182805461031290611095565b801561035d5780601f106103345761010080835404028352916020019161035d565b820191905f5260205f20905b81548152906001019060200180831161034057829003601f168201915b5050505050935081805480602002602001604051908101604052809291908181526020015f905b8282101561042c578382905f5260205f200180546103a190611095565b80601f01602080910402602001604051908101604052809291908181526020018280546103cd90611095565b80156104185780601f106103ef57610100808354040283529160200191610418565b820191905f5260205f20905b8154815290600101906020018083116103fb57829003601f168201915b505050505081526020019060010190610384565b5050505091509450945094509450509193509193565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5b81518110156105715760025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322b9e3b68383815181106104c2576104c16110c5565b5b60200260200101516040518263ffffffff1660e01b81526004016104e691906110f2565b602060405180830381865afa158015610501573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105259190611126565b610564576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055b906111c1565b60405180910390fd5b8080600101915050610469565b505f60015f8154809291906105859061120c565b9190505590505f805f8381526020019081526020015f20905081815f0181905550848160010190816105b791906113e7565b5083816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508281600301908051906020019061061292919061084f565b506001816004015f6101000a81548160ff0219169083151502179055507f328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca82856040516106609291906114b6565b60405180910390a15050505050565b5f602052805f5260405f205f91509050805f01549080600101805461069390611095565b80601f01602080910402602001604051908101604052809291908181526020018280546106bf90611095565b801561070a5780601f106106e15761010080835404028352916020019161070a565b820191905f5260205f20905b8154815290600101906020018083116106ed57829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004015f9054906101000a900460ff16905084565b5f805f8381526020019081526020015f205f01540361079f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079690611023565b60405180910390fd5b5f808281526020019081526020015f205f8082015f9055600182015f6107c591906108a6565b600282015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382015f6107fa91906108e3565b600482015f6101000a81549060ff021916905550507fb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d98160405161083e9190610fc0565b60405180910390a150565b60015481565b828054828255905f5260205f20908101928215610895579160200282015b8281111561089457825182908161088491906113e7565b509160200191906001019061086d565b5b5090506108a29190610901565b5090565b5080546108b290611095565b5f825580601f106108c357506108e0565b601f0160209004905f5260205f20908101906108df9190610924565b5b50565b5080545f8255905f5260205f20908101906108fe9190610901565b50565b5b80821115610920575f818161091791906108a6565b50600101610902565b5090565b5b8082111561093b575f815f905550600101610925565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61096281610950565b811461096c575f80fd5b50565b5f8135905061097d81610959565b92915050565b5f8115159050919050565b61099781610983565b81146109a1575f80fd5b50565b5f813590506109b28161098e565b92915050565b5f80604083850312156109ce576109cd610948565b5b5f6109db8582860161096f565b92505060206109ec858286016109a4565b9150509250929050565b5f60208284031215610a0b57610a0a610948565b5b5f610a188482850161096f565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610a6382610a21565b610a6d8185610a2b565b9350610a7d818560208601610a3b565b610a8681610a49565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610aba82610a91565b9050919050565b610aca81610ab0565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f610b1382610a21565b610b1d8185610af9565b9350610b2d818560208601610a3b565b610b3681610a49565b840191505092915050565b5f610b4c8383610b09565b905092915050565b5f602082019050919050565b5f610b6a82610ad0565b610b748185610ada565b935083602082028501610b8685610aea565b805f5b85811015610bc15784840389528151610ba28582610b41565b9450610bad83610b54565b925060208a01995050600181019050610b89565b50829750879550505050505092915050565b610bdc81610983565b82525050565b5f6080820190508181035f830152610bfa8187610a59565b9050610c096020830186610ac1565b8181036040830152610c1b8185610b60565b9050610c2a6060830184610bd3565b95945050505050565b5f819050919050565b5f610c56610c51610c4c84610a91565b610c33565b610a91565b9050919050565b5f610c6782610c3c565b9050919050565b5f610c7882610c5d565b9050919050565b610c8881610c6e565b82525050565b5f602082019050610ca15f830184610c7f565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ce582610a49565b810181811067ffffffffffffffff82111715610d0457610d03610caf565b5b80604052505050565b5f610d1661093f565b9050610d228282610cdc565b919050565b5f67ffffffffffffffff821115610d4157610d40610caf565b5b610d4a82610a49565b9050602081019050919050565b828183375f83830152505050565b5f610d77610d7284610d27565b610d0d565b905082815260208101848484011115610d9357610d92610cab565b5b610d9e848285610d57565b509392505050565b5f82601f830112610dba57610db9610ca7565b5b8135610dca848260208601610d65565b91505092915050565b610ddc81610ab0565b8114610de6575f80fd5b50565b5f81359050610df781610dd3565b92915050565b5f67ffffffffffffffff821115610e1757610e16610caf565b5b602082029050602081019050919050565b5f80fd5b5f610e3e610e3984610dfd565b610d0d565b90508083825260208201905060208402830185811115610e6157610e60610e28565b5b835b81811015610ea857803567ffffffffffffffff811115610e8657610e85610ca7565b5b808601610e938982610da6565b85526020850194505050602081019050610e63565b5050509392505050565b5f82601f830112610ec657610ec5610ca7565b5b8135610ed6848260208601610e2c565b91505092915050565b5f805f60608486031215610ef657610ef5610948565b5b5f84013567ffffffffffffffff811115610f1357610f1261094c565b5b610f1f86828701610da6565b9350506020610f3086828701610de9565b925050604084013567ffffffffffffffff811115610f5157610f5061094c565b5b610f5d86828701610eb2565b9150509250925092565b610f7081610950565b82525050565b5f608082019050610f895f830187610f67565b8181036020830152610f9b8186610a59565b9050610faa6040830185610ac1565b610fb76060830184610bd3565b95945050505050565b5f602082019050610fd35f830184610f67565b92915050565b7f4e6f646550726f766964657220646f6573206e6f742065786973742e000000005f82015250565b5f61100d601c83610a2b565b915061101882610fd9565b602082019050919050565b5f6020820190508181035f83015261103a81611001565b9050919050565b5f6040820190506110545f830185610f67565b6110616020830184610bd3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806110ac57607f821691505b6020821081036110bf576110be611068565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020820190508181035f83015261110a8184610a59565b905092915050565b5f815190506111208161098e565b92915050565b5f6020828403121561113b5761113a610948565b5b5f61114884828501611112565b91505092915050565b7f5061796d656e74206d6574686f64206e6f7420726567697374657265642061735f8201527f206120746f6b656e2e0000000000000000000000000000000000000000000000602082015250565b5f6111ab602983610a2b565b91506111b682611151565b604082019050919050565b5f6020820190508181035f8301526111d88161119f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61121682610950565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611248576112476111df565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026112af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611274565b6112b98683611274565b95508019841693508086168417925050509392505050565b5f6112eb6112e66112e184610950565b610c33565b610950565b9050919050565b5f819050919050565b611304836112d1565b611318611310826112f2565b848454611280565b825550505050565b5f90565b61132c611320565b6113378184846112fb565b505050565b5b8181101561135a5761134f5f82611324565b60018101905061133d565b5050565b601f82111561139f5761137081611253565b61137984611265565b81016020851015611388578190505b61139c61139485611265565b83018261133c565b50505b505050565b5f82821c905092915050565b5f6113bf5f19846008026113a4565b1980831691505092915050565b5f6113d783836113b0565b9150826002028217905092915050565b6113f082610a21565b67ffffffffffffffff81111561140957611408610caf565b5b6114138254611095565b61141e82828561135e565b5f60209050601f83116001811461144f575f841561143d578287015190505b61144785826113cc565b8655506114ae565b601f19841661145d86611253565b5f5b828110156114845784890151825560018201915060208501945060208101905061145f565b868310156114a1578489015161149d601f8916826113b0565b8355505b6001600288020188555050505b505050505050565b5f6040820190506114c95f830185610f67565b6114d66020830184610ac1565b939250505056fea2646970667358221220c2c3f6c654582990c996fd54fb8657db05b079a4e1163caf5442cc1c4bf8424264736f6c63430008190033",
}

// NodeProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeProviderRegistryMetaData.ABI instead.
var NodeProviderRegistryABI = NodeProviderRegistryMetaData.ABI

// NodeProviderRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NodeProviderRegistryMetaData.Bin instead.
var NodeProviderRegistryBin = NodeProviderRegistryMetaData.Bin

// DeployNodeProviderRegistry deploys a new Ethereum contract, binding an instance of NodeProviderRegistry to it.
func DeployNodeProviderRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenRegistryAddress common.Address) (common.Address, *types.Transaction, *NodeProviderRegistry, error) {
	parsed, err := NodeProviderRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NodeProviderRegistryBin), backend, _tokenRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeProviderRegistry{NodeProviderRegistryCaller: NodeProviderRegistryCaller{contract: contract}, NodeProviderRegistryTransactor: NodeProviderRegistryTransactor{contract: contract}, NodeProviderRegistryFilterer: NodeProviderRegistryFilterer{contract: contract}}, nil
}

// NodeProviderRegistry is an auto generated Go binding around an Ethereum contract.
type NodeProviderRegistry struct {
	NodeProviderRegistryCaller     // Read-only binding to the contract
	NodeProviderRegistryTransactor // Write-only binding to the contract
	NodeProviderRegistryFilterer   // Log filterer for contract events
}

// NodeProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeProviderRegistrySession struct {
	Contract     *NodeProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NodeProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeProviderRegistryCallerSession struct {
	Contract *NodeProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// NodeProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeProviderRegistryTransactorSession struct {
	Contract     *NodeProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// NodeProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeProviderRegistryRaw struct {
	Contract *NodeProviderRegistry // Generic contract binding to access the raw methods on
}

// NodeProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeProviderRegistryCallerRaw struct {
	Contract *NodeProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeProviderRegistryTransactorRaw struct {
	Contract *NodeProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeProviderRegistry creates a new instance of NodeProviderRegistry, bound to a specific deployed contract.
func NewNodeProviderRegistry(address common.Address, backend bind.ContractBackend) (*NodeProviderRegistry, error) {
	contract, err := bindNodeProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistry{NodeProviderRegistryCaller: NodeProviderRegistryCaller{contract: contract}, NodeProviderRegistryTransactor: NodeProviderRegistryTransactor{contract: contract}, NodeProviderRegistryFilterer: NodeProviderRegistryFilterer{contract: contract}}, nil
}

// NewNodeProviderRegistryCaller creates a new read-only instance of NodeProviderRegistry, bound to a specific deployed contract.
func NewNodeProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*NodeProviderRegistryCaller, error) {
	contract, err := bindNodeProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryCaller{contract: contract}, nil
}

// NewNodeProviderRegistryTransactor creates a new write-only instance of NodeProviderRegistry, bound to a specific deployed contract.
func NewNodeProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeProviderRegistryTransactor, error) {
	contract, err := bindNodeProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryTransactor{contract: contract}, nil
}

// NewNodeProviderRegistryFilterer creates a new log filterer instance of NodeProviderRegistry, bound to a specific deployed contract.
func NewNodeProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeProviderRegistryFilterer, error) {
	contract, err := bindNodeProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryFilterer{contract: contract}, nil
}

// bindNodeProviderRegistry binds a generic wrapper to an already deployed contract.
func bindNodeProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeProviderRegistry *NodeProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeProviderRegistry.Contract.NodeProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeProviderRegistry *NodeProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.NodeProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeProviderRegistry *NodeProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.NodeProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeProviderRegistry *NodeProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeProviderRegistry *NodeProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeProviderRegistry *NodeProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetNodeProvider is a free data retrieval call binding the contract method 0x915e5511.
//
// Solidity: function getNodeProvider(uint256 _validatorId) view returns(string, address, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) GetNodeProvider(opts *bind.CallOpts, _validatorId *big.Int) (string, common.Address, []string, bool, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "getNodeProvider", _validatorId)

	if err != nil {
		return *new(string), *new(common.Address), *new([]string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new([]string)).(*[]string)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetNodeProvider is a free data retrieval call binding the contract method 0x915e5511.
//
// Solidity: function getNodeProvider(uint256 _validatorId) view returns(string, address, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistrySession) GetNodeProvider(_validatorId *big.Int) (string, common.Address, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProvider(&_NodeProviderRegistry.CallOpts, _validatorId)
}

// GetNodeProvider is a free data retrieval call binding the contract method 0x915e5511.
//
// Solidity: function getNodeProvider(uint256 _validatorId) view returns(string, address, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) GetNodeProvider(_validatorId *big.Int) (string, common.Address, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProvider(&_NodeProviderRegistry.CallOpts, _validatorId)
}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) NextValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "nextValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistrySession) NextValidatorId() (*big.Int, error) {
	return _NodeProviderRegistry.Contract.NextValidatorId(&_NodeProviderRegistry.CallOpts)
}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) NextValidatorId() (*big.Int, error) {
	return _NodeProviderRegistry.Contract.NextValidatorId(&_NodeProviderRegistry.CallOpts)
}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 validatorId, string region, address walletAddress, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) NodeProviders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ValidatorId   *big.Int
	Region        string
	WalletAddress common.Address
	IsActive      bool
}, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "nodeProviders", arg0)

	outstruct := new(struct {
		ValidatorId   *big.Int
		Region        string
		WalletAddress common.Address
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Region = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.WalletAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 validatorId, string region, address walletAddress, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistrySession) NodeProviders(arg0 *big.Int) (struct {
	ValidatorId   *big.Int
	Region        string
	WalletAddress common.Address
	IsActive      bool
}, error) {
	return _NodeProviderRegistry.Contract.NodeProviders(&_NodeProviderRegistry.CallOpts, arg0)
}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 validatorId, string region, address walletAddress, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) NodeProviders(arg0 *big.Int) (struct {
	ValidatorId   *big.Int
	Region        string
	WalletAddress common.Address
	IsActive      bool
}, error) {
	return _NodeProviderRegistry.Contract.NodeProviders(&_NodeProviderRegistry.CallOpts, arg0)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "tokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_NodeProviderRegistry *NodeProviderRegistrySession) TokenRegistry() (common.Address, error) {
	return _NodeProviderRegistry.Contract.TokenRegistry(&_NodeProviderRegistry.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) TokenRegistry() (common.Address, error) {
	return _NodeProviderRegistry.Contract.TokenRegistry(&_NodeProviderRegistry.CallOpts)
}

// AddNodeProvider is a paid mutator transaction binding the contract method 0xc40a2c2e.
//
// Solidity: function addNodeProvider(string _region, address _walletAddress, string[] _paymentsAccepted) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactor) AddNodeProvider(opts *bind.TransactOpts, _region string, _walletAddress common.Address, _paymentsAccepted []string) (*types.Transaction, error) {
	return _NodeProviderRegistry.contract.Transact(opts, "addNodeProvider", _region, _walletAddress, _paymentsAccepted)
}

// AddNodeProvider is a paid mutator transaction binding the contract method 0xc40a2c2e.
//
// Solidity: function addNodeProvider(string _region, address _walletAddress, string[] _paymentsAccepted) returns()
func (_NodeProviderRegistry *NodeProviderRegistrySession) AddNodeProvider(_region string, _walletAddress common.Address, _paymentsAccepted []string) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.AddNodeProvider(&_NodeProviderRegistry.TransactOpts, _region, _walletAddress, _paymentsAccepted)
}

// AddNodeProvider is a paid mutator transaction binding the contract method 0xc40a2c2e.
//
// Solidity: function addNodeProvider(string _region, address _walletAddress, string[] _paymentsAccepted) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactorSession) AddNodeProvider(_region string, _walletAddress common.Address, _paymentsAccepted []string) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.AddNodeProvider(&_NodeProviderRegistry.TransactOpts, _region, _walletAddress, _paymentsAccepted)
}

// RemoveNodeProvider is a paid mutator transaction binding the contract method 0xf642bcda.
//
// Solidity: function removeNodeProvider(uint256 _validatorId) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactor) RemoveNodeProvider(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _NodeProviderRegistry.contract.Transact(opts, "removeNodeProvider", _validatorId)
}

// RemoveNodeProvider is a paid mutator transaction binding the contract method 0xf642bcda.
//
// Solidity: function removeNodeProvider(uint256 _validatorId) returns()
func (_NodeProviderRegistry *NodeProviderRegistrySession) RemoveNodeProvider(_validatorId *big.Int) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.RemoveNodeProvider(&_NodeProviderRegistry.TransactOpts, _validatorId)
}

// RemoveNodeProvider is a paid mutator transaction binding the contract method 0xf642bcda.
//
// Solidity: function removeNodeProvider(uint256 _validatorId) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactorSession) RemoveNodeProvider(_validatorId *big.Int) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.RemoveNodeProvider(&_NodeProviderRegistry.TransactOpts, _validatorId)
}

// SetNodeProviderStatus is a paid mutator transaction binding the contract method 0x8ad53763.
//
// Solidity: function setNodeProviderStatus(uint256 _validatorId, bool _isActive) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactor) SetNodeProviderStatus(opts *bind.TransactOpts, _validatorId *big.Int, _isActive bool) (*types.Transaction, error) {
	return _NodeProviderRegistry.contract.Transact(opts, "setNodeProviderStatus", _validatorId, _isActive)
}

// SetNodeProviderStatus is a paid mutator transaction binding the contract method 0x8ad53763.
//
// Solidity: function setNodeProviderStatus(uint256 _validatorId, bool _isActive) returns()
func (_NodeProviderRegistry *NodeProviderRegistrySession) SetNodeProviderStatus(_validatorId *big.Int, _isActive bool) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.SetNodeProviderStatus(&_NodeProviderRegistry.TransactOpts, _validatorId, _isActive)
}

// SetNodeProviderStatus is a paid mutator transaction binding the contract method 0x8ad53763.
//
// Solidity: function setNodeProviderStatus(uint256 _validatorId, bool _isActive) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactorSession) SetNodeProviderStatus(_validatorId *big.Int, _isActive bool) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.SetNodeProviderStatus(&_NodeProviderRegistry.TransactOpts, _validatorId, _isActive)
}

// NodeProviderRegistryNodeProviderAddedIterator is returned from FilterNodeProviderAdded and is used to iterate over the raw logs and unpacked data for NodeProviderAdded events raised by the NodeProviderRegistry contract.
type NodeProviderRegistryNodeProviderAddedIterator struct {
	Event *NodeProviderRegistryNodeProviderAdded // Event containing the contract specifics and raw log

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
func (it *NodeProviderRegistryNodeProviderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeProviderRegistryNodeProviderAdded)
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
		it.Event = new(NodeProviderRegistryNodeProviderAdded)
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
func (it *NodeProviderRegistryNodeProviderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeProviderRegistryNodeProviderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeProviderRegistryNodeProviderAdded represents a NodeProviderAdded event raised by the NodeProviderRegistry contract.
type NodeProviderRegistryNodeProviderAdded struct {
	ValidatorId   *big.Int
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeProviderAdded is a free log retrieval operation binding the contract event 0x328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca.
//
// Solidity: event NodeProviderAdded(uint256 validatorId, address walletAddress)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) FilterNodeProviderAdded(opts *bind.FilterOpts) (*NodeProviderRegistryNodeProviderAddedIterator, error) {

	logs, sub, err := _NodeProviderRegistry.contract.FilterLogs(opts, "NodeProviderAdded")
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryNodeProviderAddedIterator{contract: _NodeProviderRegistry.contract, event: "NodeProviderAdded", logs: logs, sub: sub}, nil
}

// WatchNodeProviderAdded is a free log subscription operation binding the contract event 0x328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca.
//
// Solidity: event NodeProviderAdded(uint256 validatorId, address walletAddress)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) WatchNodeProviderAdded(opts *bind.WatchOpts, sink chan<- *NodeProviderRegistryNodeProviderAdded) (event.Subscription, error) {

	logs, sub, err := _NodeProviderRegistry.contract.WatchLogs(opts, "NodeProviderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeProviderRegistryNodeProviderAdded)
				if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderAdded", log); err != nil {
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

// ParseNodeProviderAdded is a log parse operation binding the contract event 0x328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca.
//
// Solidity: event NodeProviderAdded(uint256 validatorId, address walletAddress)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) ParseNodeProviderAdded(log types.Log) (*NodeProviderRegistryNodeProviderAdded, error) {
	event := new(NodeProviderRegistryNodeProviderAdded)
	if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeProviderRegistryNodeProviderRemovedIterator is returned from FilterNodeProviderRemoved and is used to iterate over the raw logs and unpacked data for NodeProviderRemoved events raised by the NodeProviderRegistry contract.
type NodeProviderRegistryNodeProviderRemovedIterator struct {
	Event *NodeProviderRegistryNodeProviderRemoved // Event containing the contract specifics and raw log

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
func (it *NodeProviderRegistryNodeProviderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeProviderRegistryNodeProviderRemoved)
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
		it.Event = new(NodeProviderRegistryNodeProviderRemoved)
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
func (it *NodeProviderRegistryNodeProviderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeProviderRegistryNodeProviderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeProviderRegistryNodeProviderRemoved represents a NodeProviderRemoved event raised by the NodeProviderRegistry contract.
type NodeProviderRegistryNodeProviderRemoved struct {
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeProviderRemoved is a free log retrieval operation binding the contract event 0xb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d9.
//
// Solidity: event NodeProviderRemoved(uint256 validatorId)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) FilterNodeProviderRemoved(opts *bind.FilterOpts) (*NodeProviderRegistryNodeProviderRemovedIterator, error) {

	logs, sub, err := _NodeProviderRegistry.contract.FilterLogs(opts, "NodeProviderRemoved")
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryNodeProviderRemovedIterator{contract: _NodeProviderRegistry.contract, event: "NodeProviderRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeProviderRemoved is a free log subscription operation binding the contract event 0xb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d9.
//
// Solidity: event NodeProviderRemoved(uint256 validatorId)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) WatchNodeProviderRemoved(opts *bind.WatchOpts, sink chan<- *NodeProviderRegistryNodeProviderRemoved) (event.Subscription, error) {

	logs, sub, err := _NodeProviderRegistry.contract.WatchLogs(opts, "NodeProviderRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeProviderRegistryNodeProviderRemoved)
				if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderRemoved", log); err != nil {
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

// ParseNodeProviderRemoved is a log parse operation binding the contract event 0xb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d9.
//
// Solidity: event NodeProviderRemoved(uint256 validatorId)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) ParseNodeProviderRemoved(log types.Log) (*NodeProviderRegistryNodeProviderRemoved, error) {
	event := new(NodeProviderRegistryNodeProviderRemoved)
	if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeProviderRegistryNodeProviderUpdatedIterator is returned from FilterNodeProviderUpdated and is used to iterate over the raw logs and unpacked data for NodeProviderUpdated events raised by the NodeProviderRegistry contract.
type NodeProviderRegistryNodeProviderUpdatedIterator struct {
	Event *NodeProviderRegistryNodeProviderUpdated // Event containing the contract specifics and raw log

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
func (it *NodeProviderRegistryNodeProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeProviderRegistryNodeProviderUpdated)
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
		it.Event = new(NodeProviderRegistryNodeProviderUpdated)
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
func (it *NodeProviderRegistryNodeProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeProviderRegistryNodeProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeProviderRegistryNodeProviderUpdated represents a NodeProviderUpdated event raised by the NodeProviderRegistry contract.
type NodeProviderRegistryNodeProviderUpdated struct {
	ValidatorId *big.Int
	IsActive    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeProviderUpdated is a free log retrieval operation binding the contract event 0x8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce.
//
// Solidity: event NodeProviderUpdated(uint256 validatorId, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) FilterNodeProviderUpdated(opts *bind.FilterOpts) (*NodeProviderRegistryNodeProviderUpdatedIterator, error) {

	logs, sub, err := _NodeProviderRegistry.contract.FilterLogs(opts, "NodeProviderUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryNodeProviderUpdatedIterator{contract: _NodeProviderRegistry.contract, event: "NodeProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeProviderUpdated is a free log subscription operation binding the contract event 0x8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce.
//
// Solidity: event NodeProviderUpdated(uint256 validatorId, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) WatchNodeProviderUpdated(opts *bind.WatchOpts, sink chan<- *NodeProviderRegistryNodeProviderUpdated) (event.Subscription, error) {

	logs, sub, err := _NodeProviderRegistry.contract.WatchLogs(opts, "NodeProviderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeProviderRegistryNodeProviderUpdated)
				if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderUpdated", log); err != nil {
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

// ParseNodeProviderUpdated is a log parse operation binding the contract event 0x8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce.
//
// Solidity: event NodeProviderUpdated(uint256 validatorId, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) ParseNodeProviderUpdated(log types.Log) (*NodeProviderRegistryNodeProviderUpdated, error) {
	event := new(NodeProviderRegistryNodeProviderUpdated)
	if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
