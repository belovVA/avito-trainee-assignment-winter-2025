package handler

import (
	"avito-coin-service/internal/repository"
	"avito-coin-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoHandler(c *gin.Context) {
	// Получаем имя отправителя из JWT
	user, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован"})
		return
	}

	userRepo := repository.NewUserRepository()
	trxRepo := repository.NewTransactionRepository()
	merchRepo := repository.NewMerchRepository()
	purchaseRepo := repository.NewPurchaseRepository()

	infoService := service.NewInfoService(userRepo, trxRepo, merchRepo, purchaseRepo)

	if answ, err := infoService.GetInfo(user.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, answ)
	}

}
