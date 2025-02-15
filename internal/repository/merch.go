package repository

import (
	"avito-coin-service/internal/model"

	"gorm.io/gorm"
)

type MerchRepository interface {
	Create(tx *model.Merch) error
	GetByID(id uint) (*model.Merch, error)
	GetByName(name string) (*model.Merch, error)
}

type merchRep struct {
	DB *gorm.DB
}

func NewMerchRepository(DB *gorm.DB) MerchRepository {
	return &merchRep{
		DB: DB,
	}
}

func (r *merchRep) Create(tx *model.Merch) error {
	return r.DB.Create(tx).Error
}

func (r *merchRep) GetByID(id uint) (*model.Merch, error) {
	var merch model.Merch

	if err := r.DB.Where("id = ?", id).First(&merch).Error; err != nil {

		return nil, err
	}

	return &merch, nil
}

func (r *merchRep) GetByName(name string) (*model.Merch, error) {
	var merch model.Merch

	if err := r.DB.Where("name = ?", name).First(&merch).Error; err != nil {

		return nil, err
	}

	return &merch, nil
}
