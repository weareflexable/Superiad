// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flexablenft

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

// FlexablenftMetaData contains all meta data concerning the Flexablenft contract.
var FlexablenftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint96\",\"name\":\"percentageBasisPoint\",\"type\":\"uint96\"}],\"name\":\"RoyaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"}],\"name\":\"TicketBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metaDataURI\",\"type\":\"string\"}],\"name\":\"TicketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"count\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"TicketRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLEXABLENFT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLEXABLENFT_CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLEXABLENFT_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TicketStatus\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"redeemCount\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"redeemInfo\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burnTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"createTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint96\",\"name\":\"royaltyPercentBasisPoint\",\"type\":\"uint96\"}],\"name\":\"createTicketWithCustomRoyalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"delegateTicketCreation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"royaltyaddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyPercentBasisPoint\",\"type\":\"uint96\"}],\"name\":\"delegateTicketCreationWithCustomRoyalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"redeemTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"royaltyReciever\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"percentageBasisPoint\",\"type\":\"uint96\"}],\"name\":\"updateDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlexablenftABI is the input ABI used to generate the binding from.
// Deprecated: Use FlexablenftMetaData.ABI instead.
var FlexablenftABI = FlexablenftMetaData.ABI

// Flexablenft is an auto generated Go binding around an Ethereum contract.
type Flexablenft struct {
	FlexablenftCaller     // Read-only binding to the contract
	FlexablenftTransactor // Write-only binding to the contract
	FlexablenftFilterer   // Log filterer for contract events
}

// FlexablenftCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlexablenftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlexablenftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlexablenftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlexablenftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlexablenftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlexablenftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlexablenftSession struct {
	Contract     *Flexablenft      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlexablenftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlexablenftCallerSession struct {
	Contract *FlexablenftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FlexablenftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlexablenftTransactorSession struct {
	Contract     *FlexablenftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FlexablenftRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlexablenftRaw struct {
	Contract *Flexablenft // Generic contract binding to access the raw methods on
}

// FlexablenftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlexablenftCallerRaw struct {
	Contract *FlexablenftCaller // Generic read-only contract binding to access the raw methods on
}

// FlexablenftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlexablenftTransactorRaw struct {
	Contract *FlexablenftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlexablenft creates a new instance of Flexablenft, bound to a specific deployed contract.
func NewFlexablenft(address common.Address, backend bind.ContractBackend) (*Flexablenft, error) {
	contract, err := bindFlexablenft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flexablenft{FlexablenftCaller: FlexablenftCaller{contract: contract}, FlexablenftTransactor: FlexablenftTransactor{contract: contract}, FlexablenftFilterer: FlexablenftFilterer{contract: contract}}, nil
}

// NewFlexablenftCaller creates a new read-only instance of Flexablenft, bound to a specific deployed contract.
func NewFlexablenftCaller(address common.Address, caller bind.ContractCaller) (*FlexablenftCaller, error) {
	contract, err := bindFlexablenft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlexablenftCaller{contract: contract}, nil
}

// NewFlexablenftTransactor creates a new write-only instance of Flexablenft, bound to a specific deployed contract.
func NewFlexablenftTransactor(address common.Address, transactor bind.ContractTransactor) (*FlexablenftTransactor, error) {
	contract, err := bindFlexablenft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlexablenftTransactor{contract: contract}, nil
}

// NewFlexablenftFilterer creates a new log filterer instance of Flexablenft, bound to a specific deployed contract.
func NewFlexablenftFilterer(address common.Address, filterer bind.ContractFilterer) (*FlexablenftFilterer, error) {
	contract, err := bindFlexablenft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlexablenftFilterer{contract: contract}, nil
}

// bindFlexablenft binds a generic wrapper to an already deployed contract.
func bindFlexablenft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlexablenftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flexablenft *FlexablenftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flexablenft.Contract.FlexablenftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flexablenft *FlexablenftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flexablenft.Contract.FlexablenftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flexablenft *FlexablenftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flexablenft.Contract.FlexablenftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flexablenft *FlexablenftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flexablenft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flexablenft *FlexablenftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flexablenft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flexablenft *FlexablenftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flexablenft.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Flexablenft.Contract.DEFAULTADMINROLE(&_Flexablenft.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Flexablenft.Contract.DEFAULTADMINROLE(&_Flexablenft.CallOpts)
}

// FLEXABLENFTADMINROLE is a free data retrieval call binding the contract method 0xea7676c7.
//
// Solidity: function FLEXABLENFT_ADMIN_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCaller) FLEXABLENFTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "FLEXABLENFT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FLEXABLENFTADMINROLE is a free data retrieval call binding the contract method 0xea7676c7.
//
// Solidity: function FLEXABLENFT_ADMIN_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftSession) FLEXABLENFTADMINROLE() ([32]byte, error) {
	return _Flexablenft.Contract.FLEXABLENFTADMINROLE(&_Flexablenft.CallOpts)
}

// FLEXABLENFTADMINROLE is a free data retrieval call binding the contract method 0xea7676c7.
//
// Solidity: function FLEXABLENFT_ADMIN_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCallerSession) FLEXABLENFTADMINROLE() ([32]byte, error) {
	return _Flexablenft.Contract.FLEXABLENFTADMINROLE(&_Flexablenft.CallOpts)
}

