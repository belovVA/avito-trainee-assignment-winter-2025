package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// JwtConfig The structure that allows to get the jwtSecretKey from the .env file
type JwtConfig struct {
	Key string `envconfig:"JWT_SECRET_KEY" required:"true"`
}

// LoadJwtConfig loading jwtKey from .env
func LoadJwtConfig() *JwtConfig {
	path := ".env"
	if os.Getenv("GO_ENV") == "test" {
		path = "../../../.env"
	}
	if err := godotenv.Load(path); err != nil {
		log.Fatal(err.Error())
	}

	var cfg JwtConfig
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf(".env:  %s", err.Error())
	}

	return &cfg
}
