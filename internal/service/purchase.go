package service

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
	"avito-coin-service/internal/repository"
	"fmt"
)

type IPurchaseService interface {
	BuyMerch(userName string, merchName string) error
}

type purchaseService struct {
	userRepo     repository.IUserRepository
	merchRepo    repository.IMerchRepository
	purchaseRepo repository.IPurchaseRepository
}

func NewPurchaseService(
	userRepo repository.IUserRepository,
	merchRepo repository.IMerchRepository,
	purchaseRepo repository.IPurchaseRepository,
) *purchaseService {

	return &purchaseService{
		userRepo:     userRepo,
		merchRepo:    merchRepo,
		purchaseRepo: purchaseRepo,
	}
}

func (s *purchaseService) BuyMerch(userName string, merchName string) error {

	user, err := s.userRepo.GetByName(userName)
	if err != nil {
		return fmt.Errorf("пользователь не найден")
	}

	merch, err := s.merchRepo.GetByName(merchName)
	if err != nil {
		return fmt.Errorf("мерч не найден")
	}

	if user.Balance < merch.Price {
		return fmt.Errorf("Недостаточно средств на балансе")
	}

	// Начинаем транзакцию
	tx := database.DB.Begin()

	user.Balance -= merch.Price

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if purchase, err := s.purchaseRepo.GetByUserAndMerch(user.ID, merch.ID); err != nil {
		purchase := model.Purchase{
			UserID:  user.ID,
			MerchID: merch.ID,
			Count:   1,
		}

		if err := s.purchaseRepo.Create(&purchase); err != nil {
			tx.Rollback()
			return err
		}

	} else if purchase != nil {

		purchase.Count += 1
		if err := s.purchaseRepo.Update(*purchase); err != nil {
			tx.Rollback()
			return fmt.Errorf("Не удалось обновить данные")
		}
	} else {
		tx.Rollback()
		return fmt.Errorf("Ошибка сервера")
	}

	return tx.Commit().Error
}
