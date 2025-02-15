package service

import (
	"avito-coin-service/internal/middleware"
	models "avito-coin-service/internal/model"
	"avito-coin-service/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Authenticate(name, password string) (string, error)
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(
	userRepo repository.IUserRepository,
) *UserService {

	return &UserService{userRepo}
}

func (s *UserService) Authenticate(name, password string) (string, error) {
	user, err := s.userRepo.GetByName(name)

	if err != nil {

		if hashPass, err := hashPassword(password); err != nil {
			return "", err

		} else {
			user = &models.User{
				Name:     name,
				Password: hashPass,
				Balance:  1000,
			}
		}

		if err := s.userRepo.Create(user); err != nil {
			return "", err
		}

		if token, err := middleware.CreateToken(user.Name); err != nil {
			return "", err

		} else {
			return token, nil
		}
	}

	if !comparePasswords(user.Password, password) {
		return "", fmt.Errorf("неверный пароль")
	}

	if token, err := middleware.CreateToken(user.Name); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		return "", err
	}

	return string(hash), nil
}

// ComparePasswords сравнивает хэшированный пароль с введенным
func comparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	return err == nil
}
