// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

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

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prodmanager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logismanager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_salemanager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repairmanager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addLogisData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addProdData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addRepairData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addSaleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getLogisData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getProdData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getRepairData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getSaleData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logismanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prodmanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repairmanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salemanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_r\",\"type\":\"address\"}],\"name\":\"setLogisManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_r\",\"type\":\"address\"}],\"name\":\"setProdManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_r\",\"type\":\"address\"}],\"name\":\"setRepairManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_r\",\"type\":\"address\"}],\"name\":\"setSaleManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001b4438038062001b448339818101604052810190620000369190620001e4565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000253565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001ae8262000183565b9050919050565b620001c081620001a2565b8114620001cb575f80fd5b50565b5f81519050620001de81620001b5565b92915050565b5f805f8060808587031215620001ff57620001fe6200017f565b5b5f6200020e87828801620001ce565b94505060206200022187828801620001ce565b93505060406200023487828801620001ce565b92505060606200024787828801620001ce565b91505092959194509250565b6118e380620002615f395ff3fe608060405260043610610101575f3560e01c80638da5cb5b11610094578063ade8ca4511610063578063ade8ca451461031e578063b86a73ce1461035a578063c039bbf914610382578063de2d0005146103be578063f88b2f4d146103e657610108565b80638da5cb5b1461027857806391711cb8146102a25780639583f02d146102cc5780639f07b369146102f657610108565b80635d562e97116100d05780635d562e97146101c257806360b1adab146101fe57806367e9499514610228578063808c559f1461025057610108565b8063064fa9331461010c57806309b7f42a146101485780630e6a3d5614610170578063156bee951461019857610108565b3661010857005b5f80fd5b348015610117575f80fd5b50610132600480360381019061012d91906110b4565b61040e565b60405161013f9190611230565b60405180910390f35b348015610153575f80fd5b5061016e60048036038101906101699190611250565b610503565b005b34801561017b575f80fd5b5061019660048036038101906101919190611250565b6105e7565b005b3480156101a3575f80fd5b506101ac6106ca565b6040516101b99190611305565b60405180910390f35b3480156101cd575f80fd5b506101e860048036038101906101e391906110b4565b6106ef565b6040516101f59190611230565b60405180910390f35b348015610209575f80fd5b506102126107e4565b60405161021f9190611305565b60405180910390f35b348015610233575f80fd5b5061024e60048036038101906102499190611348565b610809565b005b34801561025b575f80fd5b5061027660048036038101906102719190611348565b6108d9565b005b348015610283575f80fd5b5061028c6109a9565b6040516102999190611305565b60405180910390f35b3480156102ad575f80fd5b506102b66109cc565b6040516102c39190611305565b60405180910390f35b3480156102d7575f80fd5b506102e06109f1565b6040516102ed9190611305565b60405180910390f35b348015610301575f80fd5b5061031c60048036038101906103179190611250565b610a16565b005b348015610329575f80fd5b50610344600480360381019061033f91906110b4565b610afa565b6040516103519190611230565b60405180910390f35b348015610365575f80fd5b50610380600480360381019061037b9190611348565b610bef565b005b34801561038d575f80fd5b506103a860048036038101906103a391906110b4565b610cbf565b6040516103b59190611230565b60405180910390f35b3480156103c9575f80fd5b506103e460048036038101906103df9190611250565b610db3565b005b3480156103f1575f80fd5b5061040c60048036038101906104079190611348565b610e97565b005b606060058260405161042091906113ad565b9081526020016040518091039020600301805480602002602001604051908101604052809291908181526020015f905b828210156104f8578382905f5260205f2001805461046d906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610499906113f0565b80156104e45780601f106104bb576101008083540402835291602001916104e4565b820191905f5260205f20905b8154815290600101906020018083116104c757829003601f168201915b505050505081526020019060010190610450565b505050509050919050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610592576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105899061147a565b60405180910390fd5b6005826040516105a291906113ad565b908152602001604051809103902060030181908060018154018082558091505060019003905f5260205f20015f9091909190915090816105e2919061163e565b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066d90611757565b60405180910390fd5b60058260405161068691906113ad565b90815260200160405180910390205f0181908060018154018082558091505060019003905f5260205f20015f9091909190915090816106c5919061163e565b505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060058260405161070191906113ad565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020015f905b828210156107d9578382905f5260205f2001805461074e906113f0565b80601f016020809104026020016040519081016040528092919081815260200182805461077a906113f0565b80156107c55780601f1061079c576101008083540402835291602001916107c5565b820191905f5260205f20905b8154815290600101906020018083116107a857829003601f168201915b505050505081526020019060010190610731565b505050509050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610896576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088d906117bf565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610966576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095d906117bf565b60405180910390fd5b8060045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610aa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9c90611827565b60405180910390fd5b600582604051610ab591906113ad565b908152602001604051809103902060020181908060018154018082558091505060019003905f5260205f20015f909190919091509081610af5919061163e565b505050565b6060600582604051610b0c91906113ad565b9081526020016040518091039020600201805480602002602001604051908101604052809291908181526020015f905b82821015610be4578382905f5260205f20018054610b59906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610b85906113f0565b8015610bd05780601f10610ba757610100808354040283529160200191610bd0565b820191905f5260205f20905b815481529060010190602001808311610bb357829003601f168201915b505050505081526020019060010190610b3c565b505050509050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c73906117bf565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6060600582604051610cd191906113ad565b90815260200160405180910390205f01805480602002602001604051908101604052809291908181526020015f905b82821015610da8578382905f5260205f20018054610d1d906113f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610d49906113f0565b8015610d945780601f10610d6b57610100808354040283529160200191610d94565b820191905f5260205f20905b815481529060010190602001808311610d7757829003601f168201915b505050505081526020019060010190610d00565b505050509050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e399061188f565b60405180910390fd5b600582604051610e5291906113ad565b908152602001604051809103902060010181908060018154018082558091505060019003905f5260205f20015f909190919091509081610e92919061163e565b505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1b906117bf565b60405180910390fd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610fc682610f80565b810181811067ffffffffffffffff82111715610fe557610fe4610f90565b5b80604052505050565b5f610ff7610f67565b90506110038282610fbd565b919050565b5f67ffffffffffffffff82111561102257611021610f90565b5b61102b82610f80565b9050602081019050919050565b828183375f83830152505050565b5f61105861105384611008565b610fee565b90508281526020810184848401111561107457611073610f7c565b5b61107f848285611038565b509392505050565b5f82601f83011261109b5761109a610f78565b5b81356110ab848260208601611046565b91505092915050565b5f602082840312156110c9576110c8610f70565b5b5f82013567ffffffffffffffff8111156110e6576110e5610f74565b5b6110f284828501611087565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561115b578082015181840152602081019050611140565b5f8484015250505050565b5f61117082611124565b61117a818561112e565b935061118a81856020860161113e565b61119381610f80565b840191505092915050565b5f6111a98383611166565b905092915050565b5f602082019050919050565b5f6111c7826110fb565b6111d18185611105565b9350836020820285016111e385611115565b805f5b8581101561121e57848403895281516111ff858261119e565b945061120a836111b1565b925060208a019950506001810190506111e6565b50829750879550505050505092915050565b5f6020820190508181035f83015261124881846111bd565b905092915050565b5f806040838503121561126657611265610f70565b5b5f83013567ffffffffffffffff81111561128357611282610f74565b5b61128f85828601611087565b925050602083013567ffffffffffffffff8111156112b0576112af610f74565b5b6112bc85828601611087565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112ef826112c6565b9050919050565b6112ff816112e5565b82525050565b5f6020820190506113185f8301846112f6565b92915050565b611327816112e5565b8114611331575f80fd5b50565b5f813590506113428161131e565b92915050565b5f6020828403121561135d5761135c610f70565b5b5f61136a84828501611334565b91505092915050565b5f81905092915050565b5f61138782611124565b6113918185611373565b93506113a181856020860161113e565b80840191505092915050565b5f6113b8828461137d565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061140757607f821691505b60208210810361141a576114196113c3565b5b50919050565b5f82825260208201905092915050565b7f4f4e4c595f5245504149525f4d414e414745525f524f4c4500000000000000005f82015250565b5f611464601883611420565b915061146f82611430565b602082019050919050565b5f6020820190508181035f83015261149181611458565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026114f47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114b9565b6114fe86836114b9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61154261153d61153884611516565b61151f565b611516565b9050919050565b5f819050919050565b61155b83611528565b61156f61156782611549565b8484546114c5565b825550505050565b5f90565b611583611577565b61158e818484611552565b505050565b5b818110156115b1576115a65f8261157b565b600181019050611594565b5050565b601f8211156115f6576115c781611498565b6115d0846114aa565b810160208510156115df578190505b6115f36115eb856114aa565b830182611593565b50505b505050565b5f82821c905092915050565b5f6116165f19846008026115fb565b1980831691505092915050565b5f61162e8383611607565b9150826002028217905092915050565b61164782611124565b67ffffffffffffffff8111156116605761165f610f90565b5b61166a82546113f0565b6116758282856115b5565b5f60209050601f8311600181146116a6575f8415611694578287015190505b61169e8582611623565b865550611705565b601f1984166116b486611498565b5f5b828110156116db578489015182556001820191506020850194506020810190506116b6565b868310156116f857848901516116f4601f891682611607565b8355505b6001600288020188555050505b505050505050565b7f4f4e4c595f50524f445f4d414e414745525f524f4c45000000000000000000005f82015250565b5f611741601683611420565b915061174c8261170d565b602082019050919050565b5f6020820190508181035f83015261176e81611735565b9050919050565b7f4f4e4c595f4f574e45525f524f4c4500000000000000000000000000000000005f82015250565b5f6117a9600f83611420565b91506117b482611775565b602082019050919050565b5f6020820190508181035f8301526117d68161179d565b9050919050565b7f4f4e4c595f53414c455f4d414e414745525f524f4c45000000000000000000005f82015250565b5f611811601683611420565b915061181c826117dd565b602082019050919050565b5f6020820190508181035f83015261183e81611805565b9050919050565b7f4f4e4c595f4c4f4749535f4d414e414745525f524f4c450000000000000000005f82015250565b5f611879601783611420565b915061188482611845565b602082019050919050565b5f6020820190508181035f8301526118a68161186d565b905091905056fea26469706673582212200c10ff5e5c846f40b8183404fb7b75c9066afa414f83abe2cc819fd4163f2ebe64736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend, _prodmanager common.Address, _logismanager common.Address, _salemanager common.Address, _repairmanager common.Address) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend, _prodmanager, _logismanager, _salemanager, _repairmanager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// GetLogisData is a free data retrieval call binding the contract method 0x5d562e97.
//
// Solidity: function getLogisData(string carId) view returns(string[])
func (_Trace *TraceCaller) GetLogisData(opts *bind.CallOpts, carId string) ([]string, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "getLogisData", carId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetLogisData is a free data retrieval call binding the contract method 0x5d562e97.
//
// Solidity: function getLogisData(string carId) view returns(string[])
func (_Trace *TraceSession) GetLogisData(carId string) ([]string, error) {
	return _Trace.Contract.GetLogisData(&_Trace.CallOpts, carId)
}

// GetLogisData is a free data retrieval call binding the contract method 0x5d562e97.
//
// Solidity: function getLogisData(string carId) view returns(string[])
func (_Trace *TraceCallerSession) GetLogisData(carId string) ([]string, error) {
	return _Trace.Contract.GetLogisData(&_Trace.CallOpts, carId)
}

// GetProdData is a free data retrieval call binding the contract method 0xc039bbf9.
//
// Solidity: function getProdData(string carId) view returns(string[])
func (_Trace *TraceCaller) GetProdData(opts *bind.CallOpts, carId string) ([]string, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "getProdData", carId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetProdData is a free data retrieval call binding the contract method 0xc039bbf9.
//
// Solidity: function getProdData(string carId) view returns(string[])
func (_Trace *TraceSession) GetProdData(carId string) ([]string, error) {
	return _Trace.Contract.GetProdData(&_Trace.CallOpts, carId)
}

// GetProdData is a free data retrieval call binding the contract method 0xc039bbf9.
//
// Solidity: function getProdData(string carId) view returns(string[])
func (_Trace *TraceCallerSession) GetProdData(carId string) ([]string, error) {
	return _Trace.Contract.GetProdData(&_Trace.CallOpts, carId)
}

// GetRepairData is a free data retrieval call binding the contract method 0x064fa933.
//
// Solidity: function getRepairData(string carId) view returns(string[])
func (_Trace *TraceCaller) GetRepairData(opts *bind.CallOpts, carId string) ([]string, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "getRepairData", carId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetRepairData is a free data retrieval call binding the contract method 0x064fa933.
//
// Solidity: function getRepairData(string carId) view returns(string[])
func (_Trace *TraceSession) GetRepairData(carId string) ([]string, error) {
	return _Trace.Contract.GetRepairData(&_Trace.CallOpts, carId)
}

// GetRepairData is a free data retrieval call binding the contract method 0x064fa933.
//
// Solidity: function getRepairData(string carId) view returns(string[])
func (_Trace *TraceCallerSession) GetRepairData(carId string) ([]string, error) {
	return _Trace.Contract.GetRepairData(&_Trace.CallOpts, carId)
}

// GetSaleData is a free data retrieval call binding the contract method 0xade8ca45.
//
// Solidity: function getSaleData(string carId) view returns(string[])
func (_Trace *TraceCaller) GetSaleData(opts *bind.CallOpts, carId string) ([]string, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "getSaleData", carId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetSaleData is a free data retrieval call binding the contract method 0xade8ca45.
//
// Solidity: function getSaleData(string carId) view returns(string[])
func (_Trace *TraceSession) GetSaleData(carId string) ([]string, error) {
	return _Trace.Contract.GetSaleData(&_Trace.CallOpts, carId)
}

// GetSaleData is a free data retrieval call binding the contract method 0xade8ca45.
//
// Solidity: function getSaleData(string carId) view returns(string[])
func (_Trace *TraceCallerSession) GetSaleData(carId string) ([]string, error) {
	return _Trace.Contract.GetSaleData(&_Trace.CallOpts, carId)
}

// Logismanager is a free data retrieval call binding the contract method 0x91711cb8.
//
// Solidity: function logismanager() view returns(address)
func (_Trace *TraceCaller) Logismanager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "logismanager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Logismanager is a free data retrieval call binding the contract method 0x91711cb8.
//
// Solidity: function logismanager() view returns(address)
func (_Trace *TraceSession) Logismanager() (common.Address, error) {
	return _Trace.Contract.Logismanager(&_Trace.CallOpts)
}

// Logismanager is a free data retrieval call binding the contract method 0x91711cb8.
//
// Solidity: function logismanager() view returns(address)
func (_Trace *TraceCallerSession) Logismanager() (common.Address, error) {
	return _Trace.Contract.Logismanager(&_Trace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceSession) Owner() (common.Address, error) {
	return _Trace.Contract.Owner(&_Trace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trace *TraceCallerSession) Owner() (common.Address, error) {
	return _Trace.Contract.Owner(&_Trace.CallOpts)
}

// Prodmanager is a free data retrieval call binding the contract method 0x60b1adab.
//
// Solidity: function prodmanager() view returns(address)
func (_Trace *TraceCaller) Prodmanager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "prodmanager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prodmanager is a free data retrieval call binding the contract method 0x60b1adab.
//
// Solidity: function prodmanager() view returns(address)
func (_Trace *TraceSession) Prodmanager() (common.Address, error) {
	return _Trace.Contract.Prodmanager(&_Trace.CallOpts)
}

// Prodmanager is a free data retrieval call binding the contract method 0x60b1adab.
//
// Solidity: function prodmanager() view returns(address)
func (_Trace *TraceCallerSession) Prodmanager() (common.Address, error) {
	return _Trace.Contract.Prodmanager(&_Trace.CallOpts)
}

// Repairmanager is a free data retrieval call binding the contract method 0x9583f02d.
//
// Solidity: function repairmanager() view returns(address)
func (_Trace *TraceCaller) Repairmanager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "repairmanager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Repairmanager is a free data retrieval call binding the contract method 0x9583f02d.
//
// Solidity: function repairmanager() view returns(address)
func (_Trace *TraceSession) Repairmanager() (common.Address, error) {
	return _Trace.Contract.Repairmanager(&_Trace.CallOpts)
}

// Repairmanager is a free data retrieval call binding the contract method 0x9583f02d.
//
// Solidity: function repairmanager() view returns(address)
func (_Trace *TraceCallerSession) Repairmanager() (common.Address, error) {
	return _Trace.Contract.Repairmanager(&_Trace.CallOpts)
}

// Salemanager is a free data retrieval call binding the contract method 0x156bee95.
//
// Solidity: function salemanager() view returns(address)
func (_Trace *TraceCaller) Salemanager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "salemanager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Salemanager is a free data retrieval call binding the contract method 0x156bee95.
//
// Solidity: function salemanager() view returns(address)
func (_Trace *TraceSession) Salemanager() (common.Address, error) {
	return _Trace.Contract.Salemanager(&_Trace.CallOpts)
}

// Salemanager is a free data retrieval call binding the contract method 0x156bee95.
//
// Solidity: function salemanager() view returns(address)
func (_Trace *TraceCallerSession) Salemanager() (common.Address, error) {
	return _Trace.Contract.Salemanager(&_Trace.CallOpts)
}

// AddLogisData is a paid mutator transaction binding the contract method 0xde2d0005.
//
// Solidity: function addLogisData(string carId, string description) returns()
func (_Trace *TraceTransactor) AddLogisData(opts *bind.TransactOpts, carId string, description string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addLogisData", carId, description)
}

// AddLogisData is a paid mutator transaction binding the contract method 0xde2d0005.
//
// Solidity: function addLogisData(string carId, string description) returns()
func (_Trace *TraceSession) AddLogisData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddLogisData(&_Trace.TransactOpts, carId, description)
}

// AddLogisData is a paid mutator transaction binding the contract method 0xde2d0005.
//
// Solidity: function addLogisData(string carId, string description) returns()
func (_Trace *TraceTransactorSession) AddLogisData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddLogisData(&_Trace.TransactOpts, carId, description)
}

// AddProdData is a paid mutator transaction binding the contract method 0x0e6a3d56.
//
// Solidity: function addProdData(string carId, string description) returns()
func (_Trace *TraceTransactor) AddProdData(opts *bind.TransactOpts, carId string, description string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addProdData", carId, description)
}

// AddProdData is a paid mutator transaction binding the contract method 0x0e6a3d56.
//
// Solidity: function addProdData(string carId, string description) returns()
func (_Trace *TraceSession) AddProdData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddProdData(&_Trace.TransactOpts, carId, description)
}

// AddProdData is a paid mutator transaction binding the contract method 0x0e6a3d56.
//
// Solidity: function addProdData(string carId, string description) returns()
func (_Trace *TraceTransactorSession) AddProdData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddProdData(&_Trace.TransactOpts, carId, description)
}

// AddRepairData is a paid mutator transaction binding the contract method 0x09b7f42a.
//
// Solidity: function addRepairData(string carId, string description) returns()
func (_Trace *TraceTransactor) AddRepairData(opts *bind.TransactOpts, carId string, description string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addRepairData", carId, description)
}

// AddRepairData is a paid mutator transaction binding the contract method 0x09b7f42a.
//
// Solidity: function addRepairData(string carId, string description) returns()
func (_Trace *TraceSession) AddRepairData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddRepairData(&_Trace.TransactOpts, carId, description)
}

// AddRepairData is a paid mutator transaction binding the contract method 0x09b7f42a.
//
// Solidity: function addRepairData(string carId, string description) returns()
func (_Trace *TraceTransactorSession) AddRepairData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddRepairData(&_Trace.TransactOpts, carId, description)
}

