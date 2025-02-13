package model

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Balance  int    `gorm:"default:1000"`
}
