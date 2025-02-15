package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type PGConfig struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     string `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
}

func LoadPGConfig() *PGConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println(".env file loaded successfully")
	}

	var config PGConfig
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &config
}

// func LoadPGConfig() *PGConfig {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	config := &PGConfig{
// 		Host:     getEnv("DATABASE_HOST", "localhost"),
// 		User:     getEnv("DATABASE_USER", "postgres"),
// 		Password: getEnv("DATABASE_PASSWORD", "postgres"),
// 		Name:     getEnv("DATABASE_NAME", "avito"),
// 		Port:     getEnv("DATABASE_PORT", "5432"),
// 		// JWTSecret:  getEnv("JWT_SECRET", "mySecretKey"),
// 	}

// 	// middleware.InitsecretKey(config.JWTSecret)

// 	return config
// }

// func getEnv(key, defaultValue string) string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists {
// 		return defaultValue
// 	}
// 	return value
// }
