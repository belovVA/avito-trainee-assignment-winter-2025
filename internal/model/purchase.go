package model

type Purchase struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	UserID  uint `gorm:"not null"`
	MerchID uint `gorm:"not null"`
	Count   uint `gorm:"default:1"`
}
