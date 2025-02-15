package service

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
	"avito-coin-service/internal/repository"
	"fmt"
)

type TransactionService struct {
	userRepo *repository.UserRepository
	txRepo   *repository.TransactionRepository
}

func NewTransactionService(userRepo *repository.UserRepository, txRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		userRepo: userRepo,
		txRepo:   txRepo,
	}
}

// SendCoins — переводит монеты от одного пользователя другому
func (s *TransactionService) SendCoins(fromUserName string, toUserName string, amount int) error {
	// Получаем отправителя
	fromUser, err := s.userRepo.GetByName(fromUserName)
	if err != nil {
		return fmt.Errorf("отправитель не найден")
	}

	// Получаем получателя
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

	// Начинаем транзакцию
	tx := database.DB.Begin()

	// Обновляем балансы
	fromUser.Balance -= amount
	toUser.Balance += amount

	// Обновляем в БД
	if err := tx.Save(fromUser).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(toUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Создаём запись о транзакции
	transaction := model.Transaction{
		FromUser: fromUser.ID,
		ToUser:   toUser.ID,
		Amount:   amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Фиксируем изменения
	return tx.Commit().Error
}
