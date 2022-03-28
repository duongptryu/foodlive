// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package foodlive

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
)

// FoodliveMetaData contains all meta data concerning the Foodlive contract.
var FoodliveMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PaymentOrder\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"paymentOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101aa806100206000396000f3fe6080604052600436106100295760003560e01c806312065fe01461002e5780631247e91d14610059575b600080fd5b34801561003a57600080fd5b50610043610075565b60405161005091906100d2565b60405180910390f35b610073600480360381019061006e919061011e565b61007d565b005b600047905090565b7f2d4515b4d2e9e432433e3ef9944a92f60d5f2e3213d302cfdb197e6f16782a3c81346040516100ae92919061014b565b60405180910390a150565b6000819050919050565b6100cc816100b9565b82525050565b60006020820190506100e760008301846100c3565b92915050565b600080fd5b6100fb816100b9565b811461010657600080fd5b50565b600081359050610118816100f2565b92915050565b600060208284031215610134576101336100ed565b5b600061014284828501610109565b91505092915050565b600060408201905061016060008301856100c3565b61016d60208301846100c3565b939250505056fea2646970667358221220de0db3d6c407a5b99c80a72126ceaa04e2285d345cfae1782cb768689fa2b81e64736f6c634300080d0033",
}

// FoodliveABI is the input ABI used to generate the binding from.
// Deprecated: Use FoodliveMetaData.ABI instead.
var FoodliveABI = FoodliveMetaData.ABI

// FoodliveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FoodliveMetaData.Bin instead.
var FoodliveBin = FoodliveMetaData.Bin

// DeployFoodlive deploys a new Ethereum contract, binding an instance of Foodlive to it.
func DeployFoodlive(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Foodlive, error) {
	parsed, err := FoodliveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FoodliveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Foodlive{FoodliveCaller: FoodliveCaller{contract: contract}, FoodliveTransactor: FoodliveTransactor{contract: contract}, FoodliveFilterer: FoodliveFilterer{contract: contract}}, nil
}

// Foodlive is an auto generated Go binding around an Ethereum contract.
type Foodlive struct {
	FoodliveCaller     // Read-only binding to the contract
	FoodliveTransactor // Write-only binding to the contract
	FoodliveFilterer   // Log filterer for contract events
}

// FoodliveCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoodliveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodliveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoodliveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodliveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoodliveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodliveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoodliveSession struct {
	Contract     *Foodlive         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoodliveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoodliveCallerSession struct {
	Contract *FoodliveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FoodliveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoodliveTransactorSession struct {
	Contract     *FoodliveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FoodliveRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoodliveRaw struct {
	Contract *Foodlive // Generic contract binding to access the raw methods on
}

// FoodliveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoodliveCallerRaw struct {
	Contract *FoodliveCaller // Generic read-only contract binding to access the raw methods on
}

// FoodliveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoodliveTransactorRaw struct {
	Contract *FoodliveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoodlive creates a new instance of Foodlive, bound to a specific deployed contract.
func NewFoodlive(address common.Address, backend bind.ContractBackend) (*Foodlive, error) {
	contract, err := bindFoodlive(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foodlive{FoodliveCaller: FoodliveCaller{contract: contract}, FoodliveTransactor: FoodliveTransactor{contract: contract}, FoodliveFilterer: FoodliveFilterer{contract: contract}}, nil
}

// NewFoodliveCaller creates a new read-only instance of Foodlive, bound to a specific deployed contract.
func NewFoodliveCaller(address common.Address, caller bind.ContractCaller) (*FoodliveCaller, error) {
	contract, err := bindFoodlive(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoodliveCaller{contract: contract}, nil
}

// NewFoodliveTransactor creates a new write-only instance of Foodlive, bound to a specific deployed contract.
func NewFoodliveTransactor(address common.Address, transactor bind.ContractTransactor) (*FoodliveTransactor, error) {
	contract, err := bindFoodlive(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoodliveTransactor{contract: contract}, nil
}

// NewFoodliveFilterer creates a new log filterer instance of Foodlive, bound to a specific deployed contract.
func NewFoodliveFilterer(address common.Address, filterer bind.ContractFilterer) (*FoodliveFilterer, error) {
	contract, err := bindFoodlive(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoodliveFilterer{contract: contract}, nil
}

// bindFoodlive binds a generic wrapper to an already deployed contract.
func bindFoodlive(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FoodliveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foodlive *FoodliveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foodlive.Contract.FoodliveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foodlive *FoodliveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foodlive.Contract.FoodliveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foodlive *FoodliveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foodlive.Contract.FoodliveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foodlive *FoodliveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foodlive.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foodlive *FoodliveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foodlive.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foodlive *FoodliveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foodlive.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Foodlive *FoodliveCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Foodlive.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Foodlive *FoodliveSession) GetBalance() (*big.Int, error) {
	return _Foodlive.Contract.GetBalance(&_Foodlive.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Foodlive *FoodliveCallerSession) GetBalance() (*big.Int, error) {
	return _Foodlive.Contract.GetBalance(&_Foodlive.CallOpts)
}

// PaymentOrder is a paid mutator transaction binding the contract method 0x1247e91d.
//
// Solidity: function paymentOrder(uint256 _orderId) payable returns()
func (_Foodlive *FoodliveTransactor) PaymentOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _Foodlive.contract.Transact(opts, "paymentOrder", _orderId)
}

// PaymentOrder is a paid mutator transaction binding the contract method 0x1247e91d.
//
// Solidity: function paymentOrder(uint256 _orderId) payable returns()
func (_Foodlive *FoodliveSession) PaymentOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _Foodlive.Contract.PaymentOrder(&_Foodlive.TransactOpts, _orderId)
}

// PaymentOrder is a paid mutator transaction binding the contract method 0x1247e91d.
//
// Solidity: function paymentOrder(uint256 _orderId) payable returns()
func (_Foodlive *FoodliveTransactorSession) PaymentOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _Foodlive.Contract.PaymentOrder(&_Foodlive.TransactOpts, _orderId)
}

// FoodlivePaymentOrderIterator is returned from FilterPaymentOrder and is used to iterate over the raw logs and unpacked data for PaymentOrder events raised by the Foodlive contract.
type FoodlivePaymentOrderIterator struct {
	Event *FoodlivePaymentOrder // Event containing the contract specifics and raw log

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
func (it *FoodlivePaymentOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodlivePaymentOrder)
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
		it.Event = new(FoodlivePaymentOrder)
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
func (it *FoodlivePaymentOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodlivePaymentOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodlivePaymentOrder represents a PaymentOrder event raised by the Foodlive contract.
type FoodlivePaymentOrder struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaymentOrder is a free log retrieval operation binding the contract event 0x2d4515b4d2e9e432433e3ef9944a92f60d5f2e3213d302cfdb197e6f16782a3c.
//
// Solidity: event PaymentOrder(uint256 arg0, uint256 arg1)
func (_Foodlive *FoodliveFilterer) FilterPaymentOrder(opts *bind.FilterOpts) (*FoodlivePaymentOrderIterator, error) {

	logs, sub, err := _Foodlive.contract.FilterLogs(opts, "PaymentOrder")
	if err != nil {
		return nil, err
	}
	return &FoodlivePaymentOrderIterator{contract: _Foodlive.contract, event: "PaymentOrder", logs: logs, sub: sub}, nil
}

// WatchPaymentOrder is a free log subscription operation binding the contract event 0x2d4515b4d2e9e432433e3ef9944a92f60d5f2e3213d302cfdb197e6f16782a3c.
//
// Solidity: event PaymentOrder(uint256 arg0, uint256 arg1)
func (_Foodlive *FoodliveFilterer) WatchPaymentOrder(opts *bind.WatchOpts, sink chan<- *FoodlivePaymentOrder) (event.Subscription, error) {

	logs, sub, err := _Foodlive.contract.WatchLogs(opts, "PaymentOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodlivePaymentOrder)
				if err := _Foodlive.contract.UnpackLog(event, "PaymentOrder", log); err != nil {
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

// ParsePaymentOrder is a log parse operation binding the contract event 0x2d4515b4d2e9e432433e3ef9944a92f60d5f2e3213d302cfdb197e6f16782a3c.
//
// Solidity: event PaymentOrder(uint256 arg0, uint256 arg1)
func (_Foodlive *FoodliveFilterer) ParsePaymentOrder(log types.Log) (*FoodlivePaymentOrder, error) {
	event := new(FoodlivePaymentOrder)
	if err := _Foodlive.contract.UnpackLog(event, "PaymentOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
