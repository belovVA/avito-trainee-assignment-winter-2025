package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Balance  int    `gorm:"default:1000"`
}

// package models

// package models

// type User struct {
//     ID       uint   `gorm:"primaryKey"` // Это поле будет автоинкрементным и использоваться как уникальный идентификатор
//     Name     string `gorm:"unique;not null"`
//     Password string `gorm:"not null"`
//     Balance  int    `gorm:"default:1000"` // Начальный баланс
// }