// AddSaleData is a paid mutator transaction binding the contract method 0x9f07b369.
//
// Solidity: function addSaleData(string carId, string description) returns()
func (_Trace *TraceTransactor) AddSaleData(opts *bind.TransactOpts, carId string, description string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addSaleData", carId, description)
}

// AddSaleData is a paid mutator transaction binding the contract method 0x9f07b369.
//
// Solidity: function addSaleData(string carId, string description) returns()
func (_Trace *TraceSession) AddSaleData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddSaleData(&_Trace.TransactOpts, carId, description)
}

// AddSaleData is a paid mutator transaction binding the contract method 0x9f07b369.
//
// Solidity: function addSaleData(string carId, string description) returns()
func (_Trace *TraceTransactorSession) AddSaleData(carId string, description string) (*types.Transaction, error) {
	return _Trace.Contract.AddSaleData(&_Trace.TransactOpts, carId, description)
}

// SetLogisManager is a paid mutator transaction binding the contract method 0x67e94995.
//
// Solidity: function setLogisManager(address _r) returns()
func (_Trace *TraceTransactor) SetLogisManager(opts *bind.TransactOpts, _r common.Address) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "setLogisManager", _r)
}

// SetLogisManager is a paid mutator transaction binding the contract method 0x67e94995.
//
// Solidity: function setLogisManager(address _r) returns()
func (_Trace *TraceSession) SetLogisManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetLogisManager(&_Trace.TransactOpts, _r)
}