// FLEXABLENFTCREATORROLE is a free data retrieval call binding the contract method 0x0aca6616.
//
// Solidity: function FLEXABLENFT_CREATOR_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCaller) FLEXABLENFTCREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "FLEXABLENFT_CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FLEXABLENFTCREATORROLE is a free data retrieval call binding the contract method 0x0aca6616.
//
// Solidity: function FLEXABLENFT_CREATOR_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftSession) FLEXABLENFTCREATORROLE() ([32]byte, error) {
	return _Flexablenft.Contract.FLEXABLENFTCREATORROLE(&_Flexablenft.CallOpts)
}

// FLEXABLENFTCREATORROLE is a free data retrieval call binding the contract method 0x0aca6616.
//
// Solidity: function FLEXABLENFT_CREATOR_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCallerSession) FLEXABLENFTCREATORROLE() ([32]byte, error) {
	return _Flexablenft.Contract.FLEXABLENFTCREATORROLE(&_Flexablenft.CallOpts)
}

// FLEXABLENFTOPERATORROLE is a free data retrieval call binding the contract method 0x36a913b5.
//
// Solidity: function FLEXABLENFT_OPERATOR_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCaller) FLEXABLENFTOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "FLEXABLENFT_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FLEXABLENFTOPERATORROLE is a free data retrieval call binding the contract method 0x36a913b5.
//
// Solidity: function FLEXABLENFT_OPERATOR_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftSession) FLEXABLENFTOPERATORROLE() ([32]byte, error) {
	return _Flexablenft.Contract.FLEXABLENFTOPERATORROLE(&_Flexablenft.CallOpts)
}

// FLEXABLENFTOPERATORROLE is a free data retrieval call binding the contract method 0x36a913b5.
//
// Solidity: function FLEXABLENFT_OPERATOR_ROLE() view returns(bytes32)
func (_Flexablenft *FlexablenftCallerSession) FLEXABLENFTOPERATORROLE() ([32]byte, error) {
	return _Flexablenft.Contract.FLEXABLENFTOPERATORROLE(&_Flexablenft.CallOpts)
}

