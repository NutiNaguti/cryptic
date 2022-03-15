package ethereum

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GetLastBlockNumber() (number uint64) {
	ctx := context.Background()
	number, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetBalance(hex string) (balance *big.Int) {
	return
}

func minERC20(hex string, amount int64) {
	address := common.HexToAddress(hex)
	if amount < 0 {
		amount *= -1
	}
	erc20.Mint(nil, address, big.NewInt(amount))
}
