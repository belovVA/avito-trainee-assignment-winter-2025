package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

type IUserRepository interface {
	Create(user *model.User) error
	GetByID(ID uint) (*model.User, error)
	GetByName(name string) (*model.User, error)
}

type UserRepository struct{}

func NewUserRepository() *UserRepository {

	return &UserRepository{}
}

func (r *UserRepository) Create(user *model.User) error {

	return database.DB.Create(user).Error
}

func (r *UserRepository) GetByID(ID uint) (*model.User, error) {
	var user model.User

	if err := database.DB.Where("id = ?", ID).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByName(name string) (*model.User, error) {
	var user model.User

	if err := database.DB.Where("name = ?", name).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}
