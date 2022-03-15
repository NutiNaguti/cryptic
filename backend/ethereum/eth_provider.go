package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func InitClient() {
	cl, err := ethclient.Dial("https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		panic(err)
	}

	client = cl
}
