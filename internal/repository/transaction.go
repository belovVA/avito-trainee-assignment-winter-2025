package repository

import (
	"avito-coin-service/internal/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *model.Transaction) error
	GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error)
	GetListSentTransactionByID(ID uint) ([]*model.Transaction, error)
	ProcessTransaction(fromUser *model.User, toUser *model.User, amount int) error
}

type transactionRep struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &transactionRep{
		DB: DB,
	}
}

func (r *transactionRep) Create(tx *model.Transaction) error {
	return r.DB.Create(tx).Error
}

func (r *transactionRep) GetListRecievedTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction

	if err := r.DB.Where("to_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}

	return trxs, nil
}

func (r *transactionRep) GetListSentTransactionByID(ID uint) ([]*model.Transaction, error) {
	var trxs []*model.Transaction

	if err := r.DB.Where("from_user = ?", ID).Find(&trxs).Error; err != nil {
		return nil, err
	}

	return trxs, nil
}

func (r *transactionRep) ProcessTransaction(fromUser *model.User, toUser *model.User, amount int) error {
	tx := r.DB.Begin()

	fromUser.Balance -= amount
	toUser.Balance += amount

	if err := tx.Save(fromUser).Error; err != nil {
		tx.Rollback()

		return err
	}

	if err := tx.Save(toUser).Error; err != nil {
		tx.Rollback()

		return err
	}

	transaction := model.Transaction{
		FromUser: fromUser.ID,
		ToUser:   toUser.ID,
		Amount:   amount,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}
