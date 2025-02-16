package handler

import (
	"net/http"

	"avito-coin-service/internal/service"

	"github.com/gin-gonic/gin"
)

// SendCoinRequest — the structure of the coin transfer request
type SendCoinRequest struct {
	ToUser string `json:"toUser" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

// TransactionHandler interface Handler
//
//	API "/api/sendCoin"
type TransactionHandler interface {
	SendCoinHandler(c *gin.Context)
}

type trxH struct {
	trxService service.TransactionService
}

func NewTransactionHandler(trxService service.TransactionService) TransactionHandler {
	return &trxH{trxService}
}

func (t *trxH) SendCoinHandler(c *gin.Context) {
	var req SendCoinRequest

	// Check body request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request body"})

		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Transaction amount must be more than 0"})

		return
	}

	fromUserName, exists := c.Get("userName")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})

		return
	}

	err := t.trxService.SendCoins(fromUserName.(string), req.ToUser, req.Amount)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successful"})
}
