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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"NodeProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NodeProviderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"NodeProviderUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_paymentsAccepted\",\"type\":\"string[]\"}],\"name\":\"addNodeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToProviderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeProvider\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"}],\"name\":\"getNodeProviderByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextProviderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"removeNodeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setNodeProviderStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractITokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600255348015610014575f80fd5b50604051611a97380380611a97833981810160405281019061003691906100da565b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610105565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a982610080565b9050919050565b6100b98161009f565b81146100c3575f80fd5b50565b5f815190506100d4816100b0565b92915050565b5f602082840312156100ef576100ee61007c565b5b5f6100fc848285016100c6565b91505092915050565b611985806101125f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80639d23c4c7116100645780639d23c4c714610135578063c40a2c2e14610153578063e835440e1461016f578063eadbc5891461019f578063f642bcda146101d257610091565b806351a8dbcf1461009557806368efbce3146100b35780638ad53763146100e6578063915e551114610102575b5f80fd5b61009d6101ee565b6040516100aa9190610cbf565b60405180910390f35b6100cd60048036038101906100c89190610d43565b6101f4565b6040516100dd9493929190610efb565b60405180910390f35b61010060048036038101906100fb9190610fa0565b610415565b005b61011c60048036038101906101179190610fde565b6104d0565b60405161012c9493929190611018565b60405180910390f35b61013d6106e2565b60405161014a91906110c4565b60405180910390f35b61016d600480360381019061016891906112eb565b610707565b005b61018960048036038101906101849190610d43565b610951565b6040516101969190610cbf565b60405180910390f35b6101b960048036038101906101b49190610fde565b610966565b6040516101c99493929190611373565b60405180910390f35b6101ec60048036038101906101e79190610fde565b610a42565b005b60025481565b5f6060805f8060015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f810361027c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102739061142d565b60405180910390fd5b5f805f8381526020019081526020015f209050805f01548160010182600301836004015f9054906101000a900460ff168280546102b890611478565b80601f01602080910402602001604051908101604052809291908181526020018280546102e490611478565b801561032f5780601f106103065761010080835404028352916020019161032f565b820191905f5260205f20905b81548152906001019060200180831161031257829003601f168201915b5050505050925081805480602002602001604051908101604052809291908181526020015f905b828210156103fe578382905f5260205f2001805461037390611478565b80601f016020809104026020016040519081016040528092919081815260200182805461039f90611478565b80156103ea5780601f106103c1576101008083540402835291602001916103ea565b820191905f5260205f20905b8154815290600101906020018083116103cd57829003601f168201915b505050505081526020019060010190610356565b505050509150955095509550955050509193509193565b5f805f8481526020019081526020015f205f015403610469576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610460906114f2565b60405180910390fd5b805f808481526020019081526020015f206004015f6101000a81548160ff0219169083151502179055507f8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce82826040516104c4929190611510565b60405180910390a15050565b60605f60605f805f808781526020019081526020015f205f01540361052a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610521906114f2565b60405180910390fd5b5f805f8781526020019081526020015f20905080600101816002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682600301836004015f9054906101000a900460ff1683805461058690611478565b80601f01602080910402602001604051908101604052809291908181526020018280546105b290611478565b80156105fd5780601f106105d4576101008083540402835291602001916105fd565b820191905f5260205f20905b8154815290600101906020018083116105e057829003601f168201915b5050505050935081805480602002602001604051908101604052809291908181526020015f905b828210156106cc578382905f5260205f2001805461064190611478565b80601f016020809104026020016040519081016040528092919081815260200182805461066d90611478565b80156106b85780601f1061068f576101008083540402835291602001916106b8565b820191905f5260205f20905b81548152906001019060200180831161069b57829003601f168201915b505050505081526020019060010190610624565b5050505091509450945094509450509193509193565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5b81518110156108115760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322b9e3b683838151811061076257610761611537565b5b60200260200101516040518263ffffffff1660e01b81526004016107869190611564565b602060405180830381865afa1580156107a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107c59190611598565b610804576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fb90611633565b60405180910390fd5b8080600101915050610709565b505f60025f8154809291906108259061167e565b9190505590505f805f8381526020019081526020015f20905081815f0181905550848160010190816108579190611859565b5083816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828160030190805190602001906108b2929190610bb7565b506001816004015f6101000a81548160ff0219169083151502179055508160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507f328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca8285604051610942929190611928565b60405180910390a15050505050565b6001602052805f5260405f205f915090505481565b5f602052805f5260405f205f91509050805f01549080600101805461098a90611478565b80601f01602080910402602001604051908101604052809291908181526020018280546109b690611478565b8015610a015780601f106109d857610100808354040283529160200191610a01565b820191905f5260205f20905b8154815290600101906020018083116109e457829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004015f9054906101000a900460ff16905084565b5f805f8381526020019081526020015f205f015403610a96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8d906114f2565b60405180910390fd5b5f805f8381526020019081526020015f206002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f90555f808381526020019081526020015f205f8082015f9055600182015f610b329190610c0e565b600282015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382015f610b679190610c4b565b600482015f6101000a81549060ff021916905550507fb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d982604051610bab9190610cbf565b60405180910390a15050565b828054828255905f5260205f20908101928215610bfd579160200282015b82811115610bfc578251829081610bec9190611859565b5091602001919060010190610bd5565b5b509050610c0a9190610c69565b5090565b508054610c1a90611478565b5f825580601f10610c2b5750610c48565b601f0160209004905f5260205f2090810190610c479190610c8c565b5b50565b5080545f8255905f5260205f2090810190610c669190610c69565b50565b5b80821115610c88575f8181610c7f9190610c0e565b50600101610c6a565b5090565b5b80821115610ca3575f815f905550600101610c8d565b5090565b5f819050919050565b610cb981610ca7565b82525050565b5f602082019050610cd25f830184610cb0565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d1282610ce9565b9050919050565b610d2281610d08565b8114610d2c575f80fd5b50565b5f81359050610d3d81610d19565b92915050565b5f60208284031215610d5857610d57610ce1565b5b5f610d6584828501610d2f565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610db082610d6e565b610dba8185610d78565b9350610dca818560208601610d88565b610dd381610d96565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f610e2182610d6e565b610e2b8185610e07565b9350610e3b818560208601610d88565b610e4481610d96565b840191505092915050565b5f610e5a8383610e17565b905092915050565b5f602082019050919050565b5f610e7882610dde565b610e828185610de8565b935083602082028501610e9485610df8565b805f5b85811015610ecf5784840389528151610eb08582610e4f565b9450610ebb83610e62565b925060208a01995050600181019050610e97565b50829750879550505050505092915050565b5f8115159050919050565b610ef581610ee1565b82525050565b5f608082019050610f0e5f830187610cb0565b8181036020830152610f208186610da6565b90508181036040830152610f348185610e6e565b9050610f436060830184610eec565b95945050505050565b610f5581610ca7565b8114610f5f575f80fd5b50565b5f81359050610f7081610f4c565b92915050565b610f7f81610ee1565b8114610f89575f80fd5b50565b5f81359050610f9a81610f76565b92915050565b5f8060408385031215610fb657610fb5610ce1565b5b5f610fc385828601610f62565b9250506020610fd485828601610f8c565b9150509250929050565b5f60208284031215610ff357610ff2610ce1565b5b5f61100084828501610f62565b91505092915050565b61101281610d08565b82525050565b5f6080820190508181035f8301526110308187610da6565b905061103f6020830186611009565b81810360408301526110518185610e6e565b90506110606060830184610eec565b95945050505050565b5f819050919050565b5f61108c61108761108284610ce9565b611069565b610ce9565b9050919050565b5f61109d82611072565b9050919050565b5f6110ae82611093565b9050919050565b6110be816110a4565b82525050565b5f6020820190506110d75f8301846110b5565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61111b82610d96565b810181811067ffffffffffffffff8211171561113a576111396110e5565b5b80604052505050565b5f61114c610cd8565b90506111588282611112565b919050565b5f67ffffffffffffffff821115611177576111766110e5565b5b61118082610d96565b9050602081019050919050565b828183375f83830152505050565b5f6111ad6111a88461115d565b611143565b9050828152602081018484840111156111c9576111c86110e1565b5b6111d484828561118d565b509392505050565b5f82601f8301126111f0576111ef6110dd565b5b813561120084826020860161119b565b91505092915050565b5f67ffffffffffffffff821115611223576112226110e5565b5b602082029050602081019050919050565b5f80fd5b5f61124a61124584611209565b611143565b9050808382526020820190506020840283018581111561126d5761126c611234565b5b835b818110156112b457803567ffffffffffffffff811115611292576112916110dd565b5b80860161129f89826111dc565b8552602085019450505060208101905061126f565b5050509392505050565b5f82601f8301126112d2576112d16110dd565b5b81356112e2848260208601611238565b91505092915050565b5f805f6060848603121561130257611301610ce1565b5b5f84013567ffffffffffffffff81111561131f5761131e610ce5565b5b61132b868287016111dc565b935050602061133c86828701610d2f565b925050604084013567ffffffffffffffff81111561135d5761135c610ce5565b5b611369868287016112be565b9150509250925092565b5f6080820190506113865f830187610cb0565b81810360208301526113988186610da6565b90506113a76040830185611009565b6113b46060830184610eec565b95945050505050565b7f4e6f646550726f766964657220776974682074686520676976656e20616464725f8201527f65737320646f6573206e6f742065786973742e00000000000000000000000000602082015250565b5f611417603383610d78565b9150611422826113bd565b604082019050919050565b5f6020820190508181035f8301526114448161140b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061148f57607f821691505b6020821081036114a2576114a161144b565b5b50919050565b7f4e6f646550726f766964657220646f6573206e6f742065786973742e000000005f82015250565b5f6114dc601c83610d78565b91506114e7826114a8565b602082019050919050565b5f6020820190508181035f830152611509816114d0565b9050919050565b5f6040820190506115235f830185610cb0565b6115306020830184610eec565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020820190508181035f83015261157c8184610da6565b905092915050565b5f8151905061159281610f76565b92915050565b5f602082840312156115ad576115ac610ce1565b5b5f6115ba84828501611584565b91505092915050565b7f5061796d656e74206d6574686f64206e6f7420726567697374657265642061735f8201527f206120746f6b656e2e0000000000000000000000000000000000000000000000602082015250565b5f61161d602983610d78565b9150611628826115c3565b604082019050919050565b5f6020820190508181035f83015261164a81611611565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61168882610ca7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116ba576116b9611651565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117217fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826116e6565b61172b86836116e6565b95508019841693508086168417925050509392505050565b5f61175d61175861175384610ca7565b611069565b610ca7565b9050919050565b5f819050919050565b61177683611743565b61178a61178282611764565b8484546116f2565b825550505050565b5f90565b61179e611792565b6117a981848461176d565b505050565b5b818110156117cc576117c15f82611796565b6001810190506117af565b5050565b601f821115611811576117e2816116c5565b6117eb846116d7565b810160208510156117fa578190505b61180e611806856116d7565b8301826117ae565b50505b505050565b5f82821c905092915050565b5f6118315f1984600802611816565b1980831691505092915050565b5f6118498383611822565b9150826002028217905092915050565b61186282610d6e565b67ffffffffffffffff81111561187b5761187a6110e5565b5b6118858254611478565b6118908282856117d0565b5f60209050601f8311600181146118c1575f84156118af578287015190505b6118b9858261183e565b865550611920565b601f1984166118cf866116c5565b5f5b828110156118f6578489015182556001820191506020850194506020810190506118d1565b86831015611913578489015161190f601f891682611822565b8355505b6001600288020188555050505b505050505050565b5f60408201905061193b5f830185610cb0565b6119486020830184611009565b939250505056fea2646970667358221220ca74aecf5870eea63fd3fd60d0c5aee8fd2b399e251f83ae9bc7e195abf8351464736f6c63430008190033",
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

