package middleware

import (
	// "avito-coin-service/config"
	"avito-coin-service/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtKey struct {
	secretKey []byte
}

func CreateToken(userName string) (string, error) {
	cfg := config.LoadJwtConfig()

	// Создаем claims
	claims := jwt.MapClaims{
		"sub": userName,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}

	// Создаем новый токен с указанием алгоритма и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.Key))
}

func ValidateToken(tokenString string) (string, error) {
	cfg := config.LoadJwtConfig()

	// Проверяем и декодируем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неподдерживаемый метод подписи: %v", token.Header["alg"])
		}
		return []byte(cfg.Key), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("неверный токен: %v", err)
	}

	// Извлекаем информацию из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	}
	return "", fmt.Errorf("неверные данные токена")
}
