package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
)

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

// Create — создаёт запись о транзакции
func (r *TransactionRepository) Create(tx *model.Transaction) error {
	return database.DB.Create(tx).Error
}
