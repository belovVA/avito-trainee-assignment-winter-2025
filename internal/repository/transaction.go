package repository

import (
	"avito-coin-service/internal/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *model.Transaction) error
	GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error)
	GetListSentTransactionByID(ID uint) ([]*model.Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: DB,
	}
}

func (r *transactionRepository) Create(tx *model.Transaction) error {
	return r.DB.Create(tx).Error
}

func (r *transactionRepository) GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := r.DB.Where("to_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}

func (r *transactionRepository) GetListSentTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction
	if err := r.DB.Where("from_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}
