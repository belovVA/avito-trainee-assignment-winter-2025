package handler

import (
	"avito-coin-service/internal/repository"
	"avito-coin-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	Merch string
}

func PurchaseHandler(c *gin.Context) {
	// item := c.Param("item") // Получаем {item} из URL
	// c.JSON(http.StatusOK, gin.H{
	// 	// "message": fmt.Sprintf("Покупка товара: %s", item),
	// })

	itemName := c.Param("item")

	// Проверяем, передан ли параметр
	if itemName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не указан предмет"})
		return
	}
	// Получаем имя отправителя из JWT
	user, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизован"})
		return
	}

	userRepo := repository.NewUserRepository()
	merchRepo := repository.NewMerchRepository()
	purchaseRepo := repository.NewPurchaseRepository()

	purchService := service.NewPurchaseService(userRepo, merchRepo, purchaseRepo)

	if err := purchService.BuyMerch(user.(string), itemName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Куплен товар"})
	return
}
