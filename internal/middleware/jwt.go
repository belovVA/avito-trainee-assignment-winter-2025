package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("mySecretKey")

func CreateToken(userName string) (string, error) {
	// Создаем claims
	claims := jwt.MapClaims{
		"sub": userName,                              // Идентификатор пользователя
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Время истечения токена
	}

	// Создаем новый токен с указанием алгоритма и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен и возвращаем строку
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (string, error) {
	// Проверяем и декодируем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неподдерживаемый метод подписи: %v", token.Header["alg"])
		}
		return secretKey, nil
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
