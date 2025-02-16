package tests

import (
	"avito-coin-service/internal/model"
	"avito-coin-service/internal/service"
	"avito-coin-service/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuyMerch(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockMerchRepo := new(mocks.MockMerchRepository)
	mockPurchaseRepo := new(mocks.MockPurchaseRepository)

	purchaseService := service.NewPurchaseService(mockUserRepo, mockMerchRepo, mockPurchaseRepo)

	// Создадим пользователей для теста

	tests := []struct {
		name        string
		Name        string
		merchName   string
		mockSetup   func()
		expectError bool
	}{
		{
			name:      "Пользователь не найден",
			Name:      "NotFound",
			merchName: "T-shirt",
			mockSetup: func() {
				mockUserRepo.On("GetByName", "NotFound").Return(nil, errors.New("not found"))

			},
			expectError: true,
		},
		{
			name:      "Мерч не найден",
			Name:      "Alice",
			merchName: "NoMerch",
			mockSetup: func() {
				mockUserRepo.On("GetByName", "Alice").Return(&model.User{Name: "Alice", Balance: 300}, nil)
				mockMerchRepo.On("GetByName", "NoMerch").Return(nil, errors.New("not found"))
				// Добавьте выводы для отладки
			},
			expectError: true,
		},
		{
			name:      "Недостаточно средств",
			Name:      "littleMoney",
			merchName: "T-shirt",
			mockSetup: func() {
				mockUserRepo.On("GetByName", "littleMoney").Return(&model.User{ID: 3, Name: "littleMoney", Password: "11", Balance: 100}, nil)
				mockMerchRepo.On("GetByName", "T-shirt").Return(&model.Merch{ID: 1, Name: "T-shirt", Price: 200}, nil)

			},
			expectError: true,
		},
		{
			name:      "Успешная покупка нового товара",
			Name:      "goodMoney",
			merchName: "T-shirt",
			mockSetup: func() {
				user := &model.User{Name: "goodMoney", Balance: 500}
				merch := &model.Merch{ID: 1, Name: "T-shirt", Price: 200}

				mockUserRepo.On("GetByName", "goodMoney").Return(user, nil)
				mockPurchaseRepo.On("Create", mock.AnythingOfType("*model.Purchase")).Return(nil)
				mockMerchRepo.On("GetByName", "T-shirt").Return(&model.Merch{ID: 1, Name: "T-shirt", Price: 200}, nil)
				mockPurchaseRepo.On("GetByUserAndMerch", uint(0), uint(1)).Return(nil, errors.New("not found"))
				mockPurchaseRepo.On("ProcessPurchase", user, merch).Return(nil)

			},
			expectError: false,
		},
		{
			name:      "Успешная покупка товара который уже был",
			Name:      "And",
			merchName: "T-shirt",
			mockSetup: func() {
				user := &model.User{ID: 7, Name: "And", Balance: 500}
				merch := &model.Merch{ID: 1, Name: "T-shirt", Price: 200}

				mockUserRepo.On("GetByName", "And").Return(user, nil)
				mockPurchaseRepo.On("Create", mock.AnythingOfType("*model.Purchase")).Return(nil)
				mockMerchRepo.On("GetByName", "T-shirt").Return(&model.Merch{ID: 1, Name: "T-shirt", Price: 200}, nil)
				mockPurchaseRepo.On("GetByUserAndMerch", user.ID, merch.ID).Return(&model.Purchase{UserID: user.ID, MerchID: merch.ID, Count: 1}, nil)
				mockPurchaseRepo.On("ProcessPurchase", user, merch).Return(nil)
				mockPurchaseRepo.On("Update", mock.AnythingOfType("*model.Purchase")).Return(nil)

			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := purchaseService.BuyMerch(tt.Name, tt.merchName)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// mockUserRepo.AssertExpectations(t)
			// mockMerchRepo.AssertExpectations(t)
			// mockPurchaseRepo.AssertExpectations(t)

		})
	}
}
