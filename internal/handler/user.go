package handler

import (
	"net/http"

	"avito-coin-service/internal/service"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserHandler interface Handler
//
//	API "/api/auth"
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

	req := userResponse{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid request body"})

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
