package service_test

import (
	"avito-coin-service/internal/model"
	"avito-coin-service/internal/service"

	"avito-coin-service/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockUserRepo)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	// Таблица тестов
	tests := []struct {
		name        string
		inputName   string
		inputPass   string
		mockUser    *model.User
		mockError   error
		expectError bool
	}{
		{

			name:      "Пользователь найден, правильный пароль",
			inputName: "testuser",
			inputPass: "password",
			mockUser: &model.User{
				Name:     "testuser",
				Password: string(hashedPassword), // Хэш "password"
				Balance:  1000,
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "Пользователь не найден, создаём нового",
			inputName:   "newuser",
			inputPass:   "newpassword",
			mockUser:    nil,
			mockError:   errors.New("пользователь не найден"),
			expectError: false,
		},
		{
			name:      "Пользователь найден, неправильный пароль",
			inputName: "testuser",
			inputPass: "wrongpassword",
			mockUser: &model.User{
				Name:     "testuser",
				Password: string(hashedPassword),
				Balance:  1000,
			},
			mockError:   nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo.On("GetByName", tt.inputName).Return(tt.mockUser, tt.mockError)

			if tt.mockUser == nil {
				mockUserRepo.On("Create", mock.Anything).Return(nil)
			}

			token, err := userService.Authenticate(tt.inputName, tt.inputPass)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}

			mockUserRepo.AssertExpectations(t)
		})
	}
}
