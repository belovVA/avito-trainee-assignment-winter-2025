package main

import (
	"avito-coin-service/internal/app"
	"log"
)

func main() {
	log.Println("Starting server on :8080")

	application := app.NewApp()
	application.Run(":8080")
}
