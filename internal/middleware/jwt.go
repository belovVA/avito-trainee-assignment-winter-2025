package middleware

import (
	"fmt"
	"time"

	"avito-coin-service/config"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userName string) (string, error) {

	cfg := config.LoadJwtConfig()

	// Creating claims
	claims := jwt.MapClaims{
		"sub": userName,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.Key))
}

func ValidateToken(tokenString string) (string, error) {
	cfg := config.LoadJwtConfig()

	// Check and decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unsupported signature method: %v", token.Header["alg"])
		}

		return []byte(cfg.Key), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("Invalid token: %v", err)
	}

	// Извлекаем информацию из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	}
	return "", fmt.Errorf("invalid token data")
}
