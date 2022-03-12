package ethereum

import (
	"math/big"

	"github.com/umbracle/ethgo"
)

func GetLastBlockNumber() (number uint64) {
	number, err := getClient().Eth().BlockNumber()
	if err != nil {
		panic(nil)
	}

	return
}

func GetBalance(hex string) (balance *big.Int) {
	address := ethgo.HexToAddress(hex)
	blockNumber := ethgo.HexToHash(hex)
	balance, err := getClient().Eth().GetBalance(address, blockNumber)
	if err != nil {
		panic(err)
	}
	return
}
