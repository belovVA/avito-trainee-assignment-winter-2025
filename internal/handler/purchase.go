package handler

import (
	"net/http"

	"avito-coin-service/internal/service"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	Merch string
}

// PurchaseHandler interface Handler
//
//	API "/api/buy/:item"
type PurchaseHandler interface {
	PurchaseHandler(c *gin.Context)
}

type purchaseH struct {
	purchService service.PurchaseService
}

func NewPurchHandler(p service.PurchaseService) PurchaseHandler {

	return &purchaseH{
		purchService: p,
	}
}

func (h *purchaseH) PurchaseHandler(c *gin.Context) {
	itemName := c.Param("item")

	if itemName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "item didnt get "})

		return
	}

	user, exists := c.Get("userName")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})

		return
	}

	if err := h.purchService.BuyMerch(user.(string), itemName); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "purchase successful completed"})
}
