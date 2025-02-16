package tests

import (
	"avito-coin-service/internal/model"
	"avito-coin-service/internal/service"
	"avito-coin-service/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoService_GetInfo(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockTxRepo := new(mocks.MockTransactionRepository)
	mockMerchRepo := new(mocks.MockMerchRepository)
	mockPurchaseRepo := new(mocks.MockPurchaseRepository)

	service := service.NewInfoService(mockUserRepo, mockTxRepo, mockMerchRepo, mockPurchaseRepo)

	tests := []struct {
		name           string
		userName       string
		mockSetup      func()
		expectedError  bool
		expectedResult *model.InfoResponse
	}{
		{
			name:     "Успешное получение информации",
			userName: "Sandy",
			mockSetup: func() {
				mockUserRepo.On("GetByName", "Sandy").Return(&model.User{ID: 1, Name: "Sandy", Balance: 500}, nil)
				mockPurchaseRepo.On("GetListByUserID", uint(1)).Return([]*model.Purchase{
					{MerchID: 1, Count: 2},
				}, nil)
				mockMerchRepo.On("GetByID", uint(1)).Return(&model.Merch{ID: 2, Name: "Pen", Price: 10}, nil)
				mockUserRepo.On("GetByID", uint(2)).Return(&model.User{ID: 2, Name: "Alice", Balance: 1000}, nil) // Настроить для ID 2

				mockTxRepo.On("GetListRecievedTransactionByID", uint(1)).Return([]*model.Transaction{
					{FromUser: 2, Amount: 100},
				}, nil)
				mockTxRepo.On("GetListSentTransactionByID", uint(1)).Return([]*model.Transaction{
					{ToUser: 2, Amount: 50},
				}, nil)
			},
			expectedError: false,
			expectedResult: &model.InfoResponse{
				Coins: 500,
				Inventory: []model.InventoryItem{
					{Type: "Pen", Quantity: 2},
				},
				CoinHistory: model.CoinHistory{
					Received: []model.CoinTransaction{
						{FromUser: "Alice", Amount: 100},
					},
					Sent: []model.CoinTransaction{
						{ToUser: "Alice", Amount: 50},
					},
				},
			},
		},
		{
			name:     "Ошибка, пользователь не найден",
			userName: "Alice",
			mockSetup: func() {
				mockUserRepo.On("GetByName", "Alice").Return(nil, fmt.Errorf("пользователь не найден"))
			},
			expectedError:  true,
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup() // вызываем настройку моков

			info, err := service.GetInfo(tt.userName)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, info)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, info)
			}
		})
	}
}
