package service

import (
	"fmt"

	"avito-coin-service/internal/repository"
)

type TransactionService interface {
	SendCoins(fromUserName string, toUserName string, amount int) error
}

type trx struct {
	userRepo repository.UserRepository
	txRepo   repository.TransactionRepository
}

func NewTransactionService(
	userRepo repository.UserRepository,
	txRepo repository.TransactionRepository,
) *trx {
	return &trx{
		userRepo: userRepo,
		txRepo:   txRepo,
	}
}

// SendCoins — переводит монеты от одного пользователя другому
func (s *trx) SendCoins(fromUserName string, toUserName string, amount int) error {

	fromUser, err := s.userRepo.GetByName(fromUserName)
	if err != nil {
		return fmt.Errorf("sender not found")
	}

	toUser, err := s.userRepo.GetByName(toUserName)
	if err != nil {
		return fmt.Errorf("recipient was not found")
	}

	if toUser.ID == fromUser.ID {
		return fmt.Errorf("impossible to make transaction to yourself")
	}

	if fromUser.Balance < amount {
		return fmt.Errorf("Insufficient funds")
	}

	return s.txRepo.ProcessTransaction(fromUser, toUser, amount)
}
