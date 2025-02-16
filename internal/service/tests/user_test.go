package service_test

import (
	"errors"
	"testing"

	"avito-coin-service/internal/model"
	"avito-coin-service/internal/service"
	"avito-coin-service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockUserRepo)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	tests := []struct {
		name        string
		inputName   string
		inputPass   string
		mockUser    *model.User
		mockError   error
		expectError bool
	}{
		{
			name:      "user has been found, and the password is correct",
			inputName: "testuser",
			inputPass: "password",
			mockUser: &model.User{
				Name:     "testuser",
				Password: string(hashedPassword),
				Balance:  1000,
			},
			mockError:   nil,
			expectError: false,
		},

		{
			name:        "user not found, create new",
			inputName:   "newuser",
			inputPass:   "newpassword",
			mockUser:    nil,
			mockError:   errors.New("user not found"),
			expectError: false,
		},

		{
			name:      "The user has been found, the password is incorrect",
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
