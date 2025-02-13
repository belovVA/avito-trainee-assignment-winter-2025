package main

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/handler"
	"avito-coin-service/internal/repository"
	"avito-coin-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.Migrate()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.POST("/api/auth", userHandler.AuthHandler)

	r.Run(":8080")
}