// AddressToProviderId is a free data retrieval call binding the contract method 0xe835440e.
//
// Solidity: function addressToProviderId(address ) view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) AddressToProviderId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "addressToProviderId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToProviderId is a free data retrieval call binding the contract method 0xe835440e.
//
// Solidity: function addressToProviderId(address ) view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistrySession) AddressToProviderId(arg0 common.Address) (*big.Int, error) {
	return _NodeProviderRegistry.Contract.AddressToProviderId(&_NodeProviderRegistry.CallOpts, arg0)
}

// AddressToProviderId is a free data retrieval call binding the contract method 0xe835440e.
//
// Solidity: function addressToProviderId(address ) view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) AddressToProviderId(arg0 common.Address) (*big.Int, error) {
	return _NodeProviderRegistry.Contract.AddressToProviderId(&_NodeProviderRegistry.CallOpts, arg0)
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

// GetNodeProviderByAddress is a free data retrieval call binding the contract method 0x68efbce3.
//
// Solidity: function getNodeProviderByAddress(address _walletAddress) view returns(uint256, string, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) GetNodeProviderByAddress(opts *bind.CallOpts, _walletAddress common.Address) (*big.Int, string, []string, bool, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "getNodeProviderByAddress", _walletAddress)

	if err != nil {
		return *new(*big.Int), *new(string), *new([]string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([]string)).(*[]string)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetNodeProviderByAddress is a free data retrieval call binding the contract method 0x68efbce3.
//
// Solidity: function getNodeProviderByAddress(address _walletAddress) view returns(uint256, string, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistrySession) GetNodeProviderByAddress(_walletAddress common.Address) (*big.Int, string, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProviderByAddress(&_NodeProviderRegistry.CallOpts, _walletAddress)
}

