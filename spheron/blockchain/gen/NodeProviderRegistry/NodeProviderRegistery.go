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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"NodeProviderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"}],\"name\":\"NodeProviderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"NodeProviderUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_region\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_paymentsAccepted\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_attributes\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hostUri\",\"type\":\"string\"}],\"name\":\"addNodeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToProviderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeProvider\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"}],\"name\":\"getNodeProviderByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextProviderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeProviders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"providerId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"attributes\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hostUri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"removeNodeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setNodeProviderStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractITokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600255348015610014575f80fd5b50604051611f20380380611f20833981810160405281019061003691906100da565b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610105565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a982610080565b9050919050565b6100b98161009f565b81146100c3575f80fd5b50565b5f815190506100d4816100b0565b92915050565b5f602082840312156100ef576100ee61007c565b5b5f6100fc848285016100c6565b91505092915050565b611e0e806101125f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c8063915e551111610064578063915e5511146101205780639d23c4c714610155578063e835440e14610173578063eadbc589146101a3578063f642bcda146101d857610091565b80631432d4f91461009557806351a8dbcf146100b157806368efbce3146100cf5780638ad5376314610104575b5f80fd5b6100af60048036038101906100aa91906112dc565b6101f4565b005b6100b9610464565b6040516100c691906113db565b60405180910390f35b6100e960048036038101906100e491906113f4565b61046a565b6040516100fb9695949392919061159c565b60405180910390f35b61011e6004803603810190610119919061166b565b6107b0565b005b61013a600480360381019061013591906116a9565b61086b565b60405161014c969594939291906116e3565b60405180910390f35b61015d610ba2565b60405161016a91906117b9565b60405180910390f35b61018d600480360381019061018891906113f4565b610bc7565b60405161019a91906113db565b60405180910390f35b6101bd60048036038101906101b891906116a9565b610bdc565b6040516101cf969594939291906117d2565b60405180910390f35b6101f260048036038101906101ed91906116a9565b610dd0565b005b5f5b83518110156102fe5760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322b9e3b685838151811061024f5761024e611846565b5b60200260200101516040518263ffffffff1660e01b81526004016102739190611873565b602060405180830381865afa15801561028e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b291906118a7565b6102f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e890611942565b60405180910390fd5b80806001019150506101f6565b505f60025f8154809291906103129061198d565b9190505590505f805f8381526020019081526020015f20905081815f0181905550868160010190816103449190611bc5565b50838160040190816103569190611bc5565b50828160050190816103689190611bc5565b5085816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550848160030190805190602001906103c3929190610f63565b506001816006015f6101000a81548160ff0219169083151502179055508160015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507f328912c436036c08cd0ce170bbf2f0810e172ddae3a9bb430c2120556d26a7ca8287604051610453929190611c94565b60405180910390a150505050505050565b60025481565b5f6060806060805f8060015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f81036104f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ec90611d2b565b60405180910390fd5b5f805f8381526020019081526020015f209050805f015481600101826004018360050184600301856006015f9054906101000a900460ff1684805461053990611a01565b80601f016020809104026020016040519081016040528092919081815260200182805461056590611a01565b80156105b05780601f10610587576101008083540402835291602001916105b0565b820191905f5260205f20905b81548152906001019060200180831161059357829003601f168201915b505050505094508380546105c390611a01565b80601f01602080910402602001604051908101604052809291908181526020018280546105ef90611a01565b801561063a5780601f106106115761010080835404028352916020019161063a565b820191905f5260205f20905b81548152906001019060200180831161061d57829003601f168201915b5050505050935082805461064d90611a01565b80601f016020809104026020016040519081016040528092919081815260200182805461067990611a01565b80156106c45780601f1061069b576101008083540402835291602001916106c4565b820191905f5260205f20905b8154815290600101906020018083116106a757829003601f168201915b5050505050925081805480602002602001604051908101604052809291908181526020015f905b82821015610793578382905f5260205f2001805461070890611a01565b80601f016020809104026020016040519081016040528092919081815260200182805461073490611a01565b801561077f5780601f106107565761010080835404028352916020019161077f565b820191905f5260205f20905b81548152906001019060200180831161076257829003601f168201915b5050505050815260200190600101906106eb565b505050509150975097509750975097509750505091939550919395565b5f805f8481526020019081526020015f205f015403610804576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fb90611d93565b60405180910390fd5b805f808481526020019081526020015f206006015f6101000a81548160ff0219169083151502179055507f8af896476cd9524899861a2c3aab4dc66f47fc9aa162ed963e8e20c7809525ce828260405161085f929190611db1565b60405180910390a15050565b60608060605f60605f805f808981526020019081526020015f205f0154036108c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bf90611d93565b60405180910390fd5b5f805f8981526020019081526020015f209050806001018160040182600501836002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600301856006015f9054906101000a900460ff1685805461092c90611a01565b80601f016020809104026020016040519081016040528092919081815260200182805461095890611a01565b80156109a35780601f1061097a576101008083540402835291602001916109a3565b820191905f5260205f20905b81548152906001019060200180831161098657829003601f168201915b505050505095508480546109b690611a01565b80601f01602080910402602001604051908101604052809291908181526020018280546109e290611a01565b8015610a2d5780601f10610a0457610100808354040283529160200191610a2d565b820191905f5260205f20905b815481529060010190602001808311610a1057829003601f168201915b50505050509450838054610a4090611a01565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6c90611a01565b8015610ab75780601f10610a8e57610100808354040283529160200191610ab7565b820191905f5260205f20905b815481529060010190602001808311610a9a57829003601f168201915b5050505050935081805480602002602001604051908101604052809291908181526020015f905b82821015610b86578382905f5260205f20018054610afb90611a01565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2790611a01565b8015610b725780601f10610b4957610100808354040283529160200191610b72565b820191905f5260205f20905b815481529060010190602001808311610b5557829003601f168201915b505050505081526020019060010190610ade565b5050505091509650965096509650965096505091939550919395565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052805f5260405f205f915090505481565b5f602052805f5260405f205f91509050805f015490806001018054610c0090611a01565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2c90611a01565b8015610c775780601f10610c4e57610100808354040283529160200191610c77565b820191905f5260205f20905b815481529060010190602001808311610c5a57829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004018054610cb190611a01565b80601f0160208091040260200160405190810160405280929190818152602001828054610cdd90611a01565b8015610d285780601f10610cff57610100808354040283529160200191610d28565b820191905f5260205f20905b815481529060010190602001808311610d0b57829003601f168201915b505050505090806005018054610d3d90611a01565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6990611a01565b8015610db45780601f10610d8b57610100808354040283529160200191610db4565b820191905f5260205f20905b815481529060010190602001808311610d9757829003601f168201915b505050505090806006015f9054906101000a900460ff16905086565b5f805f8381526020019081526020015f205f015403610e24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1b90611d93565b60405180910390fd5b5f805f8381526020019081526020015f206002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f90555f808381526020019081526020015f205f8082015f9055600182015f610ec09190610fba565b600282015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600382015f610ef59190610ff7565b600482015f610f049190610fba565b600582015f610f139190610fba565b600682015f6101000a81549060ff021916905550507fb102d5ae9794b7f3538bdd875291e4b485031d5d02363149e227b7f2500883d982604051610f5791906113db565b60405180910390a15050565b828054828255905f5260205f20908101928215610fa9579160200282015b82811115610fa8578251829081610f989190611bc5565b5091602001919060010190610f81565b5b509050610fb69190611015565b5090565b508054610fc690611a01565b5f825580601f10610fd75750610ff4565b601f0160209004905f5260205f2090810190610ff39190611038565b5b50565b5080545f8255905f5260205f20908101906110129190611015565b50565b5b80821115611034575f818161102b9190610fba565b50600101611016565b5090565b5b8082111561104f575f815f905550600101611039565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6110b28261106c565b810181811067ffffffffffffffff821117156110d1576110d061107c565b5b80604052505050565b5f6110e3611053565b90506110ef82826110a9565b919050565b5f67ffffffffffffffff82111561110e5761110d61107c565b5b6111178261106c565b9050602081019050919050565b828183375f83830152505050565b5f61114461113f846110f4565b6110da565b9050828152602081018484840111156111605761115f611068565b5b61116b848285611124565b509392505050565b5f82601f83011261118757611186611064565b5b8135611197848260208601611132565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6111c9826111a0565b9050919050565b6111d9816111bf565b81146111e3575f80fd5b50565b5f813590506111f4816111d0565b92915050565b5f67ffffffffffffffff8211156112145761121361107c565b5b602082029050602081019050919050565b5f80fd5b5f61123b611236846111fa565b6110da565b9050808382526020820190506020840283018581111561125e5761125d611225565b5b835b818110156112a557803567ffffffffffffffff81111561128357611282611064565b5b8086016112908982611173565b85526020850194505050602081019050611260565b5050509392505050565b5f82601f8301126112c3576112c2611064565b5b81356112d3848260208601611229565b91505092915050565b5f805f805f60a086880312156112f5576112f461105c565b5b5f86013567ffffffffffffffff81111561131257611311611060565b5b61131e88828901611173565b955050602061132f888289016111e6565b945050604086013567ffffffffffffffff8111156113505761134f611060565b5b61135c888289016112af565b935050606086013567ffffffffffffffff81111561137d5761137c611060565b5b61138988828901611173565b925050608086013567ffffffffffffffff8111156113aa576113a9611060565b5b6113b688828901611173565b9150509295509295909350565b5f819050919050565b6113d5816113c3565b82525050565b5f6020820190506113ee5f8301846113cc565b92915050565b5f602082840312156114095761140861105c565b5b5f611416848285016111e6565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6114518261141f565b61145b8185611429565b935061146b818560208601611439565b6114748161106c565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f6114c28261141f565b6114cc81856114a8565b93506114dc818560208601611439565b6114e58161106c565b840191505092915050565b5f6114fb83836114b8565b905092915050565b5f602082019050919050565b5f6115198261147f565b6115238185611489565b93508360208202850161153585611499565b805f5b85811015611570578484038952815161155185826114f0565b945061155c83611503565b925060208a01995050600181019050611538565b50829750879550505050505092915050565b5f8115159050919050565b61159681611582565b82525050565b5f60c0820190506115af5f8301896113cc565b81810360208301526115c18188611447565b905081810360408301526115d58187611447565b905081810360608301526115e98186611447565b905081810360808301526115fd818561150f565b905061160c60a083018461158d565b979650505050505050565b611620816113c3565b811461162a575f80fd5b50565b5f8135905061163b81611617565b92915050565b61164a81611582565b8114611654575f80fd5b50565b5f8135905061166581611641565b92915050565b5f80604083850312156116815761168061105c565b5b5f61168e8582860161162d565b925050602061169f85828601611657565b9150509250929050565b5f602082840312156116be576116bd61105c565b5b5f6116cb8482850161162d565b91505092915050565b6116dd816111bf565b82525050565b5f60c0820190508181035f8301526116fb8189611447565b9050818103602083015261170f8188611447565b905081810360408301526117238187611447565b905061173260608301866116d4565b8181036080830152611744818561150f565b905061175360a083018461158d565b979650505050505050565b5f819050919050565b5f61178161177c611777846111a0565b61175e565b6111a0565b9050919050565b5f61179282611767565b9050919050565b5f6117a382611788565b9050919050565b6117b381611799565b82525050565b5f6020820190506117cc5f8301846117aa565b92915050565b5f60c0820190506117e55f8301896113cc565b81810360208301526117f78188611447565b905061180660408301876116d4565b81810360608301526118188186611447565b9050818103608083015261182c8185611447565b905061183b60a083018461158d565b979650505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6020820190508181035f83015261188b8184611447565b905092915050565b5f815190506118a181611641565b92915050565b5f602082840312156118bc576118bb61105c565b5b5f6118c984828501611893565b91505092915050565b7f5061796d656e74206d6574686f64206e6f7420726567697374657265642061735f8201527f206120746f6b656e2e0000000000000000000000000000000000000000000000602082015250565b5f61192c602983611429565b9150611937826118d2565b604082019050919050565b5f6020820190508181035f83015261195981611920565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611997826113c3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119c9576119c8611960565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a1857607f821691505b602082108103611a2b57611a2a6119d4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611a8d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611a52565b611a978683611a52565b95508019841693508086168417925050509392505050565b5f611ac9611ac4611abf846113c3565b61175e565b6113c3565b9050919050565b5f819050919050565b611ae283611aaf565b611af6611aee82611ad0565b848454611a5e565b825550505050565b5f90565b611b0a611afe565b611b15818484611ad9565b505050565b5b81811015611b3857611b2d5f82611b02565b600181019050611b1b565b5050565b601f821115611b7d57611b4e81611a31565b611b5784611a43565b81016020851015611b66578190505b611b7a611b7285611a43565b830182611b1a565b50505b505050565b5f82821c905092915050565b5f611b9d5f1984600802611b82565b1980831691505092915050565b5f611bb58383611b8e565b9150826002028217905092915050565b611bce8261141f565b67ffffffffffffffff811115611be757611be661107c565b5b611bf18254611a01565b611bfc828285611b3c565b5f60209050601f831160018114611c2d575f8415611c1b578287015190505b611c258582611baa565b865550611c8c565b601f198416611c3b86611a31565b5f5b82811015611c6257848901518255600182019150602085019450602081019050611c3d565b86831015611c7f5784890151611c7b601f891682611b8e565b8355505b6001600288020188555050505b505050505050565b5f604082019050611ca75f8301856113cc565b611cb460208301846116d4565b9392505050565b7f4e6f646550726f766964657220776974682074686520676976656e20616464725f8201527f65737320646f6573206e6f742065786973742e00000000000000000000000000602082015250565b5f611d15603383611429565b9150611d2082611cbb565b604082019050919050565b5f6020820190508181035f830152611d4281611d09565b9050919050565b7f4e6f646550726f766964657220646f6573206e6f742065786973742e000000005f82015250565b5f611d7d601c83611429565b9150611d8882611d49565b602082019050919050565b5f6020820190508181035f830152611daa81611d71565b9050919050565b5f604082019050611dc45f8301856113cc565b611dd1602083018461158d565b939250505056fea26469706673582212201625e043dcd5e4ddba952dcbab859b63c67a5cba0c05b54882fd9bf2efa89e2f64736f6c63430008190033",
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
// Solidity: function getNodeProvider(uint256 _validatorId) view returns(string, string, string, address, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) GetNodeProvider(opts *bind.CallOpts, _validatorId *big.Int) (string, string, string, common.Address, []string, bool, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "getNodeProvider", _validatorId)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(common.Address), *new([]string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetNodeProvider is a free data retrieval call binding the contract method 0x915e5511.
//
// Solidity: function getNodeProvider(uint256 _validatorId) view returns(string, string, string, address, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistrySession) GetNodeProvider(_validatorId *big.Int) (string, string, string, common.Address, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProvider(&_NodeProviderRegistry.CallOpts, _validatorId)
}

