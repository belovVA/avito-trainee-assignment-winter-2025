package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks JWT before accessing protected resources
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Токен отсутствует"})
			c.Abort()

			return
		}

		// Expecting format "Bearer <token>"
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {

			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token format"})
			c.Abort()

			return
		}

		tokenString := parts[1]

		// Check token
		userName, err := ValidateToken(tokenString)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Invalid token"})
			c.Abort()

			return
		}

		// Save username in Cotnext request
		c.Set("userName", userName)

		c.Next() // Passing the request
	}
}
