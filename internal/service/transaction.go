package service

import (
	"avito-coin-service/internal/repository"
	"fmt"
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
		return fmt.Errorf("отправитель не найден")
	}

	toUser, err := s.userRepo.GetByName(toUserName)
	if err != nil {
		return fmt.Errorf("получатель не найден")
	}

	if toUser.ID == fromUser.ID {
		return fmt.Errorf("невозможно осуществить перевод самому себе")
	}

	if fromUser.Balance < amount {
		return fmt.Errorf("недостаточно средств")
	}

	return s.txRepo.ProcessTransaction(fromUser, toUser, amount)
}
