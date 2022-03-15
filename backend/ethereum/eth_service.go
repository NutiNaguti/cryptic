package ethereum

import (
	"context"
	"log"
	"math/big"
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
