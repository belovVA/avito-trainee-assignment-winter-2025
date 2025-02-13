package repository

import (
	"avito-coin-service/internal/database"
	models "avito-coin-service/internal/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {

	return &UserRepository{}
}

func (r *UserRepository) GetByName(name string) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("name = ?", name).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {

	return database.DB.Create(user).Error
}
