package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) Create(tx *model.Transaction) error {
	return database.DB.Create(tx).Error
}

func (r *TransactionRepository) GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := database.DB.Where("to_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}

func (r *TransactionRepository) GetListSentTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := database.DB.Where("from_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}
