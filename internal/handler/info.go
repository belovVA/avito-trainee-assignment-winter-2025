package handler

import (
	"net/http"

	"avito-coin-service/internal/service"

	"github.com/gin-gonic/gin"
)

// InfoHandler interface Handler
//
//	API "/api/info"
type InfoHandler interface {
	InfoHandler(c *gin.Context)
}

type infoH struct {
	infoService service.InfoService
}

func NewInfoHandler(infsrv service.InfoService) InfoHandler {
	return &infoH{infsrv}
}

// infoHandl processes a request for user information
func (h *infoH) InfoHandler(c *gin.Context) {
	// Get userName from JWT Token
	user, exists := c.Get("userName")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized"})

		return
	}

	if answ, err := h.infoService.GetInfo(user.(string)); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return

	} else {

		c.JSON(http.StatusOK, answ)
	}

}
