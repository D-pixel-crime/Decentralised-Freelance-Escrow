// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package escrow

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

// FreelanceEscrowMetaData contains all meta data concerning the FreelanceEscrow contract.
var FreelanceEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"freelancer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"confirmationPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_noCancel\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_noDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyArbitrator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyClient\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyFreelancer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptJobCompletion\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addClientStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addfreelancerStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"breakDeal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDealBreak\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finaliseUnilateralJob\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClientStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfirmationPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFreelancerStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectJobCompletion\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestPayment\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"clientPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"freelancerPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FreelanceEscrow__AggreementCreated\",\"inputs\":[{\"name\":\"client\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"freelancer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__BothPartyStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__ClientStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DealBroken\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DealCancelRequested\",\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DisputeResolved\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__FreelancerStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__JobCompletedAndFreelancerPaid\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__JobCompletionRejected\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__PaymentDisputeRaised\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__RandomDisputeRaised\",\"inputs\":[{\"name\":\"raiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__RevertedDealBreak\",\"inputs\":[{\"name\":\"reverter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ActiveConfirmationTimePeriod\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ClientAlreadyStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__FreelancerAlreadyStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidClient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidFreelancer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidFundsDistribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__NoCancelRequestedYet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__PaymentError\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ProcessNotAllowed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__RefundError\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreenlanceEscrow__Busy\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreenlanceEscrow__DealAlreadyBroken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// FreelanceEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use FreelanceEscrowMetaData.ABI instead.
var FreelanceEscrowABI = FreelanceEscrowMetaData.ABI

// FreelanceEscrow is an auto generated Go binding around an Ethereum contract.
type FreelanceEscrow struct {
	FreelanceEscrowCaller     // Read-only binding to the contract
	FreelanceEscrowTransactor // Write-only binding to the contract
	FreelanceEscrowFilterer   // Log filterer for contract events
}

// FreelanceEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreelanceEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreelanceEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreelanceEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreelanceEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreelanceEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreelanceEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreelanceEscrowSession struct {
	Contract     *FreelanceEscrow  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreelanceEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreelanceEscrowCallerSession struct {
	Contract *FreelanceEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FreelanceEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreelanceEscrowTransactorSession struct {
	Contract     *FreelanceEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FreelanceEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreelanceEscrowRaw struct {
	Contract *FreelanceEscrow // Generic contract binding to access the raw methods on
}

// FreelanceEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreelanceEscrowCallerRaw struct {
	Contract *FreelanceEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// FreelanceEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreelanceEscrowTransactorRaw struct {
	Contract *FreelanceEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreelanceEscrow creates a new instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrow(address common.Address, backend bind.ContractBackend) (*FreelanceEscrow, error) {
	contract, err := bindFreelanceEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrow{FreelanceEscrowCaller: FreelanceEscrowCaller{contract: contract}, FreelanceEscrowTransactor: FreelanceEscrowTransactor{contract: contract}, FreelanceEscrowFilterer: FreelanceEscrowFilterer{contract: contract}}, nil
}

// NewFreelanceEscrowCaller creates a new read-only instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrowCaller(address common.Address, caller bind.ContractCaller) (*FreelanceEscrowCaller, error) {
	contract, err := bindFreelanceEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowCaller{contract: contract}, nil
}

// NewFreelanceEscrowTransactor creates a new write-only instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*FreelanceEscrowTransactor, error) {
	contract, err := bindFreelanceEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowTransactor{contract: contract}, nil
}

// NewFreelanceEscrowFilterer creates a new log filterer instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*FreelanceEscrowFilterer, error) {
	contract, err := bindFreelanceEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFilterer{contract: contract}, nil
}

// bindFreelanceEscrow binds a generic wrapper to an already deployed contract.
func bindFreelanceEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FreelanceEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreelanceEscrow *FreelanceEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreelanceEscrow.Contract.FreelanceEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreelanceEscrow *FreelanceEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FreelanceEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreelanceEscrow *FreelanceEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FreelanceEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreelanceEscrow *FreelanceEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreelanceEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreelanceEscrow *FreelanceEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreelanceEscrow *FreelanceEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.contract.Transact(opts, method, params...)
}

// NoCancel is a free data retrieval call binding the contract method 0xd31cdf89.
//
// Solidity: function _noCancel() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) NoCancel(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_noCancel")

	if err != nil {
		return err
	}

	return err

}

