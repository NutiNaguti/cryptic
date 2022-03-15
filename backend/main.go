package main

import (
	"cryptic/config"
	"cryptic/ethereum"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
}

func main() {
	conf := config.New()

	ethereum.InitClient(conf.RPC_URL)
	ethereum.InitERC20(conf.ERC20_ADDRESS)
	router := gin.Default()

	go router.GET("/health", healthCheck)
	go router.GET("/history", getTxs)
	go router.GET("/balance", getBalance)

	go router.POST("/donate", donate)

	router.Run()
}

func healthCheck(c *gin.Context) {
	blockNumber := ethereum.GetLastBlockNumber()
	c.IndentedJSON(http.StatusOK, blockNumber)
}

func getBalance(c *gin.Context) {
	address := c.Param("address")
	balance := ethereum.GetBalance(address)
	c.IndentedJSON(http.StatusOK, balance)
}

func donate(c *gin.Context) {

}

func getTxs(c *gin.Context) {

}