// TicketStatus is a free data retrieval call binding the contract method 0xce5956cb.
//
// Solidity: function TicketStatus(uint256 ) view returns(uint16 redeemCount, string redeemInfo)
func (_Flexablenft *FlexablenftCaller) TicketStatus(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RedeemCount uint16
	RedeemInfo  string
}, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "TicketStatus", arg0)

	outstruct := new(struct {
		RedeemCount uint16
		RedeemInfo  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RedeemCount = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.RedeemInfo = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// TicketStatus is a free data retrieval call binding the contract method 0xce5956cb.
//
// Solidity: function TicketStatus(uint256 ) view returns(uint16 redeemCount, string redeemInfo)
func (_Flexablenft *FlexablenftSession) TicketStatus(arg0 *big.Int) (struct {
	RedeemCount uint16
	RedeemInfo  string
}, error) {
	return _Flexablenft.Contract.TicketStatus(&_Flexablenft.CallOpts, arg0)
}

// TicketStatus is a free data retrieval call binding the contract method 0xce5956cb.
//
// Solidity: function TicketStatus(uint256 ) view returns(uint16 redeemCount, string redeemInfo)
func (_Flexablenft *FlexablenftCallerSession) TicketStatus(arg0 *big.Int) (struct {
	RedeemCount uint16
	RedeemInfo  string
}, error) {
	return _Flexablenft.Contract.TicketStatus(&_Flexablenft.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Flexablenft *FlexablenftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Flexablenft *FlexablenftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Flexablenft.Contract.BalanceOf(&_Flexablenft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Flexablenft *FlexablenftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Flexablenft.Contract.BalanceOf(&_Flexablenft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Flexablenft *FlexablenftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Flexablenft *FlexablenftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Flexablenft.Contract.GetApproved(&_Flexablenft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Flexablenft *FlexablenftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Flexablenft.Contract.GetApproved(&_Flexablenft.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Flexablenft *FlexablenftCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Flexablenft *FlexablenftSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Flexablenft.Contract.GetRoleAdmin(&_Flexablenft.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Flexablenft *FlexablenftCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Flexablenft.Contract.GetRoleAdmin(&_Flexablenft.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Flexablenft *FlexablenftCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Flexablenft *FlexablenftSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Flexablenft.Contract.GetRoleMember(&_Flexablenft.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Flexablenft *FlexablenftCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Flexablenft.Contract.GetRoleMember(&_Flexablenft.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Flexablenft *FlexablenftCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Flexablenft *FlexablenftSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Flexablenft.Contract.GetRoleMemberCount(&_Flexablenft.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Flexablenft *FlexablenftCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Flexablenft.Contract.GetRoleMemberCount(&_Flexablenft.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Flexablenft *FlexablenftCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Flexablenft *FlexablenftSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Flexablenft.Contract.HasRole(&_Flexablenft.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Flexablenft *FlexablenftCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Flexablenft.Contract.HasRole(&_Flexablenft.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Flexablenft *FlexablenftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Flexablenft *FlexablenftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Flexablenft.Contract.IsApprovedForAll(&_Flexablenft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Flexablenft *FlexablenftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Flexablenft.Contract.IsApprovedForAll(&_Flexablenft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flexablenft *FlexablenftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flexablenft *FlexablenftSession) Name() (string, error) {
	return _Flexablenft.Contract.Name(&_Flexablenft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Flexablenft *FlexablenftCallerSession) Name() (string, error) {
	return _Flexablenft.Contract.Name(&_Flexablenft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Flexablenft *FlexablenftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Flexablenft *FlexablenftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Flexablenft.Contract.OwnerOf(&_Flexablenft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Flexablenft *FlexablenftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Flexablenft.Contract.OwnerOf(&_Flexablenft.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Flexablenft *FlexablenftCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Flexablenft *FlexablenftSession) Paused() (bool, error) {
	return _Flexablenft.Contract.Paused(&_Flexablenft.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Flexablenft *FlexablenftCallerSession) Paused() (bool, error) {
	return _Flexablenft.Contract.Paused(&_Flexablenft.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Flexablenft *FlexablenftCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Flexablenft *FlexablenftSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Flexablenft.Contract.RoyaltyInfo(&_Flexablenft.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Flexablenft *FlexablenftCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Flexablenft.Contract.RoyaltyInfo(&_Flexablenft.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Flexablenft *FlexablenftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Flexablenft *FlexablenftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Flexablenft.Contract.SupportsInterface(&_Flexablenft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Flexablenft *FlexablenftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Flexablenft.Contract.SupportsInterface(&_Flexablenft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flexablenft *FlexablenftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flexablenft *FlexablenftSession) Symbol() (string, error) {
	return _Flexablenft.Contract.Symbol(&_Flexablenft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Flexablenft *FlexablenftCallerSession) Symbol() (string, error) {
	return _Flexablenft.Contract.Symbol(&_Flexablenft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Flexablenft *FlexablenftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Flexablenft *FlexablenftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Flexablenft.Contract.TokenByIndex(&_Flexablenft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Flexablenft *FlexablenftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Flexablenft.Contract.TokenByIndex(&_Flexablenft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Flexablenft *FlexablenftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Flexablenft *FlexablenftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Flexablenft.Contract.TokenOfOwnerByIndex(&_Flexablenft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Flexablenft *FlexablenftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Flexablenft.Contract.TokenOfOwnerByIndex(&_Flexablenft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Flexablenft *FlexablenftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Flexablenft *FlexablenftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Flexablenft.Contract.TokenURI(&_Flexablenft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Flexablenft *FlexablenftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Flexablenft.Contract.TokenURI(&_Flexablenft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Flexablenft *FlexablenftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flexablenft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Flexablenft *FlexablenftSession) TotalSupply() (*big.Int, error) {
	return _Flexablenft.Contract.TotalSupply(&_Flexablenft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Flexablenft *FlexablenftCallerSession) TotalSupply() (*big.Int, error) {
	return _Flexablenft.Contract.TotalSupply(&_Flexablenft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.Approve(&_Flexablenft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.Approve(&_Flexablenft.TransactOpts, to, tokenId)
}

// BurnTicket is a paid mutator transaction binding the contract method 0xa476b73d.
//
// Solidity: function burnTicket(uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactor) BurnTicket(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "burnTicket", tokenId)
}

// BurnTicket is a paid mutator transaction binding the contract method 0xa476b73d.
//
// Solidity: function burnTicket(uint256 tokenId) returns()
func (_Flexablenft *FlexablenftSession) BurnTicket(tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.BurnTicket(&_Flexablenft.TransactOpts, tokenId)
}

// BurnTicket is a paid mutator transaction binding the contract method 0xa476b73d.
//
// Solidity: function burnTicket(uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactorSession) BurnTicket(tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.BurnTicket(&_Flexablenft.TransactOpts, tokenId)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x6897082f.
//
// Solidity: function createTicket(string metadataURI) returns(uint256)
func (_Flexablenft *FlexablenftTransactor) CreateTicket(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "createTicket", metadataURI)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x6897082f.
//
// Solidity: function createTicket(string metadataURI) returns(uint256)
func (_Flexablenft *FlexablenftSession) CreateTicket(metadataURI string) (*types.Transaction, error) {
	return _Flexablenft.Contract.CreateTicket(&_Flexablenft.TransactOpts, metadataURI)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x6897082f.
//
// Solidity: function createTicket(string metadataURI) returns(uint256)
func (_Flexablenft *FlexablenftTransactorSession) CreateTicket(metadataURI string) (*types.Transaction, error) {
	return _Flexablenft.Contract.CreateTicket(&_Flexablenft.TransactOpts, metadataURI)
}

// CreateTicketWithCustomRoyalty is a paid mutator transaction binding the contract method 0x9285a51b.
//
// Solidity: function createTicketWithCustomRoyalty(string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_Flexablenft *FlexablenftTransactor) CreateTicketWithCustomRoyalty(opts *bind.TransactOpts, metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "createTicketWithCustomRoyalty", metadataURI, royaltyPercentBasisPoint)
}

// CreateTicketWithCustomRoyalty is a paid mutator transaction binding the contract method 0x9285a51b.
//
// Solidity: function createTicketWithCustomRoyalty(string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_Flexablenft *FlexablenftSession) CreateTicketWithCustomRoyalty(metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.CreateTicketWithCustomRoyalty(&_Flexablenft.TransactOpts, metadataURI, royaltyPercentBasisPoint)
}

// CreateTicketWithCustomRoyalty is a paid mutator transaction binding the contract method 0x9285a51b.
//
// Solidity: function createTicketWithCustomRoyalty(string metadataURI, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_Flexablenft *FlexablenftTransactorSession) CreateTicketWithCustomRoyalty(metadataURI string, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.CreateTicketWithCustomRoyalty(&_Flexablenft.TransactOpts, metadataURI, royaltyPercentBasisPoint)
}

// DelegateTicketCreation is a paid mutator transaction binding the contract method 0x00c95078.
//
// Solidity: function delegateTicketCreation(address creator, string metadataURI) returns(uint256)
func (_Flexablenft *FlexablenftTransactor) DelegateTicketCreation(opts *bind.TransactOpts, creator common.Address, metadataURI string) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "delegateTicketCreation", creator, metadataURI)
}

// DelegateTicketCreation is a paid mutator transaction binding the contract method 0x00c95078.
//
// Solidity: function delegateTicketCreation(address creator, string metadataURI) returns(uint256)
func (_Flexablenft *FlexablenftSession) DelegateTicketCreation(creator common.Address, metadataURI string) (*types.Transaction, error) {
	return _Flexablenft.Contract.DelegateTicketCreation(&_Flexablenft.TransactOpts, creator, metadataURI)
}

// DelegateTicketCreation is a paid mutator transaction binding the contract method 0x00c95078.
//
// Solidity: function delegateTicketCreation(address creator, string metadataURI) returns(uint256)
func (_Flexablenft *FlexablenftTransactorSession) DelegateTicketCreation(creator common.Address, metadataURI string) (*types.Transaction, error) {
	return _Flexablenft.Contract.DelegateTicketCreation(&_Flexablenft.TransactOpts, creator, metadataURI)
}

// DelegateTicketCreationWithCustomRoyalty is a paid mutator transaction binding the contract method 0x2875970a.
//
// Solidity: function delegateTicketCreationWithCustomRoyalty(address creator, string metadataURI, address royaltyaddress, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_Flexablenft *FlexablenftTransactor) DelegateTicketCreationWithCustomRoyalty(opts *bind.TransactOpts, creator common.Address, metadataURI string, royaltyaddress common.Address, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "delegateTicketCreationWithCustomRoyalty", creator, metadataURI, royaltyaddress, royaltyPercentBasisPoint)
}

// DelegateTicketCreationWithCustomRoyalty is a paid mutator transaction binding the contract method 0x2875970a.
//
// Solidity: function delegateTicketCreationWithCustomRoyalty(address creator, string metadataURI, address royaltyaddress, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_Flexablenft *FlexablenftSession) DelegateTicketCreationWithCustomRoyalty(creator common.Address, metadataURI string, royaltyaddress common.Address, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.DelegateTicketCreationWithCustomRoyalty(&_Flexablenft.TransactOpts, creator, metadataURI, royaltyaddress, royaltyPercentBasisPoint)
}

// DelegateTicketCreationWithCustomRoyalty is a paid mutator transaction binding the contract method 0x2875970a.
//
// Solidity: function delegateTicketCreationWithCustomRoyalty(address creator, string metadataURI, address royaltyaddress, uint96 royaltyPercentBasisPoint) returns(uint256)
func (_Flexablenft *FlexablenftTransactorSession) DelegateTicketCreationWithCustomRoyalty(creator common.Address, metadataURI string, royaltyaddress common.Address, royaltyPercentBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.DelegateTicketCreationWithCustomRoyalty(&_Flexablenft.TransactOpts, creator, metadataURI, royaltyaddress, royaltyPercentBasisPoint)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.Contract.GrantRole(&_Flexablenft.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.Contract.GrantRole(&_Flexablenft.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Flexablenft *FlexablenftTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Flexablenft *FlexablenftSession) Pause() (*types.Transaction, error) {
	return _Flexablenft.Contract.Pause(&_Flexablenft.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Flexablenft *FlexablenftTransactorSession) Pause() (*types.Transaction, error) {
	return _Flexablenft.Contract.Pause(&_Flexablenft.TransactOpts)
}

// RedeemTicket is a paid mutator transaction binding the contract method 0x51a2c4e3.
//
// Solidity: function redeemTicket(uint256 tokenId, string info) returns()
func (_Flexablenft *FlexablenftTransactor) RedeemTicket(opts *bind.TransactOpts, tokenId *big.Int, info string) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "redeemTicket", tokenId, info)
}

// RedeemTicket is a paid mutator transaction binding the contract method 0x51a2c4e3.
//
// Solidity: function redeemTicket(uint256 tokenId, string info) returns()
func (_Flexablenft *FlexablenftSession) RedeemTicket(tokenId *big.Int, info string) (*types.Transaction, error) {
	return _Flexablenft.Contract.RedeemTicket(&_Flexablenft.TransactOpts, tokenId, info)
}

// RedeemTicket is a paid mutator transaction binding the contract method 0x51a2c4e3.
//
// Solidity: function redeemTicket(uint256 tokenId, string info) returns()
func (_Flexablenft *FlexablenftTransactorSession) RedeemTicket(tokenId *big.Int, info string) (*types.Transaction, error) {
	return _Flexablenft.Contract.RedeemTicket(&_Flexablenft.TransactOpts, tokenId, info)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.Contract.RenounceRole(&_Flexablenft.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.Contract.RenounceRole(&_Flexablenft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.Contract.RevokeRole(&_Flexablenft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Flexablenft *FlexablenftTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Flexablenft.Contract.RevokeRole(&_Flexablenft.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.SafeTransferFrom(&_Flexablenft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.SafeTransferFrom(&_Flexablenft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Flexablenft *FlexablenftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Flexablenft *FlexablenftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Flexablenft.Contract.SafeTransferFrom0(&_Flexablenft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Flexablenft *FlexablenftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Flexablenft.Contract.SafeTransferFrom0(&_Flexablenft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Flexablenft *FlexablenftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Flexablenft *FlexablenftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Flexablenft.Contract.SetApprovalForAll(&_Flexablenft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Flexablenft *FlexablenftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Flexablenft.Contract.SetApprovalForAll(&_Flexablenft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.TransferFrom(&_Flexablenft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Flexablenft *FlexablenftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.TransferFrom(&_Flexablenft.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Flexablenft *FlexablenftTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Flexablenft *FlexablenftSession) Unpause() (*types.Transaction, error) {
	return _Flexablenft.Contract.Unpause(&_Flexablenft.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Flexablenft *FlexablenftTransactorSession) Unpause() (*types.Transaction, error) {
	return _Flexablenft.Contract.Unpause(&_Flexablenft.TransactOpts)
}

// UpdateDefaultRoyalty is a paid mutator transaction binding the contract method 0xd691e43c.
//
// Solidity: function updateDefaultRoyalty(address royaltyReciever, uint96 percentageBasisPoint) returns()
func (_Flexablenft *FlexablenftTransactor) UpdateDefaultRoyalty(opts *bind.TransactOpts, royaltyReciever common.Address, percentageBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.contract.Transact(opts, "updateDefaultRoyalty", royaltyReciever, percentageBasisPoint)
}

// UpdateDefaultRoyalty is a paid mutator transaction binding the contract method 0xd691e43c.
//
// Solidity: function updateDefaultRoyalty(address royaltyReciever, uint96 percentageBasisPoint) returns()
func (_Flexablenft *FlexablenftSession) UpdateDefaultRoyalty(royaltyReciever common.Address, percentageBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.UpdateDefaultRoyalty(&_Flexablenft.TransactOpts, royaltyReciever, percentageBasisPoint)
}

// UpdateDefaultRoyalty is a paid mutator transaction binding the contract method 0xd691e43c.
//
// Solidity: function updateDefaultRoyalty(address royaltyReciever, uint96 percentageBasisPoint) returns()
func (_Flexablenft *FlexablenftTransactorSession) UpdateDefaultRoyalty(royaltyReciever common.Address, percentageBasisPoint *big.Int) (*types.Transaction, error) {
	return _Flexablenft.Contract.UpdateDefaultRoyalty(&_Flexablenft.TransactOpts, royaltyReciever, percentageBasisPoint)
}

// FlexablenftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Flexablenft contract.
type FlexablenftApprovalIterator struct {
	Event *FlexablenftApproval // Event containing the contract specifics and raw log

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
func (it *FlexablenftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftApproval)
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
		it.Event = new(FlexablenftApproval)
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
func (it *FlexablenftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftApproval represents a Approval event raised by the Flexablenft contract.
type FlexablenftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Flexablenft *FlexablenftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FlexablenftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftApprovalIterator{contract: _Flexablenft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Flexablenft *FlexablenftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FlexablenftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftApproval)
				if err := _Flexablenft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Flexablenft *FlexablenftFilterer) ParseApproval(log types.Log) (*FlexablenftApproval, error) {
	event := new(FlexablenftApproval)
	if err := _Flexablenft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Flexablenft contract.
type FlexablenftApprovalForAllIterator struct {
	Event *FlexablenftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FlexablenftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftApprovalForAll)
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
		it.Event = new(FlexablenftApprovalForAll)
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
func (it *FlexablenftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftApprovalForAll represents a ApprovalForAll event raised by the Flexablenft contract.
type FlexablenftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Flexablenft *FlexablenftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FlexablenftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftApprovalForAllIterator{contract: _Flexablenft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Flexablenft *FlexablenftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FlexablenftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftApprovalForAll)
				if err := _Flexablenft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Flexablenft *FlexablenftFilterer) ParseApprovalForAll(log types.Log) (*FlexablenftApprovalForAll, error) {
	event := new(FlexablenftApprovalForAll)
	if err := _Flexablenft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Flexablenft contract.
type FlexablenftPausedIterator struct {
	Event *FlexablenftPaused // Event containing the contract specifics and raw log

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
func (it *FlexablenftPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftPaused)
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
		it.Event = new(FlexablenftPaused)
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
func (it *FlexablenftPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftPaused represents a Paused event raised by the Flexablenft contract.
type FlexablenftPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Flexablenft *FlexablenftFilterer) FilterPaused(opts *bind.FilterOpts) (*FlexablenftPausedIterator, error) {

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FlexablenftPausedIterator{contract: _Flexablenft.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Flexablenft *FlexablenftFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FlexablenftPaused) (event.Subscription, error) {

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftPaused)
				if err := _Flexablenft.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Flexablenft *FlexablenftFilterer) ParsePaused(log types.Log) (*FlexablenftPaused, error) {
	event := new(FlexablenftPaused)
	if err := _Flexablenft.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Flexablenft contract.
type FlexablenftRoleAdminChangedIterator struct {
	Event *FlexablenftRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FlexablenftRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftRoleAdminChanged)
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
		it.Event = new(FlexablenftRoleAdminChanged)
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
func (it *FlexablenftRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftRoleAdminChanged represents a RoleAdminChanged event raised by the Flexablenft contract.
type FlexablenftRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Flexablenft *FlexablenftFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FlexablenftRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftRoleAdminChangedIterator{contract: _Flexablenft.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Flexablenft *FlexablenftFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FlexablenftRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftRoleAdminChanged)
				if err := _Flexablenft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Flexablenft *FlexablenftFilterer) ParseRoleAdminChanged(log types.Log) (*FlexablenftRoleAdminChanged, error) {
	event := new(FlexablenftRoleAdminChanged)
	if err := _Flexablenft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Flexablenft contract.
type FlexablenftRoleGrantedIterator struct {
	Event *FlexablenftRoleGranted // Event containing the contract specifics and raw log

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
func (it *FlexablenftRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftRoleGranted)
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
		it.Event = new(FlexablenftRoleGranted)
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
func (it *FlexablenftRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftRoleGranted represents a RoleGranted event raised by the Flexablenft contract.
type FlexablenftRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Flexablenft *FlexablenftFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FlexablenftRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftRoleGrantedIterator{contract: _Flexablenft.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Flexablenft *FlexablenftFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FlexablenftRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftRoleGranted)
				if err := _Flexablenft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Flexablenft *FlexablenftFilterer) ParseRoleGranted(log types.Log) (*FlexablenftRoleGranted, error) {
	event := new(FlexablenftRoleGranted)
	if err := _Flexablenft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Flexablenft contract.
type FlexablenftRoleRevokedIterator struct {
	Event *FlexablenftRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FlexablenftRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftRoleRevoked)
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
		it.Event = new(FlexablenftRoleRevoked)
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
func (it *FlexablenftRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftRoleRevoked represents a RoleRevoked event raised by the Flexablenft contract.
type FlexablenftRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Flexablenft *FlexablenftFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FlexablenftRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftRoleRevokedIterator{contract: _Flexablenft.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Flexablenft *FlexablenftFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FlexablenftRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftRoleRevoked)
				if err := _Flexablenft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Flexablenft *FlexablenftFilterer) ParseRoleRevoked(log types.Log) (*FlexablenftRoleRevoked, error) {
	event := new(FlexablenftRoleRevoked)
	if err := _Flexablenft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftRoyaltyUpdatedIterator is returned from FilterRoyaltyUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyUpdated events raised by the Flexablenft contract.
type FlexablenftRoyaltyUpdatedIterator struct {
	Event *FlexablenftRoyaltyUpdated // Event containing the contract specifics and raw log

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
func (it *FlexablenftRoyaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftRoyaltyUpdated)
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
		it.Event = new(FlexablenftRoyaltyUpdated)
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
func (it *FlexablenftRoyaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftRoyaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftRoyaltyUpdated represents a RoyaltyUpdated event raised by the Flexablenft contract.
type FlexablenftRoyaltyUpdated struct {
	Reciever             common.Address
	PercentageBasisPoint *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyUpdated is a free log retrieval operation binding the contract event 0x8039bd6e4e7dba001c8840eb2e118d9d131246faa7d0d04335f7305127ec0b10.
//
// Solidity: event RoyaltyUpdated(address indexed reciever, uint96 indexed percentageBasisPoint)
func (_Flexablenft *FlexablenftFilterer) FilterRoyaltyUpdated(opts *bind.FilterOpts, reciever []common.Address, percentageBasisPoint []*big.Int) (*FlexablenftRoyaltyUpdatedIterator, error) {

	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}
	var percentageBasisPointRule []interface{}
	for _, percentageBasisPointItem := range percentageBasisPoint {
		percentageBasisPointRule = append(percentageBasisPointRule, percentageBasisPointItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "RoyaltyUpdated", recieverRule, percentageBasisPointRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftRoyaltyUpdatedIterator{contract: _Flexablenft.contract, event: "RoyaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyUpdated is a free log subscription operation binding the contract event 0x8039bd6e4e7dba001c8840eb2e118d9d131246faa7d0d04335f7305127ec0b10.
//
// Solidity: event RoyaltyUpdated(address indexed reciever, uint96 indexed percentageBasisPoint)
func (_Flexablenft *FlexablenftFilterer) WatchRoyaltyUpdated(opts *bind.WatchOpts, sink chan<- *FlexablenftRoyaltyUpdated, reciever []common.Address, percentageBasisPoint []*big.Int) (event.Subscription, error) {

	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}
	var percentageBasisPointRule []interface{}
	for _, percentageBasisPointItem := range percentageBasisPoint {
		percentageBasisPointRule = append(percentageBasisPointRule, percentageBasisPointItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "RoyaltyUpdated", recieverRule, percentageBasisPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftRoyaltyUpdated)
				if err := _Flexablenft.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
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

// ParseRoyaltyUpdated is a log parse operation binding the contract event 0x8039bd6e4e7dba001c8840eb2e118d9d131246faa7d0d04335f7305127ec0b10.
//
// Solidity: event RoyaltyUpdated(address indexed reciever, uint96 indexed percentageBasisPoint)
func (_Flexablenft *FlexablenftFilterer) ParseRoyaltyUpdated(log types.Log) (*FlexablenftRoyaltyUpdated, error) {
	event := new(FlexablenftRoyaltyUpdated)
	if err := _Flexablenft.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftTicketBurntIterator is returned from FilterTicketBurnt and is used to iterate over the raw logs and unpacked data for TicketBurnt events raised by the Flexablenft contract.
type FlexablenftTicketBurntIterator struct {
	Event *FlexablenftTicketBurnt // Event containing the contract specifics and raw log

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
func (it *FlexablenftTicketBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftTicketBurnt)
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
		it.Event = new(FlexablenftTicketBurnt)
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
func (it *FlexablenftTicketBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftTicketBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftTicketBurnt represents a TicketBurnt event raised by the Flexablenft contract.
type FlexablenftTicketBurnt struct {
	TokenId         *big.Int
	OwnerOrApproved common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTicketBurnt is a free log retrieval operation binding the contract event 0x992fdab418cdc52e18511f72c6e1a38083d4e6251471f459051d80001131fe73.
//
// Solidity: event TicketBurnt(uint256 indexed tokenId, address indexed ownerOrApproved)
func (_Flexablenft *FlexablenftFilterer) FilterTicketBurnt(opts *bind.FilterOpts, tokenId []*big.Int, ownerOrApproved []common.Address) (*FlexablenftTicketBurntIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "TicketBurnt", tokenIdRule, ownerOrApprovedRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftTicketBurntIterator{contract: _Flexablenft.contract, event: "TicketBurnt", logs: logs, sub: sub}, nil
}

// WatchTicketBurnt is a free log subscription operation binding the contract event 0x992fdab418cdc52e18511f72c6e1a38083d4e6251471f459051d80001131fe73.
//
// Solidity: event TicketBurnt(uint256 indexed tokenId, address indexed ownerOrApproved)
func (_Flexablenft *FlexablenftFilterer) WatchTicketBurnt(opts *bind.WatchOpts, sink chan<- *FlexablenftTicketBurnt, tokenId []*big.Int, ownerOrApproved []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "TicketBurnt", tokenIdRule, ownerOrApprovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftTicketBurnt)
				if err := _Flexablenft.contract.UnpackLog(event, "TicketBurnt", log); err != nil {
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

// ParseTicketBurnt is a log parse operation binding the contract event 0x992fdab418cdc52e18511f72c6e1a38083d4e6251471f459051d80001131fe73.
//
// Solidity: event TicketBurnt(uint256 indexed tokenId, address indexed ownerOrApproved)
func (_Flexablenft *FlexablenftFilterer) ParseTicketBurnt(log types.Log) (*FlexablenftTicketBurnt, error) {
	event := new(FlexablenftTicketBurnt)
	if err := _Flexablenft.contract.UnpackLog(event, "TicketBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftTicketCreatedIterator is returned from FilterTicketCreated and is used to iterate over the raw logs and unpacked data for TicketCreated events raised by the Flexablenft contract.
type FlexablenftTicketCreatedIterator struct {
	Event *FlexablenftTicketCreated // Event containing the contract specifics and raw log

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
func (it *FlexablenftTicketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftTicketCreated)
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
		it.Event = new(FlexablenftTicketCreated)
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
func (it *FlexablenftTicketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftTicketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftTicketCreated represents a TicketCreated event raised by the Flexablenft contract.
type FlexablenftTicketCreated struct {
	TokenID     *big.Int
	Creator     common.Address
	MetaDataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTicketCreated is a free log retrieval operation binding the contract event 0x0484f94842646b419a0e924ba515b27093fce72662146119168089c5279dd3e4.
//
// Solidity: event TicketCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_Flexablenft *FlexablenftFilterer) FilterTicketCreated(opts *bind.FilterOpts, creator []common.Address) (*FlexablenftTicketCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "TicketCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftTicketCreatedIterator{contract: _Flexablenft.contract, event: "TicketCreated", logs: logs, sub: sub}, nil
}

// WatchTicketCreated is a free log subscription operation binding the contract event 0x0484f94842646b419a0e924ba515b27093fce72662146119168089c5279dd3e4.
//
// Solidity: event TicketCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_Flexablenft *FlexablenftFilterer) WatchTicketCreated(opts *bind.WatchOpts, sink chan<- *FlexablenftTicketCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "TicketCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftTicketCreated)
				if err := _Flexablenft.contract.UnpackLog(event, "TicketCreated", log); err != nil {
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

// ParseTicketCreated is a log parse operation binding the contract event 0x0484f94842646b419a0e924ba515b27093fce72662146119168089c5279dd3e4.
//
// Solidity: event TicketCreated(uint256 tokenID, address indexed creator, string metaDataURI)
func (_Flexablenft *FlexablenftFilterer) ParseTicketCreated(log types.Log) (*FlexablenftTicketCreated, error) {
	event := new(FlexablenftTicketCreated)
	if err := _Flexablenft.contract.UnpackLog(event, "TicketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftTicketRedeemedIterator is returned from FilterTicketRedeemed and is used to iterate over the raw logs and unpacked data for TicketRedeemed events raised by the Flexablenft contract.
type FlexablenftTicketRedeemedIterator struct {
	Event *FlexablenftTicketRedeemed // Event containing the contract specifics and raw log

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
func (it *FlexablenftTicketRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftTicketRedeemed)
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
		it.Event = new(FlexablenftTicketRedeemed)
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
func (it *FlexablenftTicketRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftTicketRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftTicketRedeemed represents a TicketRedeemed event raised by the Flexablenft contract.
type FlexablenftTicketRedeemed struct {
	TokenID *big.Int
	Count   uint16
	Info    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTicketRedeemed is a free log retrieval operation binding the contract event 0x91e9be99c353eb314b6fc412ede9f15d7187a1b918bf65a53df2a1de9eb5c8f0.
//
// Solidity: event TicketRedeemed(uint256 indexed tokenID, uint16 indexed count, string info)
func (_Flexablenft *FlexablenftFilterer) FilterTicketRedeemed(opts *bind.FilterOpts, tokenID []*big.Int, count []uint16) (*FlexablenftTicketRedeemedIterator, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "TicketRedeemed", tokenIDRule, countRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftTicketRedeemedIterator{contract: _Flexablenft.contract, event: "TicketRedeemed", logs: logs, sub: sub}, nil
}

// WatchTicketRedeemed is a free log subscription operation binding the contract event 0x91e9be99c353eb314b6fc412ede9f15d7187a1b918bf65a53df2a1de9eb5c8f0.
//
// Solidity: event TicketRedeemed(uint256 indexed tokenID, uint16 indexed count, string info)
func (_Flexablenft *FlexablenftFilterer) WatchTicketRedeemed(opts *bind.WatchOpts, sink chan<- *FlexablenftTicketRedeemed, tokenID []*big.Int, count []uint16) (event.Subscription, error) {

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "TicketRedeemed", tokenIDRule, countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftTicketRedeemed)
				if err := _Flexablenft.contract.UnpackLog(event, "TicketRedeemed", log); err != nil {
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

// ParseTicketRedeemed is a log parse operation binding the contract event 0x91e9be99c353eb314b6fc412ede9f15d7187a1b918bf65a53df2a1de9eb5c8f0.
//
// Solidity: event TicketRedeemed(uint256 indexed tokenID, uint16 indexed count, string info)
func (_Flexablenft *FlexablenftFilterer) ParseTicketRedeemed(log types.Log) (*FlexablenftTicketRedeemed, error) {
	event := new(FlexablenftTicketRedeemed)
	if err := _Flexablenft.contract.UnpackLog(event, "TicketRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Flexablenft contract.
type FlexablenftTransferIterator struct {
	Event *FlexablenftTransfer // Event containing the contract specifics and raw log

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
func (it *FlexablenftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftTransfer)
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
		it.Event = new(FlexablenftTransfer)
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
func (it *FlexablenftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftTransfer represents a Transfer event raised by the Flexablenft contract.
type FlexablenftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Flexablenft *FlexablenftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FlexablenftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FlexablenftTransferIterator{contract: _Flexablenft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Flexablenft *FlexablenftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FlexablenftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftTransfer)
				if err := _Flexablenft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Flexablenft *FlexablenftFilterer) ParseTransfer(log types.Log) (*FlexablenftTransfer, error) {
	event := new(FlexablenftTransfer)
	if err := _Flexablenft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlexablenftUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Flexablenft contract.
type FlexablenftUnpausedIterator struct {
	Event *FlexablenftUnpaused // Event containing the contract specifics and raw log

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
func (it *FlexablenftUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlexablenftUnpaused)
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
		it.Event = new(FlexablenftUnpaused)
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
func (it *FlexablenftUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlexablenftUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlexablenftUnpaused represents a Unpaused event raised by the Flexablenft contract.
type FlexablenftUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Flexablenft *FlexablenftFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FlexablenftUnpausedIterator, error) {

	logs, sub, err := _Flexablenft.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FlexablenftUnpausedIterator{contract: _Flexablenft.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Flexablenft *FlexablenftFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FlexablenftUnpaused) (event.Subscription, error) {

	logs, sub, err := _Flexablenft.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlexablenftUnpaused)
				if err := _Flexablenft.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Flexablenft *FlexablenftFilterer) ParseUnpaused(log types.Log) (*FlexablenftUnpaused, error) {
	event := new(FlexablenftUnpaused)
	if err := _Flexablenft.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
