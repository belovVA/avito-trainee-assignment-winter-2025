package database

import (
	models "avito-coin-service/internal/model"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Transaction{})
}
