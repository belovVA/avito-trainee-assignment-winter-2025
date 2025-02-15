package repository

import (
	"avito-coin-service/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(ID uint) (*model.User, error)
	GetByName(name string) (*model.User, error)
}

type userRep struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRep {

	return &userRep{
		DB: DB,
	}
}

func (r *userRep) Create(user *model.User) error {

	return r.DB.Create(user).Error
}

func (r *userRep) GetByID(ID uint) (*model.User, error) {
	var user model.User

	if err := r.DB.Where("id = ?", ID).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *userRep) GetByName(name string) (*model.User, error) {
	var user model.User

	if err := r.DB.Where("name = ?", name).First(&user).Error; err != nil {

		return nil, err
	}

	return &user, nil
}
