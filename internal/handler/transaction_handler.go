package handler

import (
	"avito-coin-service/internal/repository"
	"avito-coin-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendCoinRequest — структура запроса на перевод монет
type SendCoinRequest struct {
	ToUser string `json:"toUser" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

func SendCoinHandler(c *gin.Context) {
	var req SendCoinRequest

	// Проверяем тело запроса
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Неверный формат запроса"})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Сумма перевода должна быть больше 0"})
		return
	}

	// Получаем имя отправителя из JWT
	fromUserName, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован"})
		return
	}

	userRepo := repository.NewUserRepository()
	txRepo := repository.NewTransactionRepository()
	txService := service.NewTransactionService(userRepo, txRepo)

	// Выполняем перевод монет
	err := txService.SendCoins(fromUserName.(string), req.ToUser, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Транзакция успешна"})
}
