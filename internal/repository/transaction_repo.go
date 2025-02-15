package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

type ITransactionRepository interface {
	Create(tx *model.Transaction) error
	GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error)
	GetListSentTransactionByID(ID uint) ([]*model.Transaction, error)
}

type transactionRepository struct{}

func NewTransactionRepository() ITransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(tx *model.Transaction) error {
	return database.DB.Create(tx).Error
}

func (r *transactionRepository) GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := database.DB.Where("to_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}

func (r *transactionRepository) GetListSentTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := database.DB.Where("from_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}
