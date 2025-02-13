package model

type Transaction struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	FromUser uint `gorm:"not null"`
	ToUser   uint `gorm:"not null"`
	Amount   int  `gorm:"default:0"`
}