// NoCancel is a free data retrieval call binding the contract method 0xd31cdf89.
//
// Solidity: function _noCancel() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) NoCancel() error {
	return _FreelanceEscrow.Contract.NoCancel(&_FreelanceEscrow.CallOpts)
}

// NoCancel is a free data retrieval call binding the contract method 0xd31cdf89.
//
// Solidity: function _noCancel() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) NoCancel() error {
	return _FreelanceEscrow.Contract.NoCancel(&_FreelanceEscrow.CallOpts)
}

// NoDispute is a free data retrieval call binding the contract method 0x95744869.
//
// Solidity: function _noDispute() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) NoDispute(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_noDispute")

	if err != nil {
		return err
	}

	return err

}

// NoDispute is a free data retrieval call binding the contract method 0x95744869.
//
// Solidity: function _noDispute() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) NoDispute() error {
	return _FreelanceEscrow.Contract.NoDispute(&_FreelanceEscrow.CallOpts)
}

// NoDispute is a free data retrieval call binding the contract method 0x95744869.
//
// Solidity: function _noDispute() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) NoDispute() error {
	return _FreelanceEscrow.Contract.NoDispute(&_FreelanceEscrow.CallOpts)
}

// OnlyArbitrator is a free data retrieval call binding the contract method 0x7f25e289.
//
// Solidity: function _onlyArbitrator() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) OnlyArbitrator(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_onlyArbitrator")

	if err != nil {
		return err
	}

	return err

}

// OnlyArbitrator is a free data retrieval call binding the contract method 0x7f25e289.
//
// Solidity: function _onlyArbitrator() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) OnlyArbitrator() error {
	return _FreelanceEscrow.Contract.OnlyArbitrator(&_FreelanceEscrow.CallOpts)
}

// OnlyArbitrator is a free data retrieval call binding the contract method 0x7f25e289.
//
// Solidity: function _onlyArbitrator() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) OnlyArbitrator() error {
	return _FreelanceEscrow.Contract.OnlyArbitrator(&_FreelanceEscrow.CallOpts)
}

// OnlyClient is a free data retrieval call binding the contract method 0x08e0771e.
//
// Solidity: function _onlyClient() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) OnlyClient(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_onlyClient")

	if err != nil {
		return err
	}

	return err

}

// OnlyClient is a free data retrieval call binding the contract method 0x08e0771e.
//
// Solidity: function _onlyClient() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) OnlyClient() error {
	return _FreelanceEscrow.Contract.OnlyClient(&_FreelanceEscrow.CallOpts)
}

// OnlyClient is a free data retrieval call binding the contract method 0x08e0771e.
//
// Solidity: function _onlyClient() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) OnlyClient() error {
	return _FreelanceEscrow.Contract.OnlyClient(&_FreelanceEscrow.CallOpts)
}

// OnlyFreelancer is a free data retrieval call binding the contract method 0x3e216754.
//
// Solidity: function _onlyFreelancer() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) OnlyFreelancer(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_onlyFreelancer")

	if err != nil {
		return err
	}

	return err

}

// OnlyFreelancer is a free data retrieval call binding the contract method 0x3e216754.
//
// Solidity: function _onlyFreelancer() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) OnlyFreelancer() error {
	return _FreelanceEscrow.Contract.OnlyFreelancer(&_FreelanceEscrow.CallOpts)
}

// OnlyFreelancer is a free data retrieval call binding the contract method 0x3e216754.
//
// Solidity: function _onlyFreelancer() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) OnlyFreelancer() error {
	return _FreelanceEscrow.Contract.OnlyFreelancer(&_FreelanceEscrow.CallOpts)
}

// GetClientStake is a free data retrieval call binding the contract method 0x87d40b5a.
//
// Solidity: function getClientStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetClientStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getClientStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClientStake is a free data retrieval call binding the contract method 0x87d40b5a.
//
// Solidity: function getClientStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowSession) GetClientStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetClientStake(&_FreelanceEscrow.CallOpts)
}

// GetClientStake is a free data retrieval call binding the contract method 0x87d40b5a.
//
// Solidity: function getClientStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetClientStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetClientStake(&_FreelanceEscrow.CallOpts)
}

