// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain

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

// IFloodPlainItem is an auto generated low-level Go binding around an user-defined struct.
type IFloodPlainItem struct {
	Token  common.Address
	Amount *big.Int
}

// IFloodPlainOnChainOrdersOrderWithSignature is an auto generated low-level Go binding around an user-defined struct.
type IFloodPlainOnChainOrdersOrderWithSignature struct {
	Order     IFloodPlainOrder
	Signature []byte
}

// IFloodPlainOrder is an auto generated low-level Go binding around an user-defined struct.
type IFloodPlainOrder struct {
	Offerer       common.Address
	Zone          common.Address
	Offer         []IFloodPlainItem
	Consideration []IFloodPlainItem
	Deadline      *big.Int
	Nonce         *big.Int
}

// FloodPlainL2MetaData contains all meta data concerning the FloodPlainL2 contract.
var FloodPlainL2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectValueReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountPulled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyDecoders\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"decoderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"decoder\",\"type\":\"address\"}],\"name\":\"DecoderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"OrderEtched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT2\",\"outputs\":[{\"internalType\":\"contractISignatureTransfer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"decoder\",\"type\":\"address\"}],\"name\":\"addDecoder\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decoderId\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFloodPlainOnChainOrders.OrderWithSignature\",\"name\":\"orderWithSignature\",\"type\":\"tuple\"}],\"name\":\"etchOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"fulfillEtchedOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"fulfillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fulfillOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"decoderId\",\"type\":\"uint256\"}],\"name\":\"getDecoder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"etchedOrderId\",\"type\":\"uint256\"}],\"name\":\"getEtchedOrder\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIFloodPlainOnChainOrders.OrderWithSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getNonceStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"fulfiller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"getOrderValidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"getOrderValidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Item[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFloodPlain.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getPermitHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FloodPlainL2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FloodPlainL2MetaData.ABI instead.
var FloodPlainL2ABI = FloodPlainL2MetaData.ABI

// FloodPlainL2 is an auto generated Go binding around an Ethereum contract.
type FloodPlainL2 struct {
	FloodPlainL2Caller     // Read-only binding to the contract
	FloodPlainL2Transactor // Write-only binding to the contract
	FloodPlainL2Filterer   // Log filterer for contract events
}

// FloodPlainL2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FloodPlainL2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FloodPlainL2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FloodPlainL2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FloodPlainL2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FloodPlainL2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FloodPlainL2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FloodPlainL2Session struct {
	Contract     *FloodPlainL2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FloodPlainL2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FloodPlainL2CallerSession struct {
	Contract *FloodPlainL2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FloodPlainL2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FloodPlainL2TransactorSession struct {
	Contract     *FloodPlainL2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FloodPlainL2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FloodPlainL2Raw struct {
	Contract *FloodPlainL2 // Generic contract binding to access the raw methods on
}

// FloodPlainL2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FloodPlainL2CallerRaw struct {
	Contract *FloodPlainL2Caller // Generic read-only contract binding to access the raw methods on
}

// FloodPlainL2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FloodPlainL2TransactorRaw struct {
	Contract *FloodPlainL2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFloodPlainL2 creates a new instance of FloodPlainL2, bound to a specific deployed contract.
func NewFloodPlainL2(address common.Address, backend bind.ContractBackend) (*FloodPlainL2, error) {
	contract, err := bindFloodPlainL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2{FloodPlainL2Caller: FloodPlainL2Caller{contract: contract}, FloodPlainL2Transactor: FloodPlainL2Transactor{contract: contract}, FloodPlainL2Filterer: FloodPlainL2Filterer{contract: contract}}, nil
}

// NewFloodPlainL2Caller creates a new read-only instance of FloodPlainL2, bound to a specific deployed contract.
func NewFloodPlainL2Caller(address common.Address, caller bind.ContractCaller) (*FloodPlainL2Caller, error) {
	contract, err := bindFloodPlainL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2Caller{contract: contract}, nil
}

// NewFloodPlainL2Transactor creates a new write-only instance of FloodPlainL2, bound to a specific deployed contract.
func NewFloodPlainL2Transactor(address common.Address, transactor bind.ContractTransactor) (*FloodPlainL2Transactor, error) {
	contract, err := bindFloodPlainL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2Transactor{contract: contract}, nil
}

// NewFloodPlainL2Filterer creates a new log filterer instance of FloodPlainL2, bound to a specific deployed contract.
func NewFloodPlainL2Filterer(address common.Address, filterer bind.ContractFilterer) (*FloodPlainL2Filterer, error) {
	contract, err := bindFloodPlainL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2Filterer{contract: contract}, nil
}

// bindFloodPlainL2 binds a generic wrapper to an already deployed contract.
func bindFloodPlainL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FloodPlainL2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FloodPlainL2 *FloodPlainL2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FloodPlainL2.Contract.FloodPlainL2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FloodPlainL2 *FloodPlainL2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FloodPlainL2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FloodPlainL2 *FloodPlainL2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FloodPlainL2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FloodPlainL2 *FloodPlainL2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FloodPlainL2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FloodPlainL2 *FloodPlainL2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FloodPlainL2 *FloodPlainL2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _FloodPlainL2.Contract.DEFAULTADMINROLE(&_FloodPlainL2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FloodPlainL2.Contract.DEFAULTADMINROLE(&_FloodPlainL2.CallOpts)
}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_FloodPlainL2 *FloodPlainL2Caller) PERMIT2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "PERMIT2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_FloodPlainL2 *FloodPlainL2Session) PERMIT2() (common.Address, error) {
	return _FloodPlainL2.Contract.PERMIT2(&_FloodPlainL2.CallOpts)
}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_FloodPlainL2 *FloodPlainL2CallerSession) PERMIT2() (common.Address, error) {
	return _FloodPlainL2.Contract.PERMIT2(&_FloodPlainL2.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_FloodPlainL2 *FloodPlainL2Caller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_FloodPlainL2 *FloodPlainL2Session) DefaultAdmin() (common.Address, error) {
	return _FloodPlainL2.Contract.DefaultAdmin(&_FloodPlainL2.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_FloodPlainL2 *FloodPlainL2CallerSession) DefaultAdmin() (common.Address, error) {
	return _FloodPlainL2.Contract.DefaultAdmin(&_FloodPlainL2.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_FloodPlainL2 *FloodPlainL2Caller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_FloodPlainL2 *FloodPlainL2Session) DefaultAdminDelay() (*big.Int, error) {
	return _FloodPlainL2.Contract.DefaultAdminDelay(&_FloodPlainL2.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_FloodPlainL2 *FloodPlainL2CallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _FloodPlainL2.Contract.DefaultAdminDelay(&_FloodPlainL2.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_FloodPlainL2 *FloodPlainL2Caller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_FloodPlainL2 *FloodPlainL2Session) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _FloodPlainL2.Contract.DefaultAdminDelayIncreaseWait(&_FloodPlainL2.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_FloodPlainL2 *FloodPlainL2CallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _FloodPlainL2.Contract.DefaultAdminDelayIncreaseWait(&_FloodPlainL2.CallOpts)
}

// GetDecoder is a free data retrieval call binding the contract method 0xe77876cc.
//
// Solidity: function getDecoder(uint256 decoderId) view returns(address)
func (_FloodPlainL2 *FloodPlainL2Caller) GetDecoder(opts *bind.CallOpts, decoderId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getDecoder", decoderId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDecoder is a free data retrieval call binding the contract method 0xe77876cc.
//
// Solidity: function getDecoder(uint256 decoderId) view returns(address)
func (_FloodPlainL2 *FloodPlainL2Session) GetDecoder(decoderId *big.Int) (common.Address, error) {
	return _FloodPlainL2.Contract.GetDecoder(&_FloodPlainL2.CallOpts, decoderId)
}

// GetDecoder is a free data retrieval call binding the contract method 0xe77876cc.
//
// Solidity: function getDecoder(uint256 decoderId) view returns(address)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetDecoder(decoderId *big.Int) (common.Address, error) {
	return _FloodPlainL2.Contract.GetDecoder(&_FloodPlainL2.CallOpts, decoderId)
}

// GetEtchedOrder is a free data retrieval call binding the contract method 0x4d599400.
//
// Solidity: function getEtchedOrder(uint256 etchedOrderId) view returns(((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256),bytes))
func (_FloodPlainL2 *FloodPlainL2Caller) GetEtchedOrder(opts *bind.CallOpts, etchedOrderId *big.Int) (IFloodPlainOnChainOrdersOrderWithSignature, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getEtchedOrder", etchedOrderId)

	if err != nil {
		return *new(IFloodPlainOnChainOrdersOrderWithSignature), err
	}

	out0 := *abi.ConvertType(out[0], new(IFloodPlainOnChainOrdersOrderWithSignature)).(*IFloodPlainOnChainOrdersOrderWithSignature)

	return out0, err

}

// GetEtchedOrder is a free data retrieval call binding the contract method 0x4d599400.
//
// Solidity: function getEtchedOrder(uint256 etchedOrderId) view returns(((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256),bytes))
func (_FloodPlainL2 *FloodPlainL2Session) GetEtchedOrder(etchedOrderId *big.Int) (IFloodPlainOnChainOrdersOrderWithSignature, error) {
	return _FloodPlainL2.Contract.GetEtchedOrder(&_FloodPlainL2.CallOpts, etchedOrderId)
}

// GetEtchedOrder is a free data retrieval call binding the contract method 0x4d599400.
//
// Solidity: function getEtchedOrder(uint256 etchedOrderId) view returns(((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256),bytes))
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetEtchedOrder(etchedOrderId *big.Int) (IFloodPlainOnChainOrdersOrderWithSignature, error) {
	return _FloodPlainL2.Contract.GetEtchedOrder(&_FloodPlainL2.CallOpts, etchedOrderId)
}

// GetNonceStatus is a free data retrieval call binding the contract method 0xe9ba1e97.
//
// Solidity: function getNonceStatus(address user, uint256 nonce) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Caller) GetNonceStatus(opts *bind.CallOpts, user common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getNonceStatus", user, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNonceStatus is a free data retrieval call binding the contract method 0xe9ba1e97.
//
// Solidity: function getNonceStatus(address user, uint256 nonce) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Session) GetNonceStatus(user common.Address, nonce *big.Int) (bool, error) {
	return _FloodPlainL2.Contract.GetNonceStatus(&_FloodPlainL2.CallOpts, user, nonce)
}

// GetNonceStatus is a free data retrieval call binding the contract method 0xe9ba1e97.
//
// Solidity: function getNonceStatus(address user, uint256 nonce) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetNonceStatus(user common.Address, nonce *big.Int) (bool, error) {
	return _FloodPlainL2.Contract.GetNonceStatus(&_FloodPlainL2.CallOpts, user, nonce)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x1b8b792c.
//
// Solidity: function getOrderHash((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) pure returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Caller) GetOrderHash(opts *bind.CallOpts, order IFloodPlainOrder) ([32]byte, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getOrderHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0x1b8b792c.
//
// Solidity: function getOrderHash((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) pure returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Session) GetOrderHash(order IFloodPlainOrder) ([32]byte, error) {
	return _FloodPlainL2.Contract.GetOrderHash(&_FloodPlainL2.CallOpts, order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x1b8b792c.
//
// Solidity: function getOrderHash((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) pure returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetOrderHash(order IFloodPlainOrder) ([32]byte, error) {
	return _FloodPlainL2.Contract.GetOrderHash(&_FloodPlainL2.CallOpts, order)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x093de1d2.
//
// Solidity: function getOrderStatus((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Caller) GetOrderStatus(opts *bind.CallOpts, order IFloodPlainOrder) (bool, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getOrderStatus", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x093de1d2.
//
// Solidity: function getOrderStatus((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Session) GetOrderStatus(order IFloodPlainOrder) (bool, error) {
	return _FloodPlainL2.Contract.GetOrderStatus(&_FloodPlainL2.CallOpts, order)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x093de1d2.
//
// Solidity: function getOrderStatus((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetOrderStatus(order IFloodPlainOrder) (bool, error) {
	return _FloodPlainL2.Contract.GetOrderStatus(&_FloodPlainL2.CallOpts, order)
}

// GetOrderValidity is a free data retrieval call binding the contract method 0xa77dd3e4.
//
// Solidity: function getOrderValidity((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, address fulfiller, address caller, bytes extraData) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Caller) GetOrderValidity(opts *bind.CallOpts, order IFloodPlainOrder, fulfiller common.Address, caller common.Address, extraData []byte) (bool, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getOrderValidity", order, fulfiller, caller, extraData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOrderValidity is a free data retrieval call binding the contract method 0xa77dd3e4.
//
// Solidity: function getOrderValidity((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, address fulfiller, address caller, bytes extraData) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Session) GetOrderValidity(order IFloodPlainOrder, fulfiller common.Address, caller common.Address, extraData []byte) (bool, error) {
	return _FloodPlainL2.Contract.GetOrderValidity(&_FloodPlainL2.CallOpts, order, fulfiller, caller, extraData)
}

// GetOrderValidity is a free data retrieval call binding the contract method 0xa77dd3e4.
//
// Solidity: function getOrderValidity((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, address fulfiller, address caller, bytes extraData) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetOrderValidity(order IFloodPlainOrder, fulfiller common.Address, caller common.Address, extraData []byte) (bool, error) {
	return _FloodPlainL2.Contract.GetOrderValidity(&_FloodPlainL2.CallOpts, order, fulfiller, caller, extraData)
}

// GetOrderValidity0 is a free data retrieval call binding the contract method 0xfcb0caf2.
//
// Solidity: function getOrderValidity((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, address caller) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Caller) GetOrderValidity0(opts *bind.CallOpts, order IFloodPlainOrder, caller common.Address) (bool, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getOrderValidity0", order, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOrderValidity0 is a free data retrieval call binding the contract method 0xfcb0caf2.
//
// Solidity: function getOrderValidity((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, address caller) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Session) GetOrderValidity0(order IFloodPlainOrder, caller common.Address) (bool, error) {
	return _FloodPlainL2.Contract.GetOrderValidity0(&_FloodPlainL2.CallOpts, order, caller)
}

// GetOrderValidity0 is a free data retrieval call binding the contract method 0xfcb0caf2.
//
// Solidity: function getOrderValidity((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, address caller) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetOrderValidity0(order IFloodPlainOrder, caller common.Address) (bool, error) {
	return _FloodPlainL2.Contract.GetOrderValidity0(&_FloodPlainL2.CallOpts, order, caller)
}