// GetNodeProvider is a free data retrieval call binding the contract method 0x915e5511.
//
// Solidity: function getNodeProvider(uint256 _validatorId) view returns(string, string, string, address, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) GetNodeProvider(_validatorId *big.Int) (string, string, string, common.Address, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProvider(&_NodeProviderRegistry.CallOpts, _validatorId)
}

// GetNodeProviderByAddress is a free data retrieval call binding the contract method 0x68efbce3.
//
// Solidity: function getNodeProviderByAddress(address _walletAddress) view returns(uint256, string, string, string, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) GetNodeProviderByAddress(opts *bind.CallOpts, _walletAddress common.Address) (*big.Int, string, string, string, []string, bool, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "getNodeProviderByAddress", _walletAddress)

	if err != nil {
		return *new(*big.Int), *new(string), *new(string), *new(string), *new([]string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new([]string)).(*[]string)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// GetNodeProviderByAddress is a free data retrieval call binding the contract method 0x68efbce3.
//
// Solidity: function getNodeProviderByAddress(address _walletAddress) view returns(uint256, string, string, string, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistrySession) GetNodeProviderByAddress(_walletAddress common.Address) (*big.Int, string, string, string, []string, bool, error) {
	return _NodeProviderRegistry.Contract.GetNodeProviderByAddress(&_NodeProviderRegistry.CallOpts, _walletAddress)
}

