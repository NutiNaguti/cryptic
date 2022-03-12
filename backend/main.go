package main

import (
	"cryptic/ethereum"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", healthCheck)
	router.GET("/history", getTxs)
	router.GET("/balance", getBalance)

	router.POST("/tx", sendTx)

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

func sendTx(c *gin.Context) {

}

func getTxs(c *gin.Context) {

}