// GetConfirmationPeriod is a free data retrieval call binding the contract method 0x97c9522e.
//
// Solidity: function getConfirmationPeriod() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetConfirmationPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getConfirmationPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationPeriod is a free data retrieval call binding the contract method 0x97c9522e.
//
// Solidity: function getConfirmationPeriod() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowSession) GetConfirmationPeriod() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetConfirmationPeriod(&_FreelanceEscrow.CallOpts)
}

// GetConfirmationPeriod is a free data retrieval call binding the contract method 0x97c9522e.
//
// Solidity: function getConfirmationPeriod() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetConfirmationPeriod() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetConfirmationPeriod(&_FreelanceEscrow.CallOpts)
}

// GetEscrowState is a free data retrieval call binding the contract method 0xdf44be2c.
//
// Solidity: function getEscrowState() view returns(uint8)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetEscrowState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getEscrowState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEscrowState is a free data retrieval call binding the contract method 0xdf44be2c.
//
// Solidity: function getEscrowState() view returns(uint8)
func (_FreelanceEscrow *FreelanceEscrowSession) GetEscrowState() (uint8, error) {
	return _FreelanceEscrow.Contract.GetEscrowState(&_FreelanceEscrow.CallOpts)
}

// GetEscrowState is a free data retrieval call binding the contract method 0xdf44be2c.
//
// Solidity: function getEscrowState() view returns(uint8)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetEscrowState() (uint8, error) {
	return _FreelanceEscrow.Contract.GetEscrowState(&_FreelanceEscrow.CallOpts)
}

// GetFreelancerStake is a free data retrieval call binding the contract method 0xac755a63.
//
// Solidity: function getFreelancerStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetFreelancerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getFreelancerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFreelancerStake is a free data retrieval call binding the contract method 0xac755a63.
//
// Solidity: function getFreelancerStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowSession) GetFreelancerStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetFreelancerStake(&_FreelanceEscrow.CallOpts)
}

// GetFreelancerStake is a free data retrieval call binding the contract method 0xac755a63.
//
// Solidity: function getFreelancerStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetFreelancerStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetFreelancerStake(&_FreelanceEscrow.CallOpts)
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AcceptJobCompletion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "acceptJobCompletion")
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AcceptJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AcceptJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) AcceptJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AcceptJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// AddClientStake is a paid mutator transaction binding the contract method 0x939ae273.
//
// Solidity: function addClientStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AddClientStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "addClientStake")
}

// AddClientStake is a paid mutator transaction binding the contract method 0x939ae273.
//
// Solidity: function addClientStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AddClientStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddClientStake(&_FreelanceEscrow.TransactOpts)
}

// AddClientStake is a paid mutator transaction binding the contract method 0x939ae273.
//
// Solidity: function addClientStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) AddClientStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddClientStake(&_FreelanceEscrow.TransactOpts)
}

// AddfreelancerStake is a paid mutator transaction binding the contract method 0x24dcdf64.
//
// Solidity: function addfreelancerStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AddfreelancerStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "addfreelancerStake")
}

// AddfreelancerStake is a paid mutator transaction binding the contract method 0x24dcdf64.
//
// Solidity: function addfreelancerStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AddfreelancerStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddfreelancerStake(&_FreelanceEscrow.TransactOpts)
}

// AddfreelancerStake is a paid mutator transaction binding the contract method 0x24dcdf64.
//
// Solidity: function addfreelancerStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) AddfreelancerStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddfreelancerStake(&_FreelanceEscrow.TransactOpts)
}

// BreakDeal is a paid mutator transaction binding the contract method 0x6977e176.
//
// Solidity: function breakDeal() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) BreakDeal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "breakDeal")
}

// BreakDeal is a paid mutator transaction binding the contract method 0x6977e176.
//
// Solidity: function breakDeal() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) BreakDeal() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.BreakDeal(&_FreelanceEscrow.TransactOpts)
}

// BreakDeal is a paid mutator transaction binding the contract method 0x6977e176.
//
// Solidity: function breakDeal() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) BreakDeal() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.BreakDeal(&_FreelanceEscrow.TransactOpts)
}

// CancelDealBreak is a paid mutator transaction binding the contract method 0x75f84b40.
//
// Solidity: function cancelDealBreak() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) CancelDealBreak(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "cancelDealBreak")
}

