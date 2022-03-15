package main

import (
	"cryptic/ethereum"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found!")
	}
}

func main() {
	erc20address, exists := os.LookupEnv("ERC20_ADDRESS")
	if !exists {
		log.Fatalln("ERC20_ADDRESS is not exists")
	}

	ethereum.InitClient()
	ethereum.InitERC20(erc20address)
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
