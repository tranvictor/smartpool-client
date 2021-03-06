// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package geth

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestPoolABI is the input ABI used to generate the binding from.
const TestPoolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newVersionReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getShareIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"calcPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"canRegister\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"error\",\"type\":\"string\"}],\"name\":\"submitClientError\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"augCountersBranch\",\"type\":\"uint256[]\"},{\"name\":\"augHashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"calcPaymentShareUnits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes32\"},{\"name\":\"nonceLe\",\"type\":\"bytes8\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"branchSize\",\"type\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"hashimoto\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochData\",\"outputs\":[{\"name\":\"merkleRoot\",\"type\":\"uint128\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint64\"},{\"name\":\"branchDepth\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getClaimSeed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getMinerId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"declareNewerVersion\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\"},{\"name\":\"augMerkle\",\"type\":\"uint256\"}],\"name\":\"submitClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"budgetPerShareUnit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"existingIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"merkleRoot\",\"type\":\"uint128[]\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint64[]\"},{\"name\":\"branchDepth\",\"type\":\"uint64[]\"},{\"name\":\"epoch\",\"type\":\"uint256[]\"}],\"name\":\"setEpochData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentUnitsPer1000Blocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"budget\",\"type\":\"uint256\"}],\"name\":\"setBudgetPerShareUnit\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"numChars\",\"type\":\"uint256\"}],\"name\":\"to62Encoding\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[3]\"}],\"payable\":true,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ErrorLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ErrorNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SubmitClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetEpochData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyExtraData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"string\"}],\"name\":\"ClientError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"VerifyAgt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"Result\",\"type\":\"event\"}]"

// TestPool is an auto generated Go binding around an Ethereum contract.
type TestPool struct {
	TestPoolCaller     // Read-only binding to the contract
	TestPoolTransactor // Write-only binding to the contract
}

// TestPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPoolSession struct {
	Contract     *TestPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPoolCallerSession struct {
	Contract *TestPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPoolTransactorSession struct {
	Contract     *TestPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPoolRaw struct {
	Contract *TestPool // Generic contract binding to access the raw methods on
}

// TestPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPoolCallerRaw struct {
	Contract *TestPoolCaller // Generic read-only contract binding to access the raw methods on
}

// TestPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPoolTransactorRaw struct {
	Contract *TestPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPool creates a new instance of TestPool, bound to a specific deployed contract.
func NewTestPool(address common.Address, backend bind.ContractBackend) (*TestPool, error) {
	contract, err := bindTestPool(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPool{TestPoolCaller: TestPoolCaller{contract: contract}, TestPoolTransactor: TestPoolTransactor{contract: contract}}, nil
}

// NewTestPoolCaller creates a new read-only instance of TestPool, bound to a specific deployed contract.
func NewTestPoolCaller(address common.Address, caller bind.ContractCaller) (*TestPoolCaller, error) {
	contract, err := bindTestPool(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestPoolCaller{contract: contract}, nil
}

// NewTestPoolTransactor creates a new write-only instance of TestPool, bound to a specific deployed contract.
func NewTestPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPoolTransactor, error) {
	contract, err := bindTestPool(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestPoolTransactor{contract: contract}, nil
}

// bindTestPool binds a generic wrapper to an already deployed contract.
func bindTestPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPool *TestPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestPool.Contract.TestPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPool *TestPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.Contract.TestPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPool *TestPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPool.Contract.TestPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPool *TestPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPool *TestPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPool *TestPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPool.Contract.contract.Transact(opts, method, params...)
}

// BudgetPerShareUnit is a free data retrieval call binding the contract method 0xe97857db.
//
// Solidity: function budgetPerShareUnit() constant returns(uint256)
func (_TestPool *TestPoolCaller) BudgetPerShareUnit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "budgetPerShareUnit")
	return *ret0, err
}

// BudgetPerShareUnit is a free data retrieval call binding the contract method 0xe97857db.
//
// Solidity: function budgetPerShareUnit() constant returns(uint256)
func (_TestPool *TestPoolSession) BudgetPerShareUnit() (*big.Int, error) {
	return _TestPool.Contract.BudgetPerShareUnit(&_TestPool.CallOpts)
}

// BudgetPerShareUnit is a free data retrieval call binding the contract method 0xe97857db.
//
// Solidity: function budgetPerShareUnit() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) BudgetPerShareUnit() (*big.Int, error) {
	return _TestPool.Contract.BudgetPerShareUnit(&_TestPool.CallOpts)
}

// CalcPayment is a free data retrieval call binding the contract method 0x17a720ac.
//
// Solidity: function calcPayment(numShares uint256, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) CalcPayment(opts *bind.CallOpts, numShares *big.Int, difficulty *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "calcPayment", numShares, difficulty)
	return *ret0, err
}

