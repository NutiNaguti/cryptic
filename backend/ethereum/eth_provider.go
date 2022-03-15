package ethereum

import (
	"cryptic/ethereum/contracts"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
var erc20 *contracts.Contracts

func InitClient() {
	cl, err := ethclient.Dial("https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		panic(err)
	}

	client = cl
}

func InitERC20(hex string) {
	address := common.HexToAddress(hex)
	ctr, err := contracts.NewContracts(address, client)
	if err != nil {
		log.Default().Fatalln(err)
	}
	erc20 = ctr
}