// CancelDealBreak is a paid mutator transaction binding the contract method 0x75f84b40.
//
// Solidity: function cancelDealBreak() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) CancelDealBreak() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.CancelDealBreak(&_FreelanceEscrow.TransactOpts)
}

// CancelDealBreak is a paid mutator transaction binding the contract method 0x75f84b40.
//
// Solidity: function cancelDealBreak() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) CancelDealBreak() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.CancelDealBreak(&_FreelanceEscrow.TransactOpts)
}

// FinaliseUnilateralJob is a paid mutator transaction binding the contract method 0xdd33869f.
//
// Solidity: function finaliseUnilateralJob() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) FinaliseUnilateralJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "finaliseUnilateralJob")
}

// FinaliseUnilateralJob is a paid mutator transaction binding the contract method 0xdd33869f.
//
// Solidity: function finaliseUnilateralJob() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) FinaliseUnilateralJob() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FinaliseUnilateralJob(&_FreelanceEscrow.TransactOpts)
}

// FinaliseUnilateralJob is a paid mutator transaction binding the contract method 0xdd33869f.
//
// Solidity: function finaliseUnilateralJob() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) FinaliseUnilateralJob() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FinaliseUnilateralJob(&_FreelanceEscrow.TransactOpts)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x6daa2d44.
//
// Solidity: function raiseDispute() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) RaiseDispute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "raiseDispute")
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x6daa2d44.
//
// Solidity: function raiseDispute() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) RaiseDispute() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RaiseDispute(&_FreelanceEscrow.TransactOpts)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x6daa2d44.
//
// Solidity: function raiseDispute() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) RaiseDispute() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RaiseDispute(&_FreelanceEscrow.TransactOpts)
}

// RejectJobCompletion is a paid mutator transaction binding the contract method 0xee4b7ab8.
//
// Solidity: function rejectJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) RejectJobCompletion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "rejectJobCompletion")
}

// RejectJobCompletion is a paid mutator transaction binding the contract method 0xee4b7ab8.
//
// Solidity: function rejectJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) RejectJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RejectJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// RejectJobCompletion is a paid mutator transaction binding the contract method 0xee4b7ab8.
//
// Solidity: function rejectJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) RejectJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RejectJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x63bdb94b.
//
// Solidity: function requestPayment() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) RequestPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "requestPayment")
}

// RequestPayment is a paid mutator transaction binding the contract method 0x63bdb94b.
//
// Solidity: function requestPayment() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) RequestPayment() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RequestPayment(&_FreelanceEscrow.TransactOpts)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x63bdb94b.
//
// Solidity: function requestPayment() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) RequestPayment() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RequestPayment(&_FreelanceEscrow.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xbdc84ac3.
//
// Solidity: function resolveDispute(uint256 clientPayment, uint256 freelancerPayment) returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) ResolveDispute(opts *bind.TransactOpts, clientPayment *big.Int, freelancerPayment *big.Int) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "resolveDispute", clientPayment, freelancerPayment)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xbdc84ac3.
//
// Solidity: function resolveDispute(uint256 clientPayment, uint256 freelancerPayment) returns()
func (_FreelanceEscrow *FreelanceEscrowSession) ResolveDispute(clientPayment *big.Int, freelancerPayment *big.Int) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.ResolveDispute(&_FreelanceEscrow.TransactOpts, clientPayment, freelancerPayment)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xbdc84ac3.
//
// Solidity: function resolveDispute(uint256 clientPayment, uint256 freelancerPayment) returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) ResolveDispute(clientPayment *big.Int, freelancerPayment *big.Int) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.ResolveDispute(&_FreelanceEscrow.TransactOpts, clientPayment, freelancerPayment)
}