// GetPermitHash is a free data retrieval call binding the contract method 0x729d540d.
//
// Solidity: function getPermitHash((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Caller) GetPermitHash(opts *bind.CallOpts, order IFloodPlainOrder) ([32]byte, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getPermitHash", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPermitHash is a free data retrieval call binding the contract method 0x729d540d.
//
// Solidity: function getPermitHash((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Session) GetPermitHash(order IFloodPlainOrder) ([32]byte, error) {
	return _FloodPlainL2.Contract.GetPermitHash(&_FloodPlainL2.CallOpts, order)
}

// GetPermitHash is a free data retrieval call binding the contract method 0x729d540d.
//
// Solidity: function getPermitHash((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order) view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetPermitHash(order IFloodPlainOrder) ([32]byte, error) {
	return _FloodPlainL2.Contract.GetPermitHash(&_FloodPlainL2.CallOpts, order)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FloodPlainL2.Contract.GetRoleAdmin(&_FloodPlainL2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FloodPlainL2 *FloodPlainL2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FloodPlainL2.Contract.GetRoleAdmin(&_FloodPlainL2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FloodPlainL2.Contract.HasRole(&_FloodPlainL2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FloodPlainL2.Contract.HasRole(&_FloodPlainL2.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FloodPlainL2 *FloodPlainL2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FloodPlainL2 *FloodPlainL2Session) Owner() (common.Address, error) {
	return _FloodPlainL2.Contract.Owner(&_FloodPlainL2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FloodPlainL2 *FloodPlainL2CallerSession) Owner() (common.Address, error) {
	return _FloodPlainL2.Contract.Owner(&_FloodPlainL2.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_FloodPlainL2 *FloodPlainL2Caller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_FloodPlainL2 *FloodPlainL2Session) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _FloodPlainL2.Contract.PendingDefaultAdmin(&_FloodPlainL2.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_FloodPlainL2 *FloodPlainL2CallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _FloodPlainL2.Contract.PendingDefaultAdmin(&_FloodPlainL2.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_FloodPlainL2 *FloodPlainL2Caller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_FloodPlainL2 *FloodPlainL2Session) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _FloodPlainL2.Contract.PendingDefaultAdminDelay(&_FloodPlainL2.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_FloodPlainL2 *FloodPlainL2CallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _FloodPlainL2.Contract.PendingDefaultAdminDelay(&_FloodPlainL2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FloodPlainL2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FloodPlainL2.Contract.SupportsInterface(&_FloodPlainL2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FloodPlainL2 *FloodPlainL2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FloodPlainL2.Contract.SupportsInterface(&_FloodPlainL2.CallOpts, interfaceId)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_FloodPlainL2 *FloodPlainL2Session) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.AcceptDefaultAdminTransfer(&_FloodPlainL2.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.AcceptDefaultAdminTransfer(&_FloodPlainL2.TransactOpts)
}

// AddDecoder is a paid mutator transaction binding the contract method 0x9d481b66.
//
// Solidity: function addDecoder(address decoder) returns(uint8 decoderId)
func (_FloodPlainL2 *FloodPlainL2Transactor) AddDecoder(opts *bind.TransactOpts, decoder common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "addDecoder", decoder)
}

// AddDecoder is a paid mutator transaction binding the contract method 0x9d481b66.
//
// Solidity: function addDecoder(address decoder) returns(uint8 decoderId)
func (_FloodPlainL2 *FloodPlainL2Session) AddDecoder(decoder common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.AddDecoder(&_FloodPlainL2.TransactOpts, decoder)
}

// AddDecoder is a paid mutator transaction binding the contract method 0x9d481b66.
//
// Solidity: function addDecoder(address decoder) returns(uint8 decoderId)
func (_FloodPlainL2 *FloodPlainL2TransactorSession) AddDecoder(decoder common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.AddDecoder(&_FloodPlainL2.TransactOpts, decoder)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_FloodPlainL2 *FloodPlainL2Session) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.BeginDefaultAdminTransfer(&_FloodPlainL2.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.BeginDefaultAdminTransfer(&_FloodPlainL2.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_FloodPlainL2 *FloodPlainL2Session) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.CancelDefaultAdminTransfer(&_FloodPlainL2.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.CancelDefaultAdminTransfer(&_FloodPlainL2.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_FloodPlainL2 *FloodPlainL2Session) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.ChangeDefaultAdminDelay(&_FloodPlainL2.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.ChangeDefaultAdminDelay(&_FloodPlainL2.TransactOpts, newDelay)
}

// EtchOrder is a paid mutator transaction binding the contract method 0x1d5473a2.
//
// Solidity: function etchOrder(((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256),bytes) orderWithSignature) returns(uint256 orderId)
func (_FloodPlainL2 *FloodPlainL2Transactor) EtchOrder(opts *bind.TransactOpts, orderWithSignature IFloodPlainOnChainOrdersOrderWithSignature) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "etchOrder", orderWithSignature)
}

// EtchOrder is a paid mutator transaction binding the contract method 0x1d5473a2.
//
// Solidity: function etchOrder(((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256),bytes) orderWithSignature) returns(uint256 orderId)
func (_FloodPlainL2 *FloodPlainL2Session) EtchOrder(orderWithSignature IFloodPlainOnChainOrdersOrderWithSignature) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.EtchOrder(&_FloodPlainL2.TransactOpts, orderWithSignature)
}

// EtchOrder is a paid mutator transaction binding the contract method 0x1d5473a2.
//
// Solidity: function etchOrder(((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256),bytes) orderWithSignature) returns(uint256 orderId)
func (_FloodPlainL2 *FloodPlainL2TransactorSession) EtchOrder(orderWithSignature IFloodPlainOnChainOrdersOrderWithSignature) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.EtchOrder(&_FloodPlainL2.TransactOpts, orderWithSignature)
}

// FulfillEtchedOrder is a paid mutator transaction binding the contract method 0xa15e7907.
//
// Solidity: function fulfillEtchedOrder(uint256 orderId, address fulfiller, bytes extraData) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) FulfillEtchedOrder(opts *bind.TransactOpts, orderId *big.Int, fulfiller common.Address, extraData []byte) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "fulfillEtchedOrder", orderId, fulfiller, extraData)
}

// FulfillEtchedOrder is a paid mutator transaction binding the contract method 0xa15e7907.
//
// Solidity: function fulfillEtchedOrder(uint256 orderId, address fulfiller, bytes extraData) returns()
func (_FloodPlainL2 *FloodPlainL2Session) FulfillEtchedOrder(orderId *big.Int, fulfiller common.Address, extraData []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FulfillEtchedOrder(&_FloodPlainL2.TransactOpts, orderId, fulfiller, extraData)
}

// FulfillEtchedOrder is a paid mutator transaction binding the contract method 0xa15e7907.
//
// Solidity: function fulfillEtchedOrder(uint256 orderId, address fulfiller, bytes extraData) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) FulfillEtchedOrder(orderId *big.Int, fulfiller common.Address, extraData []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FulfillEtchedOrder(&_FloodPlainL2.TransactOpts, orderId, fulfiller, extraData)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0x064d5bc3.
//
// Solidity: function fulfillOrder((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature, address fulfiller, bytes extraData) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) FulfillOrder(opts *bind.TransactOpts, order IFloodPlainOrder, signature []byte, fulfiller common.Address, extraData []byte) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "fulfillOrder", order, signature, fulfiller, extraData)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0x064d5bc3.
//
// Solidity: function fulfillOrder((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature, address fulfiller, bytes extraData) returns()
func (_FloodPlainL2 *FloodPlainL2Session) FulfillOrder(order IFloodPlainOrder, signature []byte, fulfiller common.Address, extraData []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FulfillOrder(&_FloodPlainL2.TransactOpts, order, signature, fulfiller, extraData)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0x064d5bc3.
//
// Solidity: function fulfillOrder((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature, address fulfiller, bytes extraData) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) FulfillOrder(order IFloodPlainOrder, signature []byte, fulfiller common.Address, extraData []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FulfillOrder(&_FloodPlainL2.TransactOpts, order, signature, fulfiller, extraData)
}

// FulfillOrder0 is a paid mutator transaction binding the contract method 0x6f01c5e2.
//
// Solidity: function fulfillOrder((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature) payable returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) FulfillOrder0(opts *bind.TransactOpts, order IFloodPlainOrder, signature []byte) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "fulfillOrder0", order, signature)
}

// FulfillOrder0 is a paid mutator transaction binding the contract method 0x6f01c5e2.
//
// Solidity: function fulfillOrder((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature) payable returns()
func (_FloodPlainL2 *FloodPlainL2Session) FulfillOrder0(order IFloodPlainOrder, signature []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FulfillOrder0(&_FloodPlainL2.TransactOpts, order, signature)
}

// FulfillOrder0 is a paid mutator transaction binding the contract method 0x6f01c5e2.
//
// Solidity: function fulfillOrder((address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature) payable returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) FulfillOrder0(order IFloodPlainOrder, signature []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.FulfillOrder0(&_FloodPlainL2.TransactOpts, order, signature)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.GrantRole(&_FloodPlainL2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.GrantRole(&_FloodPlainL2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.RenounceRole(&_FloodPlainL2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.RenounceRole(&_FloodPlainL2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.RevokeRole(&_FloodPlainL2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.RevokeRole(&_FloodPlainL2.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FloodPlainL2.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_FloodPlainL2 *FloodPlainL2Session) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.RollbackDefaultAdminDelay(&_FloodPlainL2.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.RollbackDefaultAdminDelay(&_FloodPlainL2.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FloodPlainL2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FloodPlainL2 *FloodPlainL2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.Fallback(&_FloodPlainL2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FloodPlainL2.Contract.Fallback(&_FloodPlainL2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FloodPlainL2 *FloodPlainL2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FloodPlainL2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FloodPlainL2 *FloodPlainL2Session) Receive() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.Receive(&_FloodPlainL2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FloodPlainL2 *FloodPlainL2TransactorSession) Receive() (*types.Transaction, error) {
	return _FloodPlainL2.Contract.Receive(&_FloodPlainL2.TransactOpts)
}

// FloodPlainL2DecoderAddedIterator is returned from FilterDecoderAdded and is used to iterate over the raw logs and unpacked data for DecoderAdded events raised by the FloodPlainL2 contract.
type FloodPlainL2DecoderAddedIterator struct {
	Event *FloodPlainL2DecoderAdded // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2DecoderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2DecoderAdded)
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
		it.Event = new(FloodPlainL2DecoderAdded)
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
func (it *FloodPlainL2DecoderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2DecoderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2DecoderAdded represents a DecoderAdded event raised by the FloodPlainL2 contract.
type FloodPlainL2DecoderAdded struct {
	DecoderId *big.Int
	Decoder   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecoderAdded is a free log retrieval operation binding the contract event 0x4e9fefd4c8c265adad06de042ad222441165306e8ac23ea525dee33f40463e64.
//
// Solidity: event DecoderAdded(uint256 indexed decoderId, address indexed decoder)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterDecoderAdded(opts *bind.FilterOpts, decoderId []*big.Int, decoder []common.Address) (*FloodPlainL2DecoderAddedIterator, error) {

	var decoderIdRule []interface{}
	for _, decoderIdItem := range decoderId {
		decoderIdRule = append(decoderIdRule, decoderIdItem)
	}
	var decoderRule []interface{}
	for _, decoderItem := range decoder {
		decoderRule = append(decoderRule, decoderItem)
	}

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "DecoderAdded", decoderIdRule, decoderRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2DecoderAddedIterator{contract: _FloodPlainL2.contract, event: "DecoderAdded", logs: logs, sub: sub}, nil
}

// WatchDecoderAdded is a free log subscription operation binding the contract event 0x4e9fefd4c8c265adad06de042ad222441165306e8ac23ea525dee33f40463e64.
//
// Solidity: event DecoderAdded(uint256 indexed decoderId, address indexed decoder)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchDecoderAdded(opts *bind.WatchOpts, sink chan<- *FloodPlainL2DecoderAdded, decoderId []*big.Int, decoder []common.Address) (event.Subscription, error) {

	var decoderIdRule []interface{}
	for _, decoderIdItem := range decoderId {
		decoderIdRule = append(decoderIdRule, decoderIdItem)
	}
	var decoderRule []interface{}
	for _, decoderItem := range decoder {
		decoderRule = append(decoderRule, decoderItem)
	}

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "DecoderAdded", decoderIdRule, decoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2DecoderAdded)
				if err := _FloodPlainL2.contract.UnpackLog(event, "DecoderAdded", log); err != nil {
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

// ParseDecoderAdded is a log parse operation binding the contract event 0x4e9fefd4c8c265adad06de042ad222441165306e8ac23ea525dee33f40463e64.
//
// Solidity: event DecoderAdded(uint256 indexed decoderId, address indexed decoder)
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseDecoderAdded(log types.Log) (*FloodPlainL2DecoderAdded, error) {
	event := new(FloodPlainL2DecoderAdded)
	if err := _FloodPlainL2.contract.UnpackLog(event, "DecoderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2DefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminDelayChangeCanceledIterator struct {
	Event *FloodPlainL2DefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2DefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2DefaultAdminDelayChangeCanceled)
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
		it.Event = new(FloodPlainL2DefaultAdminDelayChangeCanceled)
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
func (it *FloodPlainL2DefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2DefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2DefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*FloodPlainL2DefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2DefaultAdminDelayChangeCanceledIterator{contract: _FloodPlainL2.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *FloodPlainL2DefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2DefaultAdminDelayChangeCanceled)
				if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*FloodPlainL2DefaultAdminDelayChangeCanceled, error) {
	event := new(FloodPlainL2DefaultAdminDelayChangeCanceled)
	if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2DefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminDelayChangeScheduledIterator struct {
	Event *FloodPlainL2DefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2DefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2DefaultAdminDelayChangeScheduled)
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
		it.Event = new(FloodPlainL2DefaultAdminDelayChangeScheduled)
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
func (it *FloodPlainL2DefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2DefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2DefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*FloodPlainL2DefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2DefaultAdminDelayChangeScheduledIterator{contract: _FloodPlainL2.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *FloodPlainL2DefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2DefaultAdminDelayChangeScheduled)
				if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*FloodPlainL2DefaultAdminDelayChangeScheduled, error) {
	event := new(FloodPlainL2DefaultAdminDelayChangeScheduled)
	if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2DefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminTransferCanceledIterator struct {
	Event *FloodPlainL2DefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2DefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2DefaultAdminTransferCanceled)
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
		it.Event = new(FloodPlainL2DefaultAdminTransferCanceled)
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
func (it *FloodPlainL2DefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2DefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2DefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*FloodPlainL2DefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2DefaultAdminTransferCanceledIterator{contract: _FloodPlainL2.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *FloodPlainL2DefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2DefaultAdminTransferCanceled)
				if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseDefaultAdminTransferCanceled(log types.Log) (*FloodPlainL2DefaultAdminTransferCanceled, error) {
	event := new(FloodPlainL2DefaultAdminTransferCanceled)
	if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2DefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminTransferScheduledIterator struct {
	Event *FloodPlainL2DefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2DefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2DefaultAdminTransferScheduled)
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
		it.Event = new(FloodPlainL2DefaultAdminTransferScheduled)
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
func (it *FloodPlainL2DefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2DefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2DefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the FloodPlainL2 contract.
type FloodPlainL2DefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*FloodPlainL2DefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2DefaultAdminTransferScheduledIterator{contract: _FloodPlainL2.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *FloodPlainL2DefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2DefaultAdminTransferScheduled)
				if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseDefaultAdminTransferScheduled(log types.Log) (*FloodPlainL2DefaultAdminTransferScheduled, error) {
	event := new(FloodPlainL2DefaultAdminTransferScheduled)
	if err := _FloodPlainL2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2OrderEtchedIterator is returned from FilterOrderEtched and is used to iterate over the raw logs and unpacked data for OrderEtched events raised by the FloodPlainL2 contract.
type FloodPlainL2OrderEtchedIterator struct {
	Event *FloodPlainL2OrderEtched // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2OrderEtchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2OrderEtched)
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
		it.Event = new(FloodPlainL2OrderEtched)
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
func (it *FloodPlainL2OrderEtchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2OrderEtchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2OrderEtched represents a OrderEtched event raised by the FloodPlainL2 contract.
type FloodPlainL2OrderEtched struct {
	OrderId   *big.Int
	OrderHash [32]byte
	Order     IFloodPlainOrder
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderEtched is a free log retrieval operation binding the contract event 0x7257f25711a9a80142a813ae2b3eeb94db36ad0f0f83dfc2512833ab23c33f2a.
//
// Solidity: event OrderEtched(uint256 indexed orderId, bytes32 indexed orderHash, (address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterOrderEtched(opts *bind.FilterOpts, orderId []*big.Int, orderHash [][32]byte) (*FloodPlainL2OrderEtchedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "OrderEtched", orderIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2OrderEtchedIterator{contract: _FloodPlainL2.contract, event: "OrderEtched", logs: logs, sub: sub}, nil
}

// WatchOrderEtched is a free log subscription operation binding the contract event 0x7257f25711a9a80142a813ae2b3eeb94db36ad0f0f83dfc2512833ab23c33f2a.
//
// Solidity: event OrderEtched(uint256 indexed orderId, bytes32 indexed orderHash, (address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchOrderEtched(opts *bind.WatchOpts, sink chan<- *FloodPlainL2OrderEtched, orderId []*big.Int, orderHash [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "OrderEtched", orderIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2OrderEtched)
				if err := _FloodPlainL2.contract.UnpackLog(event, "OrderEtched", log); err != nil {
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

// ParseOrderEtched is a log parse operation binding the contract event 0x7257f25711a9a80142a813ae2b3eeb94db36ad0f0f83dfc2512833ab23c33f2a.
//
// Solidity: event OrderEtched(uint256 indexed orderId, bytes32 indexed orderHash, (address,address,(address,uint256)[],(address,uint256)[],uint256,uint256) order, bytes signature)
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseOrderEtched(log types.Log) (*FloodPlainL2OrderEtched, error) {
	event := new(FloodPlainL2OrderEtched)
	if err := _FloodPlainL2.contract.UnpackLog(event, "OrderEtched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2OrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the FloodPlainL2 contract.
type FloodPlainL2OrderFulfilledIterator struct {
	Event *FloodPlainL2OrderFulfilled // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2OrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2OrderFulfilled)
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
		it.Event = new(FloodPlainL2OrderFulfilled)
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
func (it *FloodPlainL2OrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2OrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2OrderFulfilled represents a OrderFulfilled event raised by the FloodPlainL2 contract.
type FloodPlainL2OrderFulfilled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Fulfiller common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xf524bb99c8acc2d7dcf2632c53dac061fafb00b208f763b6552c96d805a452a7.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderHash, address indexed offerer, address indexed fulfiller)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterOrderFulfilled(opts *bind.FilterOpts, orderHash [][32]byte, offerer []common.Address, fulfiller []common.Address) (*FloodPlainL2OrderFulfilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var fulfillerRule []interface{}
	for _, fulfillerItem := range fulfiller {
		fulfillerRule = append(fulfillerRule, fulfillerItem)
	}

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "OrderFulfilled", orderHashRule, offererRule, fulfillerRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2OrderFulfilledIterator{contract: _FloodPlainL2.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xf524bb99c8acc2d7dcf2632c53dac061fafb00b208f763b6552c96d805a452a7.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderHash, address indexed offerer, address indexed fulfiller)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *FloodPlainL2OrderFulfilled, orderHash [][32]byte, offerer []common.Address, fulfiller []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var fulfillerRule []interface{}
	for _, fulfillerItem := range fulfiller {
		fulfillerRule = append(fulfillerRule, fulfillerItem)
	}

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "OrderFulfilled", orderHashRule, offererRule, fulfillerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2OrderFulfilled)
				if err := _FloodPlainL2.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// ParseOrderFulfilled is a log parse operation binding the contract event 0xf524bb99c8acc2d7dcf2632c53dac061fafb00b208f763b6552c96d805a452a7.
//
// Solidity: event OrderFulfilled(bytes32 indexed orderHash, address indexed offerer, address indexed fulfiller)
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseOrderFulfilled(log types.Log) (*FloodPlainL2OrderFulfilled, error) {
	event := new(FloodPlainL2OrderFulfilled)
	if err := _FloodPlainL2.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FloodPlainL2 contract.
type FloodPlainL2RoleAdminChangedIterator struct {
	Event *FloodPlainL2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2RoleAdminChanged)
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
		it.Event = new(FloodPlainL2RoleAdminChanged)
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
func (it *FloodPlainL2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2RoleAdminChanged represents a RoleAdminChanged event raised by the FloodPlainL2 contract.
type FloodPlainL2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FloodPlainL2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2RoleAdminChangedIterator{contract: _FloodPlainL2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FloodPlainL2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2RoleAdminChanged)
				if err := _FloodPlainL2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseRoleAdminChanged(log types.Log) (*FloodPlainL2RoleAdminChanged, error) {
	event := new(FloodPlainL2RoleAdminChanged)
	if err := _FloodPlainL2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FloodPlainL2 contract.
type FloodPlainL2RoleGrantedIterator struct {
	Event *FloodPlainL2RoleGranted // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2RoleGranted)
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
		it.Event = new(FloodPlainL2RoleGranted)
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
func (it *FloodPlainL2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2RoleGranted represents a RoleGranted event raised by the FloodPlainL2 contract.
type FloodPlainL2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FloodPlainL2RoleGrantedIterator, error) {

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

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2RoleGrantedIterator{contract: _FloodPlainL2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FloodPlainL2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2RoleGranted)
				if err := _FloodPlainL2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseRoleGranted(log types.Log) (*FloodPlainL2RoleGranted, error) {
	event := new(FloodPlainL2RoleGranted)
	if err := _FloodPlainL2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FloodPlainL2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FloodPlainL2 contract.
type FloodPlainL2RoleRevokedIterator struct {
	Event *FloodPlainL2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *FloodPlainL2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloodPlainL2RoleRevoked)
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
		it.Event = new(FloodPlainL2RoleRevoked)
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
func (it *FloodPlainL2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloodPlainL2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloodPlainL2RoleRevoked represents a RoleRevoked event raised by the FloodPlainL2 contract.
type FloodPlainL2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FloodPlainL2 *FloodPlainL2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FloodPlainL2RoleRevokedIterator, error) {

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

	logs, sub, err := _FloodPlainL2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FloodPlainL2RoleRevokedIterator{contract: _FloodPlainL2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FloodPlainL2 *FloodPlainL2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FloodPlainL2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FloodPlainL2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloodPlainL2RoleRevoked)
				if err := _FloodPlainL2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FloodPlainL2 *FloodPlainL2Filterer) ParseRoleRevoked(log types.Log) (*FloodPlainL2RoleRevoked, error) {
	event := new(FloodPlainL2RoleRevoked)
	if err := _FloodPlainL2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
