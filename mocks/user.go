package mocks

import (
	"avito-coin-service/internal/model"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

// GetByName - мок метода получения пользователя по имени
func (m *MockUserRepository) GetByName(name string) (*model.User, error) {
	args := m.Called(name)
	err := args.Error(1)
	if err != nil {
		return nil, fmt.Errorf("user not found")

	}
	return args.Get(0).(*model.User), nil // Прямо извлекаем значение
}

// GetByID - мок метода получения пользователя по ID
func (m *MockUserRepository) GetByID(ID uint) (*model.User, error) {
	args := m.Called(ID)
	return args.Get(0).(*model.User), args.Error(1)
}

// Create - мок метода создания пользователя
func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}
