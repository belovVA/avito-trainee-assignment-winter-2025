package main

import (
	"avito-coin-service/config"
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/handler"
	"avito-coin-service/internal/middleware"
	"avito-coin-service/internal/repository"
	"avito-coin-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadPGConfig()

	log.Println("Config has been uploaded")

	// Инициализируем подключение к базе данных с конфигурацией
	pg := database.InitDB(cfg)

	// Инициализация реопзиториев
	userRepo := repository.NewUserRepository(pg.DBptr)
	trxRepo := repository.NewTransactionRepository(pg.DBptr)
	purchRepo := repository.NewPurchaseRepository(pg.DBptr)
	merchRepo := repository.NewMerchRepository(pg.DBptr)

	// service
	usrService := service.NewUserService(userRepo)
	trxService := service.NewTransactionService(userRepo, trxRepo)
	purchService := service.NewPurchaseService(userRepo, merchRepo, purchRepo)
	infoService := service.NewInfoService(userRepo, trxRepo, merchRepo, purchRepo)

	// handlers
	a := handler.NewUserHandler(usrService)
	s := handler.NewTransactionHandler(trxService)
	p := handler.NewPurchHandler(purchService)
	i := handler.NewInfoHandler(infoService)
	// database.Migrate()

	r := gin.Default()

	apiAuth := r.Group("/api")

	apiAuth.POST("/auth", a.AuthHandler)

	apiAuth.Use(middleware.AuthMiddleware())
	{
		apiAuth.POST("/sendCoin", s.SendCoinHandler)
		apiAuth.GET("/buy/:item", p.PurchaseHandler)
		apiAuth.GET("/info", i.InfoHandler)
	}

	r.Run(":8080")
}
