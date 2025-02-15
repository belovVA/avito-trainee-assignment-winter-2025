package handler

import (
	"avito-coin-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	Merch string
}

type purchaseH interface {
	PurchaseHandler(c *gin.Context)
}

type purchaseHandler struct {
	purchService service.PurchaseService
}

func NewPurchHandler(p service.PurchaseService) purchaseH {
	return &purchaseHandler{
		purchService: p,
	}
}
func (h *purchaseHandler) PurchaseHandler(c *gin.Context) {
	// item := c.Param("item") // Получаем {item} из URL
	// c.JSON(http.StatusOK, gin.H{
	// 	// "message": fmt.Sprintf("Покупка товара: %s", item),
	// })

	itemName := c.Param("item")

	// Проверяем, передан ли параметр
	if itemName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "не указан предмет"})
		return
	}
	// Получаем имя отправителя из JWT
	user, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован"})
		return
	}

	// purchService := service.NewPurchaseService(userRepo, merchRepo, purchaseRepo)

	if err := h.purchService.BuyMerch(user.(string), itemName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Куплен товар"})
	return
}
