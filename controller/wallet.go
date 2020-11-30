package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	// token := c.PostForm("token")
	// operatorId := "kkc"
	// appSecret := "buucl6eikmobttvpdaa0"
	c.JSON(200, gin.H{
		"currency": "RMB",
		"playerID": "kkc",
		"nickname": "mike",
		"test":     false,
		"time":     time.Now().Unix(),
	})
	fmt.Println("Validate")
	return
}

func WalletBalance(c *gin.Context) {

	c.JSON(200, gin.H{
		"balance":  "1000",
		"currency": "RMB",
		"time":     time.Now().Unix(),
	})
	fmt.Println("WalletBalance")
	return
}

func WalletDebit(c *gin.Context) {

	c.JSON(200, gin.H{
		"balance":  "1000",
		"currency": "RMB",
		"time":     time.Now().Unix(),
	})
	fmt.Println("WalletDebit")
	return
}

func WalletCredit(c *gin.Context) {

	c.JSON(200, gin.H{
		"balance":  "1000",
		"currency": "RMB",
		"time":     time.Now().Unix(),
	})
	fmt.Println("WalletCredit")
	return
}

func Rollback(c *gin.Context) {
	return
}
