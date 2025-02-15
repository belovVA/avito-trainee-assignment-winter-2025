package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type JwtConfig struct {
	Key string `envconfig:"JWT_SECRET_KEY" required:"true"`
}

func LoadJwtConfig() *JwtConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println(".env file loaded successfully")
	}
	var cfg JwtConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("pizda %s", err.Error())
	}
	return &cfg
}
