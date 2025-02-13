package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Balance  int    `gorm:"default:1000"`
}
