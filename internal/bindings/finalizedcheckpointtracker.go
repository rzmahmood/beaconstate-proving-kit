// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IMerkleProofVerifierProofInputs is an auto generated low-level Go binding around an user-defined struct.
type IMerkleProofVerifierProofInputs struct {
	Index  *big.Int
	Branch [][32]byte
	Value  [32]byte
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"beaconRootsAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconRootsAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestFinalizedEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestFinalizedTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveCheckpointFinalized\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIMerkleProofVerifier.ProofInputs\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"IMerkleProofVerifier.MerkleRoot\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIMerkleProofVerifier.ProofInputs\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"FinalizedCheckpointUpdated\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BeaconRootsCallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BranchHasExtraItems\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BranchIsMissingItems\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootMismatch\",\"inputs\":[]}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// BeaconRootsAddress is a free data retrieval call binding the contract method 0x38c8c9de.
//
// Solidity: function beaconRootsAddress() view returns(address)
func (_Bindings *BindingsCaller) BeaconRootsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "beaconRootsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconRootsAddress is a free data retrieval call binding the contract method 0x38c8c9de.
//
// Solidity: function beaconRootsAddress() view returns(address)
func (_Bindings *BindingsSession) BeaconRootsAddress() (common.Address, error) {
	return _Bindings.Contract.BeaconRootsAddress(&_Bindings.CallOpts)
}

// BeaconRootsAddress is a free data retrieval call binding the contract method 0x38c8c9de.
//
// Solidity: function beaconRootsAddress() view returns(address)
func (_Bindings *BindingsCallerSession) BeaconRootsAddress() (common.Address, error) {
	return _Bindings.Contract.BeaconRootsAddress(&_Bindings.CallOpts)
}

// HighestFinalizedEpoch is a free data retrieval call binding the contract method 0xe8b0d84c.
//
// Solidity: function highestFinalizedEpoch() view returns(uint256)
func (_Bindings *BindingsCaller) HighestFinalizedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "highestFinalizedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestFinalizedEpoch is a free data retrieval call binding the contract method 0xe8b0d84c.
//
// Solidity: function highestFinalizedEpoch() view returns(uint256)
func (_Bindings *BindingsSession) HighestFinalizedEpoch() (*big.Int, error) {
	return _Bindings.Contract.HighestFinalizedEpoch(&_Bindings.CallOpts)
}

// HighestFinalizedEpoch is a free data retrieval call binding the contract method 0xe8b0d84c.
//
// Solidity: function highestFinalizedEpoch() view returns(uint256)
func (_Bindings *BindingsCallerSession) HighestFinalizedEpoch() (*big.Int, error) {
	return _Bindings.Contract.HighestFinalizedEpoch(&_Bindings.CallOpts)
}

// HighestFinalizedTimestamp is a free data retrieval call binding the contract method 0x4d80bc74.
//
// Solidity: function highestFinalizedTimestamp() view returns(uint256)
func (_Bindings *BindingsCaller) HighestFinalizedTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "highestFinalizedTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestFinalizedTimestamp is a free data retrieval call binding the contract method 0x4d80bc74.
//
// Solidity: function highestFinalizedTimestamp() view returns(uint256)
func (_Bindings *BindingsSession) HighestFinalizedTimestamp() (*big.Int, error) {
	return _Bindings.Contract.HighestFinalizedTimestamp(&_Bindings.CallOpts)
}

// HighestFinalizedTimestamp is a free data retrieval call binding the contract method 0x4d80bc74.
//
// Solidity: function highestFinalizedTimestamp() view returns(uint256)
func (_Bindings *BindingsCallerSession) HighestFinalizedTimestamp() (*big.Int, error) {
	return _Bindings.Contract.HighestFinalizedTimestamp(&_Bindings.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0x950a4102.
//
// Solidity: function verifyProof(bytes32 root, (uint256,bytes32[],bytes32) proof) pure returns()
func (_Bindings *BindingsCaller) VerifyProof(opts *bind.CallOpts, root [32]byte, proof IMerkleProofVerifierProofInputs) error {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "verifyProof", root, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x950a4102.
//
// Solidity: function verifyProof(bytes32 root, (uint256,bytes32[],bytes32) proof) pure returns()
func (_Bindings *BindingsSession) VerifyProof(root [32]byte, proof IMerkleProofVerifierProofInputs) error {
	return _Bindings.Contract.VerifyProof(&_Bindings.CallOpts, root, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x950a4102.
//
// Solidity: function verifyProof(bytes32 root, (uint256,bytes32[],bytes32) proof) pure returns()
func (_Bindings *BindingsCallerSession) VerifyProof(root [32]byte, proof IMerkleProofVerifierProofInputs) error {
	return _Bindings.Contract.VerifyProof(&_Bindings.CallOpts, root, proof)
}

// ProveCheckpointFinalized is a paid mutator transaction binding the contract method 0x4ae5d418.
//
// Solidity: function proveCheckpointFinalized(uint256 timestamp, (uint256,bytes32[],bytes32) proof) returns()
func (_Bindings *BindingsTransactor) ProveCheckpointFinalized(opts *bind.TransactOpts, timestamp *big.Int, proof IMerkleProofVerifierProofInputs) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "proveCheckpointFinalized", timestamp, proof)
}

// ProveCheckpointFinalized is a paid mutator transaction binding the contract method 0x4ae5d418.
//
// Solidity: function proveCheckpointFinalized(uint256 timestamp, (uint256,bytes32[],bytes32) proof) returns()
func (_Bindings *BindingsSession) ProveCheckpointFinalized(timestamp *big.Int, proof IMerkleProofVerifierProofInputs) (*types.Transaction, error) {
	return _Bindings.Contract.ProveCheckpointFinalized(&_Bindings.TransactOpts, timestamp, proof)
}

// ProveCheckpointFinalized is a paid mutator transaction binding the contract method 0x4ae5d418.
//
// Solidity: function proveCheckpointFinalized(uint256 timestamp, (uint256,bytes32[],bytes32) proof) returns()
func (_Bindings *BindingsTransactorSession) ProveCheckpointFinalized(timestamp *big.Int, proof IMerkleProofVerifierProofInputs) (*types.Transaction, error) {
	return _Bindings.Contract.ProveCheckpointFinalized(&_Bindings.TransactOpts, timestamp, proof)
}

// BindingsFinalizedCheckpointUpdatedIterator is returned from FilterFinalizedCheckpointUpdated and is used to iterate over the raw logs and unpacked data for FinalizedCheckpointUpdated events raised by the Bindings contract.
type BindingsFinalizedCheckpointUpdatedIterator struct {
	Event *BindingsFinalizedCheckpointUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsFinalizedCheckpointUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsFinalizedCheckpointUpdated)
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
		it.Event = new(BindingsFinalizedCheckpointUpdated)
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
func (it *BindingsFinalizedCheckpointUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsFinalizedCheckpointUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsFinalizedCheckpointUpdated represents a FinalizedCheckpointUpdated event raised by the Bindings contract.
type BindingsFinalizedCheckpointUpdated struct {
	Epoch     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFinalizedCheckpointUpdated is a free log retrieval operation binding the contract event 0x4f39a67a5211ca86ad9972c17597c62bdb4fbae486921bb54160cc15cb285b85.
//
// Solidity: event FinalizedCheckpointUpdated(uint256 indexed epoch, uint256 indexed timestamp)
func (_Bindings *BindingsFilterer) FilterFinalizedCheckpointUpdated(opts *bind.FilterOpts, epoch []*big.Int, timestamp []*big.Int) (*BindingsFinalizedCheckpointUpdatedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "FinalizedCheckpointUpdated", epochRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &BindingsFinalizedCheckpointUpdatedIterator{contract: _Bindings.contract, event: "FinalizedCheckpointUpdated", logs: logs, sub: sub}, nil
}

// WatchFinalizedCheckpointUpdated is a free log subscription operation binding the contract event 0x4f39a67a5211ca86ad9972c17597c62bdb4fbae486921bb54160cc15cb285b85.
//
// Solidity: event FinalizedCheckpointUpdated(uint256 indexed epoch, uint256 indexed timestamp)
func (_Bindings *BindingsFilterer) WatchFinalizedCheckpointUpdated(opts *bind.WatchOpts, sink chan<- *BindingsFinalizedCheckpointUpdated, epoch []*big.Int, timestamp []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "FinalizedCheckpointUpdated", epochRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsFinalizedCheckpointUpdated)
				if err := _Bindings.contract.UnpackLog(event, "FinalizedCheckpointUpdated", log); err != nil {
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

// ParseFinalizedCheckpointUpdated is a log parse operation binding the contract event 0x4f39a67a5211ca86ad9972c17597c62bdb4fbae486921bb54160cc15cb285b85.
//
// Solidity: event FinalizedCheckpointUpdated(uint256 indexed epoch, uint256 indexed timestamp)
func (_Bindings *BindingsFilterer) ParseFinalizedCheckpointUpdated(log types.Log) (*BindingsFinalizedCheckpointUpdated, error) {
	event := new(BindingsFinalizedCheckpointUpdated)
	if err := _Bindings.contract.UnpackLog(event, "FinalizedCheckpointUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
