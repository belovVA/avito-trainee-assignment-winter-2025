package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

type MertchRepository struct{}

func NewMertchRepository() *MertchRepository {
	return &MertchRepository{}
}

func (r *MertchRepository) Create(tx *model.Mertch) error {
	return database.DB.Create(tx).Error
}

func (r *MertchRepository) GetByName(name string) (*model.Mertch, error) {
	var mertch model.Mertch

	if err := database.DB.Where("name = ?", name).First(&mertch).Error; err != nil {

		return nil, err
	}

	return &mertch, nil
}
