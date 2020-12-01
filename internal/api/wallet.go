package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/noworldwar/api-web/internal/entity"
	"github.com/noworldwar/api-web/internal/model"
	"github.com/noworldwar/api-web/internal/pkg/utils"
)

// @Func Validate
// @Request Bodys {token,string,true} @
// @Request Bodys {gpID,string,true} @
// @Request Bodys {appSecret,string,true} @
// @Response 200 {"playerID":string,"nickname":string,"currency":string,"test":bool,"time":bigint} @
// @Response 400  @
// @Response 401  @
// @Response 403  @
// @Response 404  @
// @Response 500 {"error":string} @
func Validate(c *gin.Context) {

	token := c.PostForm("token")
	// operatorID := c.PostForm("operatorID")
	// appSecret := c.PostForm("appSecret")

	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "token", "operatorID", "appSecret"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}

	player, err := entity.GetPlayer(token)
	if err != nil || player.PlayerID == "" {
		utils.ErrorResponse(c, 500, "Get Player Failed", err)
	} else {
		c.JSON(200, gin.H{
			"currency": "RMB",
			"playerID": player.PlayerID,
			"nickname": player.NickName,
			"test":     false,
			"time":     time.Now().Unix(),
		})
	}
	// operatorId := "kkc"
	// appSecret := "buucl6eikmobttvpdaa0"

	return
}

// @Func WalletBalance
// @Request Bodys {token,string,true} @
// @Request Bodys {operatorID,string,true} @
// @Request Bodys {appSecret,string,true} @
// @Request Bodys {playerID,string,true} @
// @Response 200 {"balance":bigint,"currency":string,"time":bigint} @
// @Response 400  @
// @Response 404  @
// @Response 500 {"error":string} @
func WalletBalance(c *gin.Context) {

	token := c.PostForm("token")
	// operatorID := c.PostForm("operatorID")
	// appSecret := c.PostForm("appSecret")
	// playerID := c.PostForm("playerID")
	// Step 1: Check the required parameters
	if missing := utils.CheckPostFormData(c, "token", "operatorID", "appSecret", "playerID"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}
	player, err := entity.GetPlayer(token)
	if err != nil {
		utils.ErrorResponse(c, 500, "Get Player Failed", err)
		return
	} else if player.PlayerID == "" {
		utils.ErrorResponse(c, 404, "Player Not Found", fmt.Errorf("%s", token))
		return
	} else {
		c.JSON(200, gin.H{
			"balance":  player.Balance,
			"currency": "RMB",
			"time":     time.Now().Unix(),
		})
	}

	return
}

// @Func WalletDebit
// @Request Bodys {token,string,true} @
// @Request Bodys {operatorID,string,true} @
// @Request Bodys {appSecret,string,true} @
// @Request Bodys {playerID,string,true} @
// @Request Bodys {amount,int,true} @
// @Response 200 {"balance":bigint,"currency":string,"time":bigint} @
// @Response 400  @
// @Response 401  @
// @Response 403  @
// @Response 404  @
// @Response 500 {"error":string} @
func WalletDebit(c *gin.Context) {

	calWallet(c, "debit")

	return
}

// @Func WalletDebit
// @Request Bodys {token,string,true} @
// @Request Bodys {operatorID,string,true} @
// @Request Bodys {appSecret,string,true} @
// @Request Bodys {playerID,string,true} @
// @Request Bodys {amount,int,true} @
// @Response 200 {"balance":bigint,"currency":string,"time":bigint} @
// @Response 400  @
// @Response 401  @
// @Response 403  @
// @Response 404  @
// @Response 500 {"error":string} @
func WalletCredit(c *gin.Context) {

	calWallet(c, "credit")

	return
}

func Rollback(c *gin.Context) {
	return
}

func calWallet(c *gin.Context, calType string) {

	if missing := utils.CheckPostFormData(c, "token", "operatorID", "appSecret", "playerID", "amount"); missing != "" {
		utils.ErrorResponse(c, 400, "Missing parameter: "+missing, nil)
		return
	}
	amount, err := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, 400, "Something went wrong with data type", nil)
		return
	} else if amount < 0 {
		utils.ErrorResponse(c, 400, "The amount must not be less than zero", nil)
		return
	}
	if calType == "debit" {
		amount = -amount
	}

	data, err := entity.GetPlayer(c.PostForm("token"))
	if err != nil {
		utils.ErrorResponse(c, 500, "Get Player Failed", err)
		return
	} else if data.PlayerID == "" {
		utils.ErrorResponse(c, 404, "Player Not Found", fmt.Errorf("%s", c.PostForm("token")))
		return
	} else {
		// UpdatePlayer model
		result := data.Balance + amount
		player := model.Player{
			PlayerID: data.PlayerID,
			Balance:  result,
		}

		affected, err := entity.UpdatePlayer(player)

		if err != nil {
			utils.ErrorResponse(c, 500, "Update Player Failed", err)
			return
		}
		fmt.Println("Updated ", affected, " row")
		c.JSON(200, gin.H{
			"balance":  result,
			"currency": "RMB",
			"time":     time.Now().Unix(),
		})
	}
	return

}
