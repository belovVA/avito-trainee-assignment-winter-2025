package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// JwtConfig The structure that allows to get the jwtSecretKey from the .env file
type JwtConfig struct {
	Key string `envconfig:"JWT_SECRET_KEY" required:"true"`
}

// LoadJwtConfig loading jwtKey from .env
func LoadJwtConfig() *JwtConfig {

	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal(err.Error())
	}

	var cfg JwtConfig
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf(".env:  %s", err.Error())
	}

	return &cfg
}