// GetNodeProviderByAddress is a free data retrieval call binding the contract method 0x68efbce3.
//
// Solidity: function getNodeProviderByAddress(address _walletAddress) view returns(uint256, string, string, string, string[], bool)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) GetNodeProviderByAddress(_walletAddress common.Address) (*big.Int, string, string, string, []string, bool, error) {
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
// Solidity: function nodeProviders(uint256 ) view returns(uint256 providerId, string region, address walletAddress, string attributes, string hostUri, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryCaller) NodeProviders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProviderId    *big.Int
	Region        string
	WalletAddress common.Address
	Attributes    string
	HostUri       string
	IsActive      bool
}, error) {
	var out []interface{}
	err := _NodeProviderRegistry.contract.Call(opts, &out, "nodeProviders", arg0)

	outstruct := new(struct {
		ProviderId    *big.Int
		Region        string
		WalletAddress common.Address
		Attributes    string
		HostUri       string
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProviderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Region = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.WalletAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Attributes = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.HostUri = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 providerId, string region, address walletAddress, string attributes, string hostUri, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistrySession) NodeProviders(arg0 *big.Int) (struct {
	ProviderId    *big.Int
	Region        string
	WalletAddress common.Address
	Attributes    string
	HostUri       string
	IsActive      bool
}, error) {
	return _NodeProviderRegistry.Contract.NodeProviders(&_NodeProviderRegistry.CallOpts, arg0)
}

// NodeProviders is a free data retrieval call binding the contract method 0xeadbc589.
//
// Solidity: function nodeProviders(uint256 ) view returns(uint256 providerId, string region, address walletAddress, string attributes, string hostUri, bool isActive)
func (_NodeProviderRegistry *NodeProviderRegistryCallerSession) NodeProviders(arg0 *big.Int) (struct {
	ProviderId    *big.Int
	Region        string
	WalletAddress common.Address
	Attributes    string
	HostUri       string
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

// AddNodeProvider is a paid mutator transaction binding the contract method 0x1432d4f9.
//
// Solidity: function addNodeProvider(string _region, address _walletAddress, string[] _paymentsAccepted, string _attributes, string _hostUri) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactor) AddNodeProvider(opts *bind.TransactOpts, _region string, _walletAddress common.Address, _paymentsAccepted []string, _attributes string, _hostUri string) (*types.Transaction, error) {
	return _NodeProviderRegistry.contract.Transact(opts, "addNodeProvider", _region, _walletAddress, _paymentsAccepted, _attributes, _hostUri)
}

// AddNodeProvider is a paid mutator transaction binding the contract method 0x1432d4f9.
//
// Solidity: function addNodeProvider(string _region, address _walletAddress, string[] _paymentsAccepted, string _attributes, string _hostUri) returns()
func (_NodeProviderRegistry *NodeProviderRegistrySession) AddNodeProvider(_region string, _walletAddress common.Address, _paymentsAccepted []string, _attributes string, _hostUri string) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.AddNodeProvider(&_NodeProviderRegistry.TransactOpts, _region, _walletAddress, _paymentsAccepted, _attributes, _hostUri)
}

// AddNodeProvider is a paid mutator transaction binding the contract method 0x1432d4f9.
//
// Solidity: function addNodeProvider(string _region, address _walletAddress, string[] _paymentsAccepted, string _attributes, string _hostUri) returns()
func (_NodeProviderRegistry *NodeProviderRegistryTransactorSession) AddNodeProvider(_region string, _walletAddress common.Address, _paymentsAccepted []string, _attributes string, _hostUri string) (*types.Transaction, error) {
	return _NodeProviderRegistry.Contract.AddNodeProvider(&_NodeProviderRegistry.TransactOpts, _region, _walletAddress, _paymentsAccepted, _attributes, _hostUri)
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
