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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prodmanager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logismanager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_salemanager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_repairmanager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addLogisData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addProdData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addRepairData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addSaleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getLogisData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getProdData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getRepairData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"carId\",\"type\":\"string\"}],\"name\":\"getSaleData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logismanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prodmanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repairmanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salemanager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620015bf380380620015bf8339818101604052810190620000369190620001a4565b835f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000213565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200016e8262000143565b9050919050565b620001808162000162565b81146200018b575f80fd5b50565b5f815190506200019e8162000175565b92915050565b5f805f8060808587031215620001bf57620001be6200013f565b5b5f620001ce878288016200018e565b9450506020620001e1878288016200018e565b9350506040620001f4878288016200018e565b925050606062000207878288016200018e565b91505092959194509250565b61139e80620002215f395ff3fe6080604052600436106100aa575f3560e01c806391711cb81161006357806391711cb8146101d15780639583f02d146101fb5780639f07b36914610225578063ade8ca451461024d578063c039bbf914610289578063de2d0005146102c5576100b1565b8063064fa933146100b557806309b7f42a146100f15780630e6a3d5614610119578063156bee95146101415780635d562e971461016b57806360b1adab146101a7576100b1565b366100b157005b5f80fd5b3480156100c0575f80fd5b506100db60048036038101906100d69190610c2c565b6102ed565b6040516100e89190610da8565b60405180910390f35b3480156100fc575f80fd5b5061011760048036038101906101129190610dc8565b6103e2565b005b348015610124575f80fd5b5061013f600480360381019061013a9190610dc8565b6104c6565b005b34801561014c575f80fd5b506101556105a7565b6040516101629190610e7d565b60405180910390f35b348015610176575f80fd5b50610191600480360381019061018c9190610c2c565b6105cc565b60405161019e9190610da8565b60405180910390f35b3480156101b2575f80fd5b506101bb6106c1565b6040516101c89190610e7d565b60405180910390f35b3480156101dc575f80fd5b506101e56106e4565b6040516101f29190610e7d565b60405180910390f35b348015610206575f80fd5b5061020f610709565b60405161021c9190610e7d565b60405180910390f35b348015610230575f80fd5b5061024b60048036038101906102469190610dc8565b61072e565b005b348015610258575f80fd5b50610273600480360381019061026e9190610c2c565b610812565b6040516102809190610da8565b60405180910390f35b348015610294575f80fd5b506102af60048036038101906102aa9190610c2c565b610907565b6040516102bc9190610da8565b60405180910390f35b3480156102d0575f80fd5b506102eb60048036038101906102e69190610dc8565b6109fb565b005b60606004826040516102ff9190610ed0565b9081526020016040518091039020600301805480602002602001604051908101604052809291908181526020015f905b828210156103d7578382905f5260205f2001805461034c90610f13565b80601f016020809104026020016040519081016040528092919081815260200182805461037890610f13565b80156103c35780601f1061039a576101008083540402835291602001916103c3565b820191905f5260205f20905b8154815290600101906020018083116103a657829003601f168201915b50505050508152602001906001019061032f565b505050509050919050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046890610f9d565b60405180910390fd5b6004826040516104819190610ed0565b908152602001604051809103902060030181908060018154018082558091505060019003905f5260205f20015f9091909190915090816104c19190611161565b505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610553576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054a9061127a565b60405180910390fd5b6004826040516105639190610ed0565b90815260200160405180910390205f0181908060018154018082558091505060019003905f5260205f20015f9091909190915090816105a29190611161565b505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606004826040516105de9190610ed0565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020015f905b828210156106b6578382905f5260205f2001805461062b90610f13565b80601f016020809104026020016040519081016040528092919081815260200182805461065790610f13565b80156106a25780601f10610679576101008083540402835291602001916106a2565b820191905f5260205f20905b81548152906001019060200180831161068557829003601f168201915b50505050508152602001906001019061060e565b505050509050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b4906112e2565b60405180910390fd5b6004826040516107cd9190610ed0565b908152602001604051809103902060020181908060018154018082558091505060019003905f5260205f20015f90919091909150908161080d9190611161565b505050565b60606004826040516108249190610ed0565b9081526020016040518091039020600201805480602002602001604051908101604052809291908181526020015f905b828210156108fc578382905f5260205f2001805461087190610f13565b80601f016020809104026020016040519081016040528092919081815260200182805461089d90610f13565b80156108e85780601f106108bf576101008083540402835291602001916108e8565b820191905f5260205f20905b8154815290600101906020018083116108cb57829003601f168201915b505050505081526020019060010190610854565b505050509050919050565b60606004826040516109199190610ed0565b90815260200160405180910390205f01805480602002602001604051908101604052809291908181526020015f905b828210156109f0578382905f5260205f2001805461096590610f13565b80601f016020809104026020016040519081016040528092919081815260200182805461099190610f13565b80156109dc5780601f106109b3576101008083540402835291602001916109dc565b820191905f5260205f20905b8154815290600101906020018083116109bf57829003601f168201915b505050505081526020019060010190610948565b505050509050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a819061134a565b60405180910390fd5b600482604051610a9a9190610ed0565b908152602001604051809103902060010181908060018154018082558091505060019003905f5260205f20015f909190919091509081610ada9190611161565b505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b3e82610af8565b810181811067ffffffffffffffff82111715610b5d57610b5c610b08565b5b80604052505050565b5f610b6f610adf565b9050610b7b8282610b35565b919050565b5f67ffffffffffffffff821115610b9a57610b99610b08565b5b610ba382610af8565b9050602081019050919050565b828183375f83830152505050565b5f610bd0610bcb84610b80565b610b66565b905082815260208101848484011115610bec57610beb610af4565b5b610bf7848285610bb0565b509392505050565b5f82601f830112610c1357610c12610af0565b5b8135610c23848260208601610bbe565b91505092915050565b5f60208284031215610c4157610c40610ae8565b5b5f82013567ffffffffffffffff811115610c5e57610c5d610aec565b5b610c6a84828501610bff565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610cd3578082015181840152602081019050610cb8565b5f8484015250505050565b5f610ce882610c9c565b610cf28185610ca6565b9350610d02818560208601610cb6565b610d0b81610af8565b840191505092915050565b5f610d218383610cde565b905092915050565b5f602082019050919050565b5f610d3f82610c73565b610d498185610c7d565b935083602082028501610d5b85610c8d565b805f5b85811015610d965784840389528151610d778582610d16565b9450610d8283610d29565b925060208a01995050600181019050610d5e565b50829750879550505050505092915050565b5f6020820190508181035f830152610dc08184610d35565b905092915050565b5f8060408385031215610dde57610ddd610ae8565b5b5f83013567ffffffffffffffff811115610dfb57610dfa610aec565b5b610e0785828601610bff565b925050602083013567ffffffffffffffff811115610e2857610e27610aec565b5b610e3485828601610bff565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e6782610e3e565b9050919050565b610e7781610e5d565b82525050565b5f602082019050610e905f830184610e6e565b92915050565b5f81905092915050565b5f610eaa82610c9c565b610eb48185610e96565b9350610ec4818560208601610cb6565b80840191505092915050565b5f610edb8284610ea0565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f2a57607f821691505b602082108103610f3d57610f3c610ee6565b5b50919050565b5f82825260208201905092915050565b7f4f4e4c595f5245504149525f4d414e414745525f524f4c4500000000000000005f82015250565b5f610f87601883610f43565b9150610f9282610f53565b602082019050919050565b5f6020820190508181035f830152610fb481610f7b565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026110177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fdc565b6110218683610fdc565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61106561106061105b84611039565b611042565b611039565b9050919050565b5f819050919050565b61107e8361104b565b61109261108a8261106c565b848454610fe8565b825550505050565b5f90565b6110a661109a565b6110b1818484611075565b505050565b5b818110156110d4576110c95f8261109e565b6001810190506110b7565b5050565b601f821115611119576110ea81610fbb565b6110f384610fcd565b81016020851015611102578190505b61111661110e85610fcd565b8301826110b6565b50505b505050565b5f82821c905092915050565b5f6111395f198460080261111e565b1980831691505092915050565b5f611151838361112a565b9150826002028217905092915050565b61116a82610c9c565b67ffffffffffffffff81111561118357611182610b08565b5b61118d8254610f13565b6111988282856110d8565b5f60209050601f8311600181146111c9575f84156111b7578287015190505b6111c18582611146565b865550611228565b601f1984166111d786610fbb565b5f5b828110156111fe578489015182556001820191506020850194506020810190506111d9565b8683101561121b5784890151611217601f89168261112a565b8355505b6001600288020188555050505b505050505050565b7f4f4e4c595f50524f445f4d414e414745525f524f4c45000000000000000000005f82015250565b5f611264601683610f43565b915061126f82611230565b602082019050919050565b5f6020820190508181035f83015261129181611258565b9050919050565b7f4f4e4c595f53414c455f4d414e414745525f524f4c45000000000000000000005f82015250565b5f6112cc601683610f43565b91506112d782611298565b602082019050919050565b5f6020820190508181035f8301526112f9816112c0565b9050919050565b7f4f4e4c595f4c4f4749535f4d414e414745525f524f4c450000000000000000005f82015250565b5f611334601783610f43565b915061133f82611300565b602082019050919050565b5f6020820190508181035f83015261136181611328565b905091905056fea2646970667358221220b1b6c1f9e5143776f3774165026d9d5a3802bee19e5eabc5aa020695e88b55ba64736f6c63430008180033",
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