// SetLogisManager is a paid mutator transaction binding the contract method 0x67e94995.
//
// Solidity: function setLogisManager(address _r) returns()
func (_Trace *TraceTransactorSession) SetLogisManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetLogisManager(&_Trace.TransactOpts, _r)
}

// SetProdManager is a paid mutator transaction binding the contract method 0xb86a73ce.
//
// Solidity: function setProdManager(address _r) returns()
func (_Trace *TraceTransactor) SetProdManager(opts *bind.TransactOpts, _r common.Address) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "setProdManager", _r)
}

// SetProdManager is a paid mutator transaction binding the contract method 0xb86a73ce.
//
// Solidity: function setProdManager(address _r) returns()
func (_Trace *TraceSession) SetProdManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetProdManager(&_Trace.TransactOpts, _r)
}

// SetProdManager is a paid mutator transaction binding the contract method 0xb86a73ce.
//
// Solidity: function setProdManager(address _r) returns()
func (_Trace *TraceTransactorSession) SetProdManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetProdManager(&_Trace.TransactOpts, _r)
}

// SetRepairManager is a paid mutator transaction binding the contract method 0x808c559f.
//
// Solidity: function setRepairManager(address _r) returns()
func (_Trace *TraceTransactor) SetRepairManager(opts *bind.TransactOpts, _r common.Address) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "setRepairManager", _r)
}

