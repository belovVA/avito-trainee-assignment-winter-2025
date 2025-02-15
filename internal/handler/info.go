package handler

import (
	"avito-coin-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InfoH interface {
	InfoHandler(c *gin.Context)
}

type infohdl struct {
	infoService service.InfoService
}

func NewInfoHandler(infsrv service.InfoService) InfoH {
	return &infohdl{infsrv}
}

func (h *infohdl) InfoHandler(c *gin.Context) {
	// Получаем имя отправителя из JWT
	user, exists := c.Get("userName")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неавторизован"})
		return
	}

	if answ, err := h.infoService.GetInfo(user.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, answ)
	}

}
