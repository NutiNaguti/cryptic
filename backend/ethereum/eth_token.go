package ethereum

import (
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const tokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"nme\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_releaseTime\",\"type\":\"uint256\"}],\"name\":\"mintTimelocked\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

type token struct {
	caller
	transactor
}

type caller struct {
	contract *bind.BoundContract
}

type transactor struct {
	contract *bind.BoundContract
}

type session struct {
	Contract     *token
	CallerOpts   bind.CallOpts
	TransactOpts bind.TransactOpts
}

type callerSession struct {
	Contract *caller
	CallOpts bind.CallOpts
}

type transactionSession struct {
	Contract     *transactor
	TransactOpts bind.TransactOpts
}

type raw struct {
	Contract *token
}

type callerRaw struct {
	Contract *caller
}

type transactorRaw struct {
	Contract *transactor
}

func newTokenCaller(address common.Address, c bind.ContractCaller) (*caller, error) {
	contract, err := bindToken(address, c, nil)
	if err != nil {
		return nil, err
	}
	return &caller{contract: contract}, nil
}

func bindToken(address common.Address, c bind.ContractCaller, t bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, c, t, nil), nil
}

func (_Token *caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", _owner)
	out0 := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	log.Println(out0)
	return &out0, err
}

func (_Token *caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "decimals")
	out0 := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	log.Println(out0)
	return &out0, err
}

func (_Token *caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	log.Println(out0)
	return out0, err
}

func (_Token *caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	log.Println(out0)
	return out0, err
}

func (tb *TokenBalance) query() error {
	var err error

	token, err := newTokenCaller(tb.Contract, client)
	if err != nil {
		log.Println("Faild to instantiate a token contract: ", err)
		return err
	}

	block, err := client.BlockByNumber(tb.ctx, nil)
	if err != nil {
		log.Println("Failed to get current block number: ", err)
		return err
	}
	tb.Block = block.Number().Int64()

	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Println("Failed to get decimals from contract: ", err)
		//return err
	}
	tb.Decimals = decimals.Int64()

	tb.ETH, err = client.BalanceAt(tb.ctx, tb.Wallet, nil)
	if err != nil {
		log.Println("Failed to get ethereum balance from address: ", err)
		return err
	}

	tb.Balance, err = token.BalanceOf(nil, tb.Wallet)
	if err != nil {
		log.Println("Faild to get balance from address: ", err)
		return err
	}

	tb.Symbol, err = token.Symbol(nil)
	if err != nil {
		log.Println("Faild to get symbol from contract: ", err)
		return err
	}

	tb.Name, err = token.Name(nil)
	if err != nil {
		log.Println("Faild to get name from contract: ", err)
		return err
	}

	return err
}
