package model

type Merch struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"unique;not null"`
	Price int    `gorm:"default:1000"`
}

func (Merch) TableName() string {
	return "merch"
}
