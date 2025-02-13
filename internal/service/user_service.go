package service

import (
	models "avito-coin-service/internal/model"
	"avito-coin-service/internal/repository"
	"errors"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) Authenticate(name, password string) (*models.User, error) {
	user, err := s.userRepo.GetByName(name)
	if err != nil {
		// Если пользователя нет — создаем
		user = &models.User{
			Name:     name,
			Password: password, // ⚠️ Надо бы хешировать, но пока так
			Balance:  1000,
		}
		if err := s.userRepo.Create(user); err != nil {
			return nil, err
		}
		return user, nil
	}

	// Проверка пароля (позже заменим на хеш)
	if user.Password != password {
		return nil, errors.New("неверный пароль")
	}

	return user, nil
}
