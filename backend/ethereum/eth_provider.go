package ethereum

import (
	"github.com/umbracle/ethgo/jsonrpc"
)

var client *jsonrpc.Client

func getClient() *jsonrpc.Client {
	if client == nil {
		c, err := jsonrpc.NewClient("wss://eth-ropsten.alchemyapi.io/v2/Hh2at-WBCmQNXsvbdah59LY7IPwtpEAs")
		if err != nil {
			panic(err)
		}
		client = c
	}

	return client
}