// CalcPayment is a free data retrieval call binding the contract method 0x17a720ac.
//
// Solidity: function calcPayment(numShares uint256, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) CalcPayment(numShares *big.Int, difficulty *big.Int) (*big.Int, error) {
	return _TestPool.Contract.CalcPayment(&_TestPool.CallOpts, numShares, difficulty)
}

// CalcPayment is a free data retrieval call binding the contract method 0x17a720ac.
//
// Solidity: function calcPayment(numShares uint256, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) CalcPayment(numShares *big.Int, difficulty *big.Int) (*big.Int, error) {
	return _TestPool.Contract.CalcPayment(&_TestPool.CallOpts, numShares, difficulty)
}

// CalcPaymentShareUnits is a free data retrieval call binding the contract method 0x3c00146b.
//
// Solidity: function calcPaymentShareUnits(numShares uint256, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) CalcPaymentShareUnits(opts *bind.CallOpts, numShares *big.Int, difficulty *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "calcPaymentShareUnits", numShares, difficulty)
	return *ret0, err
}

// CalcPaymentShareUnits is a free data retrieval call binding the contract method 0x3c00146b.
//
// Solidity: function calcPaymentShareUnits(numShares uint256, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) CalcPaymentShareUnits(numShares *big.Int, difficulty *big.Int) (*big.Int, error) {
	return _TestPool.Contract.CalcPaymentShareUnits(&_TestPool.CallOpts, numShares, difficulty)
}

// CalcPaymentShareUnits is a free data retrieval call binding the contract method 0x3c00146b.
//
// Solidity: function calcPaymentShareUnits(numShares uint256, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) CalcPaymentShareUnits(numShares *big.Int, difficulty *big.Int) (*big.Int, error) {
	return _TestPool.Contract.CalcPaymentShareUnits(&_TestPool.CallOpts, numShares, difficulty)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_TestPool *TestPoolCaller) CanRegister(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "canRegister", sender)
	return *ret0, err
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_TestPool *TestPoolSession) CanRegister(sender common.Address) (bool, error) {
	return _TestPool.Contract.CanRegister(&_TestPool.CallOpts, sender)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_TestPool *TestPoolCallerSession) CanRegister(sender common.Address) (bool, error) {
	return _TestPool.Contract.CanRegister(&_TestPool.CallOpts, sender)
}

// EpochData is a free data retrieval call binding the contract method 0x6e821b2e.
//
// Solidity: function epochData( uint256) constant returns(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64)
func (_TestPool *TestPoolCaller) EpochData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MerkleRoot             *big.Int
	FullSizeIn128Resultion uint64
	BranchDepth            uint64
}, error) {
	ret := new(struct {
		MerkleRoot             *big.Int
		FullSizeIn128Resultion uint64
		BranchDepth            uint64
	})
	out := ret
	err := _TestPool.contract.Call(opts, out, "epochData", arg0)
	return *ret, err
}

// EpochData is a free data retrieval call binding the contract method 0x6e821b2e.
//
// Solidity: function epochData( uint256) constant returns(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64)
func (_TestPool *TestPoolSession) EpochData(arg0 *big.Int) (struct {
	MerkleRoot             *big.Int
	FullSizeIn128Resultion uint64
	BranchDepth            uint64
}, error) {
	return _TestPool.Contract.EpochData(&_TestPool.CallOpts, arg0)
}

// EpochData is a free data retrieval call binding the contract method 0x6e821b2e.
//
// Solidity: function epochData( uint256) constant returns(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64)
func (_TestPool *TestPoolCallerSession) EpochData(arg0 *big.Int) (struct {
	MerkleRoot             *big.Int
	FullSizeIn128Resultion uint64
	BranchDepth            uint64
}, error) {
	return _TestPool.Contract.EpochData(&_TestPool.CallOpts, arg0)
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_TestPool *TestPoolCaller) ExistingIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "existingIds", arg0)
	return *ret0, err
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_TestPool *TestPoolSession) ExistingIds(arg0 [32]byte) (bool, error) {
	return _TestPool.Contract.ExistingIds(&_TestPool.CallOpts, arg0)
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_TestPool *TestPoolCallerSession) ExistingIds(arg0 [32]byte) (bool, error) {
	return _TestPool.Contract.ExistingIds(&_TestPool.CallOpts, arg0)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_TestPool *TestPoolCaller) GetClaimSeed(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getClaimSeed", sender)
	return *ret0, err
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_TestPool *TestPoolSession) GetClaimSeed(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts, sender)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetClaimSeed(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts, sender)
}

