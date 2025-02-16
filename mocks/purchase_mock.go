package mocks

import (
	"avito-coin-service/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockPurchaseRepository struct {
	mock.Mock
}

func (m *MockPurchaseRepository) Create(tx *model.Purchase) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockPurchaseRepository) GetByUserAndMerch(userID uint, merchId uint) (*model.Purchase, error) {
	args := m.Called(userID, merchId)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Purchase), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPurchaseRepository) GetListByUserID(userID uint) ([]*model.Purchase, error) {
	args := m.Called(userID)
	return args.Get(0).([]*model.Purchase), args.Error(1)
}

func (m *MockPurchaseRepository) Update(p model.Purchase) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPurchaseRepository) ProcessPurchase(user *model.User, merch *model.Merch) error {
	args := m.Called(user, merch)
	return args.Error(0)
}
