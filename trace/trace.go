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

// UserPassStoreOrgInfo is an auto generated low-level Go binding around an user-defined struct.
type UserPassStoreOrgInfo struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workamount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"persion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workmethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"addorupdateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tracedata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workamount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"persion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workmethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"}],\"name\":\"tracedataByNum\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workamount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"persion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workmethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"internalType\":\"structUserPassStore.OrgInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506111228061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80631bb2bfea146100435780635c7e9495146100735780638a205085146100a8575b5f80fd5b61005d60048036038101906100589190610a3e565b6100c4565b60405161006a9190610ba8565b60405180910390f35b61008d60048036038101906100889190610a3e565b61045a565b60405161009f96959493929190610c10565b60405180910390f35b6100c260048036038101906100bd9190610c99565b6107cd565b005b6100cc6108bb565b5f826040516100db9190610e04565b90815260200160405180910390206040518060c00160405290815f8201805461010390610e47565b80601f016020809104026020016040519081016040528092919081815260200182805461012f90610e47565b801561017a5780601f106101515761010080835404028352916020019161017a565b820191905f5260205f20905b81548152906001019060200180831161015d57829003601f168201915b5050505050815260200160018201805461019390610e47565b80601f01602080910402602001604051908101604052809291908181526020018280546101bf90610e47565b801561020a5780601f106101e15761010080835404028352916020019161020a565b820191905f5260205f20905b8154815290600101906020018083116101ed57829003601f168201915b5050505050815260200160028201805461022390610e47565b80601f016020809104026020016040519081016040528092919081815260200182805461024f90610e47565b801561029a5780601f106102715761010080835404028352916020019161029a565b820191905f5260205f20905b81548152906001019060200180831161027d57829003601f168201915b505050505081526020016003820180546102b390610e47565b80601f01602080910402602001604051908101604052809291908181526020018280546102df90610e47565b801561032a5780601f106103015761010080835404028352916020019161032a565b820191905f5260205f20905b81548152906001019060200180831161030d57829003601f168201915b5050505050815260200160048201805461034390610e47565b80601f016020809104026020016040519081016040528092919081815260200182805461036f90610e47565b80156103ba5780601f10610391576101008083540402835291602001916103ba565b820191905f5260205f20905b81548152906001019060200180831161039d57829003601f168201915b505050505081526020016005820180546103d390610e47565b80601f01602080910402602001604051908101604052809291908181526020018280546103ff90610e47565b801561044a5780601f106104215761010080835404028352916020019161044a565b820191905f5260205f20905b81548152906001019060200180831161042d57829003601f168201915b5050505050815250509050919050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461049090610e47565b80601f01602080910402602001604051908101604052809291908181526020018280546104bc90610e47565b80156105075780601f106104de57610100808354040283529160200191610507565b820191905f5260205f20905b8154815290600101906020018083116104ea57829003601f168201915b50505050509080600101805461051c90610e47565b80601f016020809104026020016040519081016040528092919081815260200182805461054890610e47565b80156105935780601f1061056a57610100808354040283529160200191610593565b820191905f5260205f20905b81548152906001019060200180831161057657829003601f168201915b5050505050908060020180546105a890610e47565b80601f01602080910402602001604051908101604052809291908181526020018280546105d490610e47565b801561061f5780601f106105f65761010080835404028352916020019161061f565b820191905f5260205f20905b81548152906001019060200180831161060257829003601f168201915b50505050509080600301805461063490610e47565b80601f016020809104026020016040519081016040528092919081815260200182805461066090610e47565b80156106ab5780601f10610682576101008083540402835291602001916106ab565b820191905f5260205f20905b81548152906001019060200180831161068e57829003601f168201915b5050505050908060040180546106c090610e47565b80601f01602080910402602001604051908101604052809291908181526020018280546106ec90610e47565b80156107375780601f1061070e57610100808354040283529160200191610737565b820191905f5260205f20905b81548152906001019060200180831161071a57829003601f168201915b50505050509080600501805461074c90610e47565b80601f016020809104026020016040519081016040528092919081815260200182805461077890610e47565b80156107c35780601f1061079a576101008083540402835291602001916107c3565b820191905f5260205f20905b8154815290600101906020018083116107a657829003601f168201915b5050505050905086565b845f876040516107dd9190610e04565b908152602001604051809103902060010190816107fa919061101d565b50835f8760405161080b9190610e04565b90815260200160405180910390206002019081610828919061101d565b50825f876040516108399190610e04565b90815260200160405180910390206003019081610856919061101d565b50815f876040516108679190610e04565b90815260200160405180910390206004019081610884919061101d565b50805f876040516108959190610e04565b908152602001604051809103902060050190816108b2919061101d565b50505050505050565b6040518060c001604052806060815260200160608152602001606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6109508261090a565b810181811067ffffffffffffffff8211171561096f5761096e61091a565b5b80604052505050565b5f6109816108f1565b905061098d8282610947565b919050565b5f67ffffffffffffffff8211156109ac576109ab61091a565b5b6109b58261090a565b9050602081019050919050565b828183375f83830152505050565b5f6109e26109dd84610992565b610978565b9050828152602081018484840111156109fe576109fd610906565b5b610a098482856109c2565b509392505050565b5f82601f830112610a2557610a24610902565b5b8135610a358482602086016109d0565b91505092915050565b5f60208284031215610a5357610a526108fa565b5b5f82013567ffffffffffffffff811115610a7057610a6f6108fe565b5b610a7c84828501610a11565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610abc578082015181840152602081019050610aa1565b5f8484015250505050565b5f610ad182610a85565b610adb8185610a8f565b9350610aeb818560208601610a9f565b610af48161090a565b840191505092915050565b5f60c083015f8301518482035f860152610b198282610ac7565b91505060208301518482036020860152610b338282610ac7565b91505060408301518482036040860152610b4d8282610ac7565b91505060608301518482036060860152610b678282610ac7565b91505060808301518482036080860152610b818282610ac7565b91505060a083015184820360a0860152610b9b8282610ac7565b9150508091505092915050565b5f6020820190508181035f830152610bc08184610aff565b905092915050565b5f82825260208201905092915050565b5f610be282610a85565b610bec8185610bc8565b9350610bfc818560208601610a9f565b610c058161090a565b840191505092915050565b5f60c0820190508181035f830152610c288189610bd8565b90508181036020830152610c3c8188610bd8565b90508181036040830152610c508187610bd8565b90508181036060830152610c648186610bd8565b90508181036080830152610c788185610bd8565b905081810360a0830152610c8c8184610bd8565b9050979650505050505050565b5f805f805f8060c08789031215610cb357610cb26108fa565b5b5f87013567ffffffffffffffff811115610cd057610ccf6108fe565b5b610cdc89828a01610a11565b965050602087013567ffffffffffffffff811115610cfd57610cfc6108fe565b5b610d0989828a01610a11565b955050604087013567ffffffffffffffff811115610d2a57610d296108fe565b5b610d3689828a01610a11565b945050606087013567ffffffffffffffff811115610d5757610d566108fe565b5b610d6389828a01610a11565b935050608087013567ffffffffffffffff811115610d8457610d836108fe565b5b610d9089828a01610a11565b92505060a087013567ffffffffffffffff811115610db157610db06108fe565b5b610dbd89828a01610a11565b9150509295509295509295565b5f81905092915050565b5f610dde82610a85565b610de88185610dca565b9350610df8818560208601610a9f565b80840191505092915050565b5f610e0f8284610dd4565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610e5e57607f821691505b602082108103610e7157610e70610e1a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ed37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e98565b610edd8683610e98565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610f21610f1c610f1784610ef5565b610efe565b610ef5565b9050919050565b5f819050919050565b610f3a83610f07565b610f4e610f4682610f28565b848454610ea4565b825550505050565b5f90565b610f62610f56565b610f6d818484610f31565b505050565b5b81811015610f9057610f855f82610f5a565b600181019050610f73565b5050565b601f821115610fd557610fa681610e77565b610faf84610e89565b81016020851015610fbe578190505b610fd2610fca85610e89565b830182610f72565b50505b505050565b5f82821c905092915050565b5f610ff55f1984600802610fda565b1980831691505092915050565b5f61100d8383610fe6565b9150826002028217905092915050565b61102682610a85565b67ffffffffffffffff81111561103f5761103e61091a565b5b6110498254610e47565b611054828285610f94565b5f60209050601f831160018114611085575f8415611073578287015190505b61107d8582611002565b8655506110e4565b601f19841661109386610e77565b5f5b828110156110ba57848901518255600182019150602085019450602081019050611095565b868310156110d757848901516110d3601f891682610fe6565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220731cd340da8c597cf0b2179566f20ea9056d801393963ea344cdd3bc19a3930964736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend)
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

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string number, string workamount, string persion, string workmethod, string worktime, string remarks)
func (_Trace *TraceCaller) Tracedata(opts *bind.CallOpts, arg0 string) (struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedata", arg0)

	outstruct := new(struct {
		Number     string
		Workamount string
		Persion    string
		Workmethod string
		Worktime   string
		Remarks    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Number = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Workamount = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Persion = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Workmethod = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Worktime = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Remarks = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string number, string workamount, string persion, string workmethod, string worktime, string remarks)
func (_Trace *TraceSession) Tracedata(arg0 string) (struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string number, string workamount, string persion, string workmethod, string worktime, string remarks)
func (_Trace *TraceCallerSession) Tracedata(arg0 string) (struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// TracedataByNum is a free data retrieval call binding the contract method 0x1bb2bfea.
//
// Solidity: function tracedataByNum(string number) view returns((string,string,string,string,string,string))
func (_Trace *TraceCaller) TracedataByNum(opts *bind.CallOpts, number string) (UserPassStoreOrgInfo, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedataByNum", number)

	if err != nil {
		return *new(UserPassStoreOrgInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UserPassStoreOrgInfo)).(*UserPassStoreOrgInfo)

	return out0, err

}

// TracedataByNum is a free data retrieval call binding the contract method 0x1bb2bfea.
//
// Solidity: function tracedataByNum(string number) view returns((string,string,string,string,string,string))
func (_Trace *TraceSession) TracedataByNum(number string) (UserPassStoreOrgInfo, error) {
	return _Trace.Contract.TracedataByNum(&_Trace.CallOpts, number)
}

// TracedataByNum is a free data retrieval call binding the contract method 0x1bb2bfea.
//
// Solidity: function tracedataByNum(string number) view returns((string,string,string,string,string,string))
func (_Trace *TraceCallerSession) TracedataByNum(number string) (UserPassStoreOrgInfo, error) {
	return _Trace.Contract.TracedataByNum(&_Trace.CallOpts, number)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x8a205085.
//
// Solidity: function addorupdateData(string number, string workamount, string persion, string workmethod, string worktime, string remarks) returns()
func (_Trace *TraceTransactor) AddorupdateData(opts *bind.TransactOpts, number string, workamount string, persion string, workmethod string, worktime string, remarks string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addorupdateData", number, workamount, persion, workmethod, worktime, remarks)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x8a205085.
//
// Solidity: function addorupdateData(string number, string workamount, string persion, string workmethod, string worktime, string remarks) returns()
func (_Trace *TraceSession) AddorupdateData(number string, workamount string, persion string, workmethod string, worktime string, remarks string) (*types.Transaction, error) {
	return _Trace.Contract.AddorupdateData(&_Trace.TransactOpts, number, workamount, persion, workmethod, worktime, remarks)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x8a205085.
//
// Solidity: function addorupdateData(string number, string workamount, string persion, string workmethod, string worktime, string remarks) returns()
func (_Trace *TraceTransactorSession) AddorupdateData(number string, workamount string, persion string, workmethod string, worktime string, remarks string) (*types.Transaction, error) {
	return _Trace.Contract.AddorupdateData(&_Trace.TransactOpts, number, workamount, persion, workmethod, worktime, remarks)
}
