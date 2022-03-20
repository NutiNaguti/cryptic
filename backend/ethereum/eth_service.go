package ethereum

import (
	"context"
	"log"
	"math"
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

func GetERC20Balance(hexAddress string, hexWallet string) (balance *TokenBalance) {
	address := common.HexToAddress(hexAddress)
	wallet := common.HexToAddress(hexWallet)
	tb := &TokenBalance{
		Contract: address,
		Wallet:   wallet,
		Decimals: 0,
		Balance:  big.NewInt(0),
		ctx:      context.TODO(),
	}
	err := tb.query()
	if err != nil {
		log.Fatalln(err)
	}

	balance = tb
	return
}

func GetBalance(hex string) *big.Float {
	log.Println(hex)
	ctx := context.Background()
	address := common.HexToAddress(hex)
	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}

func minERC20(hex string, amount int64) {
	address := common.HexToAddress(hex)
	if amount < 0 {
		amount *= -1
	}
	erc20.Mint(nil, address, big.NewInt(amount))
}
