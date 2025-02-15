package config

import (
	"avito-coin-service/internal/middleware"
	"os"

	"github.com/joho/godotenv"
)

// Config хранит все конфигурационные параметры
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	JWTSecret  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "avito"),
		DBPort:     getEnv("DB_PORT", "5432"),
		JWTSecret:  getEnv("JWT_SECRET", "mySecretKey"),
	}

	middleware.InitsecretKey(config.JWTSecret)

	return config, err
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
