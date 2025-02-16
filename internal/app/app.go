package app

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

// App present the main application
type App struct {
	Router *gin.Engine
}

// NewApp creates pointer on new up
func NewApp() *App {
	cfg := config.LoadPGConfig()

	log.Println("Config has been uploaded")

	pg := database.InitDB(cfg)

	// repository
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

	r := gin.Default()

	apiAuth := r.Group("/api")

	apiAuth.POST("/auth", a.AuthHandler)

	// Requesting authorization
	apiAuth.Use(middleware.AuthMiddleware())
	{
		apiAuth.POST("/sendCoin", s.SendCoinHandler)
		apiAuth.GET("/buy/:item", p.PurchaseHandler)
		apiAuth.GET("/info", i.InfoHandler)
	}

	return &App{
		Router: r,
	}
}

// Run starts the server
func (a *App) Run(address string) {
	log.Printf("Server is running on %s", address)
	if err := a.Router.Run(address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
