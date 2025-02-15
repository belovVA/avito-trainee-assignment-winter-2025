package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

type IMerchRepository interface {
	Create(tx *model.Merch) error
	GetByID(id uint) (*model.Merch, error)
	GetByName(name string) (*model.Merch, error)
}

type merchRepository struct{}

func NewMerchRepository() IMerchRepository {
	return &merchRepository{}
}

func (r *merchRepository) Create(tx *model.Merch) error {
	return database.DB.Create(tx).Error
}

func (r *merchRepository) GetByID(id uint) (*model.Merch, error) {
	var merch model.Merch

	if err := database.DB.Where("id = ?", id).First(&merch).Error; err != nil {

		return nil, err
	}

	return &merch, nil
}

func (r *merchRepository) GetByName(name string) (*model.Merch, error) {
	var merch model.Merch

	if err := database.DB.Where("name = ?", name).First(&merch).Error; err != nil {

		return nil, err
	}

	return &merch, nil
}
