package database

import (
	"fmt"
	"log"

	"avito-coin-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PG struct {
	DBptr *gorm.DB
}

func InitDB(cfg *config.PGConfig) *PG {
	// Используем данные из конфигурации
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
	)

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect DB: %s", err)
	}

	log.Println("DB Connected")

	return &PG{
		DBptr: db,
	}

}
