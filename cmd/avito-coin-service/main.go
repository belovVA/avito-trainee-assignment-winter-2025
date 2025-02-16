package main

import (
	"log"

	"avito-coin-service/internal/app"
)

func main() {
	application := app.NewApp()

	log.Println("Starting server on :8080")
	application.Run(":8080")
}
