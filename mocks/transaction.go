package mocks

import (
	"avito-coin-service/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(tx *model.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockTransactionRepository) ProcessTransaction(fromUser *model.User, toUser *model.User, amount int) error {
	args := m.Called(fromUser, toUser, amount)

	transaction := model.Transaction{
		FromUser: fromUser.ID,
		ToUser:   toUser.ID,
		Amount:   amount,
	}

	_ = m.Create(&transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error) {
	args := m.Called(ID)

	return args.Get(0).([]*model.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetListSentTransactionByID(ID uint) ([]*model.Transaction, error) {
	args := m.Called(ID)

	return args.Get(0).([]*model.Transaction), args.Error(1)
}
