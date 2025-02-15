package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

// type MerchRepository

type MerchRepository struct{}

func NewMerchRepository() *MerchRepository {
	return &MerchRepository{}
}

func (r *MerchRepository) Create(tx *model.Merch) error {
	return database.DB.Create(tx).Error
}

func (r *MerchRepository) GetByID(id uint) (*model.Merch, error) {
	var merch model.Merch

	if err := database.DB.Where("id = ?", id).First(&merch).Error; err != nil {

		return nil, err
	}

	return &merch, nil
}

func (r *MerchRepository) GetByName(name string) (*model.Merch, error) {
	var merch model.Merch

	if err := database.DB.Where("name = ?", name).First(&merch).Error; err != nil {

		return nil, err
	}

	return &merch, nil
}
