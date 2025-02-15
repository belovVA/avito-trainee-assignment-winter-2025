package service

import (
	"avito-coin-service/internal/repository"
	"fmt"
)

type PurchaseService interface {
	BuyMerch(userName string, merchName string) error
}

type purchase struct {
	userRepo     repository.UserRepository
	merchRepo    repository.MerchRepository
	purchaseRepo repository.PurchaseRepository
}

func NewPurchaseService(
	userRepo repository.UserRepository,
	merchRepo repository.MerchRepository,
	purchaseRepo repository.PurchaseRepository,
) *purchase {

	return &purchase{
		userRepo:     userRepo,
		merchRepo:    merchRepo,
		purchaseRepo: purchaseRepo,
	}
}

func (s *purchase) BuyMerch(userName string, merchName string) error {

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

	return s.purchaseRepo.ProcessPurchase(user, merch)
}
