package handler

import (
	"avito-coin-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	AuthHandler(c *gin.Context)
}

type userH struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userH{userService}
}

func (u *userH) AuthHandler(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "неверный формат запроса"})
		return
	}

	token, err := u.userService.Authenticate(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