// SetRepairManager is a paid mutator transaction binding the contract method 0x808c559f.
//
// Solidity: function setRepairManager(address _r) returns()
func (_Trace *TraceSession) SetRepairManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetRepairManager(&_Trace.TransactOpts, _r)
}

// SetRepairManager is a paid mutator transaction binding the contract method 0x808c559f.
//
// Solidity: function setRepairManager(address _r) returns()
func (_Trace *TraceTransactorSession) SetRepairManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetRepairManager(&_Trace.TransactOpts, _r)
}

// SetSaleManager is a paid mutator transaction binding the contract method 0xf88b2f4d.
//
// Solidity: function setSaleManager(address _r) returns()
func (_Trace *TraceTransactor) SetSaleManager(opts *bind.TransactOpts, _r common.Address) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "setSaleManager", _r)
}

// SetSaleManager is a paid mutator transaction binding the contract method 0xf88b2f4d.
//
// Solidity: function setSaleManager(address _r) returns()
func (_Trace *TraceSession) SetSaleManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetSaleManager(&_Trace.TransactOpts, _r)
}

// SetSaleManager is a paid mutator transaction binding the contract method 0xf88b2f4d.
//
// Solidity: function setSaleManager(address _r) returns()
func (_Trace *TraceTransactorSession) SetSaleManager(_r common.Address) (*types.Transaction, error) {
	return _Trace.Contract.SetSaleManager(&_Trace.TransactOpts, _r)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Trace *TraceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Trace *TraceSession) Receive() (*types.Transaction, error) {
	return _Trace.Contract.Receive(&_Trace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Trace *TraceTransactorSession) Receive() (*types.Transaction, error) {
	return _Trace.Contract.Receive(&_Trace.TransactOpts)
}
