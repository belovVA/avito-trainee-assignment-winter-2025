package mocks

import (
	"avito-coin-service/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockMerchRepository struct {
	mock.Mock
}

func (m *MockMerchRepository) Create(tx *model.Merch) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockMerchRepository) GetByID(id uint) (*model.Merch, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Merch), args.Error(1)
}

func (m *MockMerchRepository) GetByName(name string) (*model.Merch, error) {
	args := m.Called(name)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Merch), args.Error(1)
	}
	return nil, args.Error(1)
}
