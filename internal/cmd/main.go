package main

import (
	log "github.com/sirupsen/logrus"

	"avito-coin-service/config"
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/handler"
	"avito-coin-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Errorf("Error loading config: %s. Default values uploaded", err)
	} else {
		log.Info("Config has been uploaded")
	}

	// Инициализируем подключение к базе данных с конфигурацией
	database.InitDB(cfg)
	// database.Migrate()

	r := gin.Default()

	auth := r.Group("/api")

	auth.POST("/auth", handler.AuthHandler)

	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/sendCoin", handler.SendCoinHandler)
		auth.GET("/buy/:item", handler.PurchaseHandler)
		auth.GET("/info", handler.InfoHandler)
	}

	r.Run(":8080")
}