// GetNodeProviderByAddress is a free data retrieval call binding the contract method 0x68efbce3.
//
// Solidity: function getNodeProviderByAddress(address _walletAddress) view returns(uint256, string, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) GetNodeProviderByAddress(_walletAddress common.Address) (*big.Int, string, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProviderByAddress(&_NodeProviderRegistry.CallOpts, _walletAddress)
}

// NextProviderId is a free data retrieval call binding the contract method 0x51a8dbcf.
//
// Solidity: function nextProviderId() view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) NextProviderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "nextProviderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextProviderId is a free data retrieval call binding the contract method 0x51a8dbcf.
//
// Solidity: function nextProviderId() view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistrySession) NextProviderId() (*big.Int, error) {
	return _NodeProviderRegistry.Contract.NextProviderId(&_NodeProviderRegistry.CallOpts)
}

// NextProviderId is a free data retrieval call binding the contract method 0x51a8dbcf.
//
// Solidity: function nextProviderId() view returns(uint256)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) NextProviderId() (*big.Int, error) {
	return _NodeProviderRegistry.Contract.NextProviderId(&_NodeProviderRegistry.CallOpts)
}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 providerId, string region, address walletAddress, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) NodeProviders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProviderId    *big.Int
	Region        string
	WalletAddress common.Address
	IsActive      bool
}, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "nodeProviders", arg0)

	outstruct := new(struct {
		ProviderId    *big.Int
		Region        string
		WalletAddress common.Address
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProviderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Region = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.WalletAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 providerId, string region, address walletAddress, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistrySession) NodeProviders(arg0 *big.Int) (struct {
	ProviderId    *big.Int
	Region        string
	WalletAddress common.Address
	IsActive      bool
}, error) {
	return _NodeProviderRegistry.Contract.NodeProviders(&_NodeProviderRegistry.CallOpts, arg0)
}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 providerId, string region, address walletAddress, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) NodeProviders(arg0 *big.Int) (struct {
	ProviderId    *big.Int
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
	ProviderId    *big.Int
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNodeProviderAdded is a free log retrieval operation binding the contract event 0x328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca.
//
// Solidity: event NodeProviderAdded(uint256 providerId, address walletAddress)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) FilterNodeProviderAdded(opts *bind.FilterOpts) (*NodeProviderRegistryNodeProviderAddedIterator, error) {

	logs, sub, err := _NodeProviderRegistry.contract.FilterLogs(opts, "NodeProviderAdded")
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryNodeProviderAddedIterator{contract: _NodeProviderRegistry.contract, event: "NodeProviderAdded", logs: logs, sub: sub}, nil
}

// WatchNodeProviderAdded is a free log subscription operation binding the contract event 0x328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca.
//
// Solidity: event NodeProviderAdded(uint256 providerId, address walletAddress)
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
// Solidity: event NodeProviderAdded(uint256 providerId, address walletAddress)
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
	ProviderId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeProviderRemoved is a free log retrieval operation binding the contract event 0xb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d9.
//
// Solidity: event NodeProviderRemoved(uint256 providerId)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) FilterNodeProviderRemoved(opts *bind.FilterOpts) (*NodeProviderRegistryNodeProviderRemovedIterator, error) {

	logs, sub, err := _NodeProviderRegistry.contract.FilterLogs(opts, "NodeProviderRemoved")
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryNodeProviderRemovedIterator{contract: _NodeProviderRegistry.contract, event: "NodeProviderRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeProviderRemoved is a free log subscription operation binding the contract event 0xb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d9.
//
// Solidity: event NodeProviderRemoved(uint256 providerId)
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
// Solidity: event NodeProviderRemoved(uint256 providerId)
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
	ProviderId *big.Int
	IsActive   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeProviderUpdated is a free log retrieval operation binding the contract event 0x8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce.
//
// Solidity: event NodeProviderUpdated(uint256 providerId, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) FilterNodeProviderUpdated(opts *bind.FilterOpts) (*NodeProviderRegistryNodeProviderUpdatedIterator, error) {

	logs, sub, err := _NodeProviderRegistry.contract.FilterLogs(opts, "NodeProviderUpdated")
	if err != nil {
		return nil, err
	}
	return &NodeProviderRegistryNodeProviderUpdatedIterator{contract: _NodeProviderRegistry.contract, event: "NodeProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeProviderUpdated is a free log subscription operation binding the contract event 0x8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce.
//
// Solidity: event NodeProviderUpdated(uint256 providerId, bool isActive)
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
// Solidity: event NodeProviderUpdated(uint256 providerId, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryFilterer) ParseNodeProviderUpdated(log types.Log) (*NodeProviderRegistryNodeProviderUpdated, error) {
	event := new(NodeProviderRegistryNodeProviderUpdated)
	if err := _NodeProviderRegistry.contract.UnpackLog(event, "NodeProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
