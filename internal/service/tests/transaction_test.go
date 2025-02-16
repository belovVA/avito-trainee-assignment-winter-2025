package service_test

import (
	"errors"
	"testing"

	"avito-coin-service/internal/model"
	"avito-coin-service/internal/service"
	"avito-coin-service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendCoins(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockTxRepo := new(mocks.MockTransactionRepository)
	transactionService := service.NewTransactionService(mockUserRepo, mockTxRepo)

	fromUser := &model.User{
		ID:       1,
		Name:     "Alice",
		Balance:  1000,
		Password: "hashedpassword",
	}

	toUser := &model.User{
		ID:       2,
		Name:     "Bob",
		Balance:  500,
		Password: "hashedpassword",
	}

	tests := []struct {
		name          string
		fromUserName  string
		toUserName    string
		amount        int
		mockSetup     func()
		expectError   bool
		expectedError string
	}{
		{
			name:         "Success transaction",
			fromUserName: "Alice",
			toUserName:   "Bob",
			amount:       200,
			mockSetup: func() {

				mockUserRepo.On("GetByName", "Alice").
					Return(fromUser, nil)

				mockUserRepo.On("GetByName", "Bob").
					Return(toUser, nil)

				mockTxRepo.On("Create", mock.AnythingOfType("*model.Transaction")).
					Return(nil)

				mockTxRepo.On("ProcessTransaction", fromUser, toUser, 200).
					Return(nil)
			},
			expectError:   false,
			expectedError: "",
		},

		{
			name:         "sender not found",
			fromUserName: "Andrew",
			toUserName:   "Tim",
			amount:       200,
			mockSetup: func() {
				mockUserRepo.On("GetByName", "Andrew").
					Return(nil, errors.New("sender not found"))
			},
			expectError:   true,
			expectedError: "sender not found",
		},

		{
			name:         "recipient was not found",
			fromUserName: "test3",
			toUserName:   "test4",
			amount:       200,
			mockSetup: func() {

				mockUserRepo.On("GetByName", "test3").
					Return(fromUser, nil)

				mockUserRepo.On("GetByName", "test4").
					Return(nil, errors.New("recipient was not found"))
			},
			expectError:   true,
			expectedError: "recipient was not found",
		},

		{
			name:         "impossible to make transaction to yourself",
			fromUserName: "Alice",
			toUserName:   "Alice",
			amount:       200,
			mockSetup: func() {

				mockUserRepo.On("GetByName", "Alice").
					Return(fromUser, nil)

			},
			expectError:   true,
			expectedError: "impossible to make transaction to yourself",
		},

		{
			name:         "Insufficient funds",
			fromUserName: "Alice",
			toUserName:   "Bob",
			amount:       2000,
			mockSetup: func() {
				mockUserRepo.On("GetByName", "Alice").
					Return(fromUser, nil)
				mockUserRepo.On("GetByName", "Bob").
					Return(toUser, nil)
			},
			expectError:   true,
			expectedError: "Insufficient funds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := transactionService.SendCoins(tt.fromUserName, tt.toUserName, tt.amount)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockUserRepo.AssertExpectations(t)
			mockTxRepo.AssertExpectations(t)
		})
	}
}
