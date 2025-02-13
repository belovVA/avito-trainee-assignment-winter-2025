package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	FromUser string `gorm:"not null"`
	ToUser   string `gorm:"not null"`
	Amount   int    `gorm:"default:0"`
}