// FreelanceEscrowFreelanceEscrowAggreementCreatedIterator is returned from FilterFreelanceEscrowAggreementCreated and is used to iterate over the raw logs and unpacked data for FreelanceEscrowAggreementCreated events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowAggreementCreatedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowAggreementCreated // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowAggreementCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowAggreementCreated)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowAggreementCreated)
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
func (it *FreelanceEscrowFreelanceEscrowAggreementCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowAggreementCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowAggreementCreated represents a FreelanceEscrowAggreementCreated event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowAggreementCreated struct {
	Client     common.Address
	Freelancer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowAggreementCreated is a free log retrieval operation binding the contract event 0xdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f.
//
// Solidity: event FreelanceEscrow__AggreementCreated(address indexed client, address indexed freelancer)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowAggreementCreated(opts *bind.FilterOpts, client []common.Address, freelancer []common.Address) (*FreelanceEscrowFreelanceEscrowAggreementCreatedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var freelancerRule []interface{}
	for _, freelancerItem := range freelancer {
		freelancerRule = append(freelancerRule, freelancerItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__AggreementCreated", clientRule, freelancerRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowAggreementCreatedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__AggreementCreated", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowAggreementCreated is a free log subscription operation binding the contract event 0xdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f.
//
// Solidity: event FreelanceEscrow__AggreementCreated(address indexed client, address indexed freelancer)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowAggreementCreated(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowAggreementCreated, client []common.Address, freelancer []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var freelancerRule []interface{}
	for _, freelancerItem := range freelancer {
		freelancerRule = append(freelancerRule, freelancerItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__AggreementCreated", clientRule, freelancerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowAggreementCreated)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__AggreementCreated", log); err != nil {
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

// ParseFreelanceEscrowAggreementCreated is a log parse operation binding the contract event 0xdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f.
//
// Solidity: event FreelanceEscrow__AggreementCreated(address indexed client, address indexed freelancer)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowAggreementCreated(log types.Log) (*FreelanceEscrowFreelanceEscrowAggreementCreated, error) {
	event := new(FreelanceEscrowFreelanceEscrowAggreementCreated)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__AggreementCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator is returned from FilterFreelanceEscrowBothPartyStakeCompleted and is used to iterate over the raw logs and unpacked data for FreelanceEscrowBothPartyStakeCompleted events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
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
func (it *FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted represents a FreelanceEscrowBothPartyStakeCompleted event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowBothPartyStakeCompleted is a free log retrieval operation binding the contract event 0x84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da.
//
// Solidity: event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowBothPartyStakeCompleted(opts *bind.FilterOpts, timestamp []*big.Int) (*FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__BothPartyStakeCompleted", timestampRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__BothPartyStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowBothPartyStakeCompleted is a free log subscription operation binding the contract event 0x84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da.
//
// Solidity: event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowBothPartyStakeCompleted(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__BothPartyStakeCompleted", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__BothPartyStakeCompleted", log); err != nil {
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

// ParseFreelanceEscrowBothPartyStakeCompleted is a log parse operation binding the contract event 0x84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da.
//
// Solidity: event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowBothPartyStakeCompleted(log types.Log) (*FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted, error) {
	event := new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__BothPartyStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator is returned from FilterFreelanceEscrowClientStakeCompleted and is used to iterate over the raw logs and unpacked data for FreelanceEscrowClientStakeCompleted events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowClientStakeCompleted // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
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
func (it *FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowClientStakeCompleted represents a FreelanceEscrowClientStakeCompleted event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowClientStakeCompleted struct {
	Timestamp *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowClientStakeCompleted is a free log retrieval operation binding the contract event 0x218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc.
//
// Solidity: event FreelanceEscrow__ClientStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowClientStakeCompleted(opts *bind.FilterOpts, timestamp []*big.Int, amount []*big.Int) (*FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__ClientStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__ClientStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowClientStakeCompleted is a free log subscription operation binding the contract event 0x218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc.
//
// Solidity: event FreelanceEscrow__ClientStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowClientStakeCompleted(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowClientStakeCompleted, timestamp []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__ClientStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__ClientStakeCompleted", log); err != nil {
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

// ParseFreelanceEscrowClientStakeCompleted is a log parse operation binding the contract event 0x218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc.
//
// Solidity: event FreelanceEscrow__ClientStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowClientStakeCompleted(log types.Log) (*FreelanceEscrowFreelanceEscrowClientStakeCompleted, error) {
	event := new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__ClientStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowDealBrokenIterator is returned from FilterFreelanceEscrowDealBroken and is used to iterate over the raw logs and unpacked data for FreelanceEscrowDealBroken events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealBrokenIterator struct {
	Event *FreelanceEscrowFreelanceEscrowDealBroken // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowDealBrokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowDealBroken)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowDealBroken)
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
func (it *FreelanceEscrowFreelanceEscrowDealBrokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowDealBrokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowDealBroken represents a FreelanceEscrowDealBroken event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealBroken struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowDealBroken is a free log retrieval operation binding the contract event 0xa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf24.
//
// Solidity: event FreelanceEscrow__DealBroken(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowDealBroken(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowDealBrokenIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__DealBroken")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowDealBrokenIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__DealBroken", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowDealBroken is a free log subscription operation binding the contract event 0xa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf24.
//
// Solidity: event FreelanceEscrow__DealBroken(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowDealBroken(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowDealBroken) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__DealBroken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowDealBroken)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealBroken", log); err != nil {
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

// ParseFreelanceEscrowDealBroken is a log parse operation binding the contract event 0xa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf24.
//
// Solidity: event FreelanceEscrow__DealBroken(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowDealBroken(log types.Log) (*FreelanceEscrowFreelanceEscrowDealBroken, error) {
	event := new(FreelanceEscrowFreelanceEscrowDealBroken)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealBroken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator is returned from FilterFreelanceEscrowDealCancelRequested and is used to iterate over the raw logs and unpacked data for FreelanceEscrowDealCancelRequested events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowDealCancelRequested // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
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
func (it *FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowDealCancelRequested represents a FreelanceEscrowDealCancelRequested event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealCancelRequested struct {
	Initiator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowDealCancelRequested is a free log retrieval operation binding the contract event 0xf0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce.
//
// Solidity: event FreelanceEscrow__DealCancelRequested(address indexed initiator)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowDealCancelRequested(opts *bind.FilterOpts, initiator []common.Address) (*FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__DealCancelRequested", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__DealCancelRequested", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowDealCancelRequested is a free log subscription operation binding the contract event 0xf0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce.
//
// Solidity: event FreelanceEscrow__DealCancelRequested(address indexed initiator)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowDealCancelRequested(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowDealCancelRequested, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__DealCancelRequested", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealCancelRequested", log); err != nil {
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

// ParseFreelanceEscrowDealCancelRequested is a log parse operation binding the contract event 0xf0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce.
//
// Solidity: event FreelanceEscrow__DealCancelRequested(address indexed initiator)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowDealCancelRequested(log types.Log) (*FreelanceEscrowFreelanceEscrowDealCancelRequested, error) {
	event := new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealCancelRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowDisputeResolvedIterator is returned from FilterFreelanceEscrowDisputeResolved and is used to iterate over the raw logs and unpacked data for FreelanceEscrowDisputeResolved events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDisputeResolvedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowDisputeResolved // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowDisputeResolved)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowDisputeResolved)
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
func (it *FreelanceEscrowFreelanceEscrowDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowDisputeResolved represents a FreelanceEscrowDisputeResolved event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDisputeResolved struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowDisputeResolved is a free log retrieval operation binding the contract event 0x0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e.
//
// Solidity: event FreelanceEscrow__DisputeResolved(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowDisputeResolved(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowDisputeResolvedIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowDisputeResolvedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowDisputeResolved is a free log subscription operation binding the contract event 0x0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e.
//
// Solidity: event FreelanceEscrow__DisputeResolved(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowDisputeResolved(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowDisputeResolved)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DisputeResolved", log); err != nil {
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

// ParseFreelanceEscrowDisputeResolved is a log parse operation binding the contract event 0x0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e.
//
// Solidity: event FreelanceEscrow__DisputeResolved(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowDisputeResolved(log types.Log) (*FreelanceEscrowFreelanceEscrowDisputeResolved, error) {
	event := new(FreelanceEscrowFreelanceEscrowDisputeResolved)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator is returned from FilterFreelanceEscrowFreelancerCompletedAndClientConfirmationPending and is used to iterate over the raw logs and unpacked data for FreelanceEscrowFreelancerCompletedAndClientConfirmationPending events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator struct {
	Event *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
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
func (it *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending represents a FreelanceEscrowFreelancerCompletedAndClientConfirmationPending event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowFreelancerCompletedAndClientConfirmationPending is a free log retrieval operation binding the contract event 0xfbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6.
//
// Solidity: event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowFreelancerCompletedAndClientConfirmationPending is a free log subscription operation binding the contract event 0xfbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6.
//
// Solidity: event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending", log); err != nil {
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

// ParseFreelanceEscrowFreelancerCompletedAndClientConfirmationPending is a log parse operation binding the contract event 0xfbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6.
//
// Solidity: event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(log types.Log) (*FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending, error) {
	event := new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator is returned from FilterFreelanceEscrowFreelancerStakeCompleted and is used to iterate over the raw logs and unpacked data for FreelanceEscrowFreelancerStakeCompleted events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
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
func (it *FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted represents a FreelanceEscrowFreelancerStakeCompleted event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted struct {
	Timestamp *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowFreelancerStakeCompleted is a free log retrieval operation binding the contract event 0x4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e.
//
// Solidity: event FreelanceEscrow__FreelancerStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowFreelancerStakeCompleted(opts *bind.FilterOpts, timestamp []*big.Int, amount []*big.Int) (*FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__FreelancerStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__FreelancerStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowFreelancerStakeCompleted is a free log subscription operation binding the contract event 0x4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e.
//
// Solidity: event FreelanceEscrow__FreelancerStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowFreelancerStakeCompleted(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted, timestamp []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__FreelancerStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerStakeCompleted", log); err != nil {
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

// ParseFreelanceEscrowFreelancerStakeCompleted is a log parse operation binding the contract event 0x4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e.
//
// Solidity: event FreelanceEscrow__FreelancerStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowFreelancerStakeCompleted(log types.Log) (*FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted, error) {
	event := new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator is returned from FilterFreelanceEscrowJobCompletedAndFreelancerPaid and is used to iterate over the raw logs and unpacked data for FreelanceEscrowJobCompletedAndFreelancerPaid events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator struct {
	Event *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
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
func (it *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid represents a FreelanceEscrowJobCompletedAndFreelancerPaid event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid struct {
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowJobCompletedAndFreelancerPaid is a free log retrieval operation binding the contract event 0x4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e981.
//
// Solidity: event FreelanceEscrow__JobCompletedAndFreelancerPaid(uint256 indexed amount, uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowJobCompletedAndFreelancerPaid(opts *bind.FilterOpts, amount []*big.Int, timestamp []*big.Int) (*FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__JobCompletedAndFreelancerPaid", amountRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__JobCompletedAndFreelancerPaid", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowJobCompletedAndFreelancerPaid is a free log subscription operation binding the contract event 0x4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e981.
//
// Solidity: event FreelanceEscrow__JobCompletedAndFreelancerPaid(uint256 indexed amount, uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowJobCompletedAndFreelancerPaid(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid, amount []*big.Int, timestamp []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__JobCompletedAndFreelancerPaid", amountRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletedAndFreelancerPaid", log); err != nil {
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

// ParseFreelanceEscrowJobCompletedAndFreelancerPaid is a log parse operation binding the contract event 0x4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e981.
//
// Solidity: event FreelanceEscrow__JobCompletedAndFreelancerPaid(uint256 indexed amount, uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowJobCompletedAndFreelancerPaid(log types.Log) (*FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid, error) {
	event := new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletedAndFreelancerPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator is returned from FilterFreelanceEscrowJobCompletionRejected and is used to iterate over the raw logs and unpacked data for FreelanceEscrowJobCompletionRejected events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowJobCompletionRejected // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
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
func (it *FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowJobCompletionRejected represents a FreelanceEscrowJobCompletionRejected event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletionRejected struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowJobCompletionRejected is a free log retrieval operation binding the contract event 0x52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd.
//
// Solidity: event FreelanceEscrow__JobCompletionRejected(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowJobCompletionRejected(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__JobCompletionRejected")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__JobCompletionRejected", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowJobCompletionRejected is a free log subscription operation binding the contract event 0x52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd.
//
// Solidity: event FreelanceEscrow__JobCompletionRejected(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowJobCompletionRejected(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowJobCompletionRejected) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__JobCompletionRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletionRejected", log); err != nil {
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

// ParseFreelanceEscrowJobCompletionRejected is a log parse operation binding the contract event 0x52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd.
//
// Solidity: event FreelanceEscrow__JobCompletionRejected(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowJobCompletionRejected(log types.Log) (*FreelanceEscrowFreelanceEscrowJobCompletionRejected, error) {
	event := new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator is returned from FilterFreelanceEscrowPaymentDisputeRaised and is used to iterate over the raw logs and unpacked data for FreelanceEscrowPaymentDisputeRaised events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowPaymentDisputeRaised // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
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
func (it *FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowPaymentDisputeRaised represents a FreelanceEscrowPaymentDisputeRaised event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowPaymentDisputeRaised struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowPaymentDisputeRaised is a free log retrieval operation binding the contract event 0x7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e.
//
// Solidity: event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowPaymentDisputeRaised(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__PaymentDisputeRaised")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__PaymentDisputeRaised", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowPaymentDisputeRaised is a free log subscription operation binding the contract event 0x7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e.
//
// Solidity: event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowPaymentDisputeRaised(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowPaymentDisputeRaised) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__PaymentDisputeRaised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__PaymentDisputeRaised", log); err != nil {
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

// ParseFreelanceEscrowPaymentDisputeRaised is a log parse operation binding the contract event 0x7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e.
//
// Solidity: event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowPaymentDisputeRaised(log types.Log) (*FreelanceEscrowFreelanceEscrowPaymentDisputeRaised, error) {
	event := new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__PaymentDisputeRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator is returned from FilterFreelanceEscrowRandomDisputeRaised and is used to iterate over the raw logs and unpacked data for FreelanceEscrowRandomDisputeRaised events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowRandomDisputeRaised // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
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
func (it *FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowRandomDisputeRaised represents a FreelanceEscrowRandomDisputeRaised event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRandomDisputeRaised struct {
	Raiser    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowRandomDisputeRaised is a free log retrieval operation binding the contract event 0x022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991.
//
// Solidity: event FreelanceEscrow__RandomDisputeRaised(address indexed raiser, uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowRandomDisputeRaised(opts *bind.FilterOpts, raiser []common.Address) (*FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator, error) {

	var raiserRule []interface{}
	for _, raiserItem := range raiser {
		raiserRule = append(raiserRule, raiserItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__RandomDisputeRaised", raiserRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__RandomDisputeRaised", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowRandomDisputeRaised is a free log subscription operation binding the contract event 0x022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991.
//
// Solidity: event FreelanceEscrow__RandomDisputeRaised(address indexed raiser, uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowRandomDisputeRaised(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowRandomDisputeRaised, raiser []common.Address) (event.Subscription, error) {

	var raiserRule []interface{}
	for _, raiserItem := range raiser {
		raiserRule = append(raiserRule, raiserItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__RandomDisputeRaised", raiserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RandomDisputeRaised", log); err != nil {
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

// ParseFreelanceEscrowRandomDisputeRaised is a log parse operation binding the contract event 0x022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991.
//
// Solidity: event FreelanceEscrow__RandomDisputeRaised(address indexed raiser, uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowRandomDisputeRaised(log types.Log) (*FreelanceEscrowFreelanceEscrowRandomDisputeRaised, error) {
	event := new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RandomDisputeRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator is returned from FilterFreelanceEscrowRevertedDealBreak and is used to iterate over the raw logs and unpacked data for FreelanceEscrowRevertedDealBreak events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator struct {
	Event *FreelanceEscrowFreelanceEscrowRevertedDealBreak // Event containing the contract specifics and raw log

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
func (it *FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
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
		it.Event = new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
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
func (it *FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowRevertedDealBreak represents a FreelanceEscrowRevertedDealBreak event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRevertedDealBreak struct {
	Reverter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowRevertedDealBreak is a free log retrieval operation binding the contract event 0x9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a8.
//
// Solidity: event FreelanceEscrow__RevertedDealBreak(address indexed reverter)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowRevertedDealBreak(opts *bind.FilterOpts, reverter []common.Address) (*FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator, error) {

	var reverterRule []interface{}
	for _, reverterItem := range reverter {
		reverterRule = append(reverterRule, reverterItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__RevertedDealBreak", reverterRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__RevertedDealBreak", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowRevertedDealBreak is a free log subscription operation binding the contract event 0x9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a8.
//
// Solidity: event FreelanceEscrow__RevertedDealBreak(address indexed reverter)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowRevertedDealBreak(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowRevertedDealBreak, reverter []common.Address) (event.Subscription, error) {

	var reverterRule []interface{}
	for _, reverterItem := range reverter {
		reverterRule = append(reverterRule, reverterItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__RevertedDealBreak", reverterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RevertedDealBreak", log); err != nil {
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

// ParseFreelanceEscrowRevertedDealBreak is a log parse operation binding the contract event 0x9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a8.
//
// Solidity: event FreelanceEscrow__RevertedDealBreak(address indexed reverter)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowRevertedDealBreak(log types.Log) (*FreelanceEscrowFreelanceEscrowRevertedDealBreak, error) {
	event := new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RevertedDealBreak", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
