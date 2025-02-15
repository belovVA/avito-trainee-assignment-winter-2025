package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет JWT перед доступом к защищённым ресурсам
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Токен отсутствует"})
			c.Abort()
			return
		}

		// Ожидаем формат "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неверный формат токена"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Проверяем токен
		userName, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": "Неверный токен"})
			c.Abort()
			return
		}

		// Сохраняем имя пользователя в контексте запроса
		c.Set("userName", userName)

		c.Next() // Передаём запрос дальше
	}
}