// GetEpochData is a free data retrieval call binding the contract method 0x859e7d32.
//
// Solidity: function getEpochData(epoch uint256) constant returns(uint256[3])
func (_TestPool *TestPoolCaller) GetEpochData(opts *bind.CallOpts, epoch *big.Int) ([3]*big.Int, error) {
	var (
		ret0 = new([3]*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getEpochData", epoch)
	return *ret0, err
}

// GetEpochData is a free data retrieval call binding the contract method 0x859e7d32.
//
// Solidity: function getEpochData(epoch uint256) constant returns(uint256[3])
func (_TestPool *TestPoolSession) GetEpochData(epoch *big.Int) ([3]*big.Int, error) {
	return _TestPool.Contract.GetEpochData(&_TestPool.CallOpts, epoch)
}

// GetEpochData is a free data retrieval call binding the contract method 0x859e7d32.
//
// Solidity: function getEpochData(epoch uint256) constant returns(uint256[3])
func (_TestPool *TestPoolCallerSession) GetEpochData(epoch *big.Int) ([3]*big.Int, error) {
	return _TestPool.Contract.GetEpochData(&_TestPool.CallOpts, epoch)
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_TestPool *TestPoolCaller) GetMinerId(opts *bind.CallOpts, sender common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getMinerId", sender)
	return *ret0, err
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_TestPool *TestPoolSession) GetMinerId(sender common.Address) ([32]byte, error) {
	return _TestPool.Contract.GetMinerId(&_TestPool.CallOpts, sender)
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_TestPool *TestPoolCallerSession) GetMinerId(sender common.Address) ([32]byte, error) {
	return _TestPool.Contract.GetMinerId(&_TestPool.CallOpts, sender)
}

// GetShareIndex is a free data retrieval call binding the contract method 0x0f136527.
//
// Solidity: function getShareIndex(sender address) constant returns(uint256)
func (_TestPool *TestPoolCaller) GetShareIndex(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getShareIndex", sender)
	return *ret0, err
}

// GetShareIndex is a free data retrieval call binding the contract method 0x0f136527.
//
// Solidity: function getShareIndex(sender address) constant returns(uint256)
func (_TestPool *TestPoolSession) GetShareIndex(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetShareIndex(&_TestPool.CallOpts, sender)
}

// GetShareIndex is a free data retrieval call binding the contract method 0x0f136527.
//
// Solidity: function getShareIndex(sender address) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetShareIndex(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetShareIndex(&_TestPool.CallOpts, sender)
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Hashimoto(opts *bind.CallOpts, header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "hashimoto", header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
	return *ret0, err
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Hashimoto(header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hashimoto(&_TestPool.CallOpts, header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Hashimoto(header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hashimoto(&_TestPool.CallOpts, header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_TestPool *TestPoolCaller) IsRegistered(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "isRegistered", sender)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_TestPool *TestPoolSession) IsRegistered(sender common.Address) (bool, error) {
	return _TestPool.Contract.IsRegistered(&_TestPool.CallOpts, sender)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_TestPool *TestPoolCallerSession) IsRegistered(sender common.Address) (bool, error) {
	return _TestPool.Contract.IsRegistered(&_TestPool.CallOpts, sender)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolCaller) NewVersionReleased(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "newVersionReleased")
	return *ret0, err
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolSession) NewVersionReleased() (bool, error) {
	return _TestPool.Contract.NewVersionReleased(&_TestPool.CallOpts)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolCallerSession) NewVersionReleased() (bool, error) {
	return _TestPool.Contract.NewVersionReleased(&_TestPool.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_TestPool *TestPoolCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_TestPool *TestPoolSession) Owners(arg0 common.Address) (bool, error) {
	return _TestPool.Contract.Owners(&_TestPool.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_TestPool *TestPoolCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _TestPool.Contract.Owners(&_TestPool.CallOpts, arg0)
}

// PaymentUnitsPer1000Blocks is a free data retrieval call binding the contract method 0xf4d00f99.
//
// Solidity: function paymentUnitsPer1000Blocks( uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) PaymentUnitsPer1000Blocks(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "paymentUnitsPer1000Blocks", arg0)
	return *ret0, err
}

// PaymentUnitsPer1000Blocks is a free data retrieval call binding the contract method 0xf4d00f99.
//
// Solidity: function paymentUnitsPer1000Blocks( uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) PaymentUnitsPer1000Blocks(arg0 *big.Int) (*big.Int, error) {
	return _TestPool.Contract.PaymentUnitsPer1000Blocks(&_TestPool.CallOpts, arg0)
}

// PaymentUnitsPer1000Blocks is a free data retrieval call binding the contract method 0xf4d00f99.
//
// Solidity: function paymentUnitsPer1000Blocks( uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) PaymentUnitsPer1000Blocks(arg0 *big.Int) (*big.Int, error) {
	return _TestPool.Contract.PaymentUnitsPer1000Blocks(&_TestPool.CallOpts, arg0)
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_TestPool *TestPoolCaller) To62Encoding(opts *bind.CallOpts, id *big.Int, numChars *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "to62Encoding", id, numChars)
	return *ret0, err
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_TestPool *TestPoolSession) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return _TestPool.Contract.To62Encoding(&_TestPool.CallOpts, id, numChars)
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_TestPool *TestPoolCallerSession) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return _TestPool.Contract.To62Encoding(&_TestPool.CallOpts, id, numChars)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolSession) Version() (string, error) {
	return _TestPool.Contract.Version(&_TestPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolCallerSession) Version() (string, error) {
	return _TestPool.Contract.Version(&_TestPool.CallOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolTransactor) DeclareNewerVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "declareNewerVersion")
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _TestPool.Contract.DeclareNewerVersion(&_TestPool.TransactOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolTransactorSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _TestPool.Contract.DeclareNewerVersion(&_TestPool.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_TestPool *TestPoolTransactor) Register(opts *bind.TransactOpts, paymentAddress common.Address) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "register", paymentAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_TestPool *TestPoolSession) Register(paymentAddress common.Address) (*types.Transaction, error) {
	return _TestPool.Contract.Register(&_TestPool.TransactOpts, paymentAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_TestPool *TestPoolTransactorSession) Register(paymentAddress common.Address) (*types.Transaction, error) {
	return _TestPool.Contract.Register(&_TestPool.TransactOpts, paymentAddress)
}

// SetBudgetPerShareUnit is a paid mutator transaction binding the contract method 0xfad574be.
//
// Solidity: function setBudgetPerShareUnit(budget uint256) returns()
func (_TestPool *TestPoolTransactor) SetBudgetPerShareUnit(opts *bind.TransactOpts, budget *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "setBudgetPerShareUnit", budget)
}

// SetBudgetPerShareUnit is a paid mutator transaction binding the contract method 0xfad574be.
//
// Solidity: function setBudgetPerShareUnit(budget uint256) returns()
func (_TestPool *TestPoolSession) SetBudgetPerShareUnit(budget *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetBudgetPerShareUnit(&_TestPool.TransactOpts, budget)
}

// SetBudgetPerShareUnit is a paid mutator transaction binding the contract method 0xfad574be.
//
// Solidity: function setBudgetPerShareUnit(budget uint256) returns()
func (_TestPool *TestPoolTransactorSession) SetBudgetPerShareUnit(budget *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetBudgetPerShareUnit(&_TestPool.TransactOpts, budget)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xf2cfe267.
//
// Solidity: function setEpochData(merkleRoot uint128[], fullSizeIn128Resultion uint64[], branchDepth uint64[], epoch uint256[]) returns()
func (_TestPool *TestPoolTransactor) SetEpochData(opts *bind.TransactOpts, merkleRoot []*big.Int, fullSizeIn128Resultion []uint64, branchDepth []uint64, epoch []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "setEpochData", merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xf2cfe267.
//
// Solidity: function setEpochData(merkleRoot uint128[], fullSizeIn128Resultion uint64[], branchDepth uint64[], epoch uint256[]) returns()
func (_TestPool *TestPoolSession) SetEpochData(merkleRoot []*big.Int, fullSizeIn128Resultion []uint64, branchDepth []uint64, epoch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetEpochData(&_TestPool.TransactOpts, merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xf2cfe267.
//
// Solidity: function setEpochData(merkleRoot uint128[], fullSizeIn128Resultion uint64[], branchDepth uint64[], epoch uint256[]) returns()
func (_TestPool *TestPoolTransactorSession) SetEpochData(merkleRoot []*big.Int, fullSizeIn128Resultion []uint64, branchDepth []uint64, epoch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetEpochData(&_TestPool.TransactOpts, merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolTransactor) SubmitClaim(opts *bind.TransactOpts, numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "submitClaim", numShares, difficulty, min, max, augMerkle)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augMerkle)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolTransactorSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augMerkle)
}

// SubmitClientError is a paid mutator transaction binding the contract method 0x33401c58.
//
// Solidity: function submitClientError(error string) returns()
func (_TestPool *TestPoolTransactor) SubmitClientError(opts *bind.TransactOpts, error string) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "submitClientError", error)
}

// SubmitClientError is a paid mutator transaction binding the contract method 0x33401c58.
//
// Solidity: function submitClientError(error string) returns()
func (_TestPool *TestPoolSession) SubmitClientError(error string) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClientError(&_TestPool.TransactOpts, error)
}

// SubmitClientError is a paid mutator transaction binding the contract method 0x33401c58.
//
// Solidity: function submitClientError(error string) returns()
func (_TestPool *TestPoolTransactorSession) SubmitClientError(error string) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClientError(&_TestPool.TransactOpts, error)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolTransactor) VerifyClaim(opts *bind.TransactOpts, rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "verifyClaim", rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolTransactorSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}
