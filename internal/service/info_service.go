package service

import (
	"avito-coin-service/internal/model"
	"avito-coin-service/internal/repository"
	"fmt"
)

type InfoService struct {
	userRepo     *repository.UserRepository
	txRepo       *repository.TransactionRepository
	merchRepo    *repository.MerchRepository
	purchaseRepo *repository.PurchaseRepository
}

func NewInfoService(
	userRepo *repository.UserRepository,
	txRepo *repository.TransactionRepository,
	merchRepo *repository.MerchRepository,
	purchaseRepo *repository.PurchaseRepository,

) *InfoService {

	return &InfoService{
		userRepo:     userRepo,
		txRepo:       txRepo,
		merchRepo:    merchRepo,
		purchaseRepo: purchaseRepo,
	}
}

func (s *InfoService) GetInfo(userName string) (*model.InfoResponse, error) {
	user, err := s.userRepo.GetByName(userName)
	if err != nil {
		return nil, fmt.Errorf("пользователь не найден")
	}
	var info model.InfoResponse
	info.Coins = user.Balance
	inventory, _ := s.purchaseRepo.GetListByUserID(user.ID)
	info.Inventory = make([]model.InventoryItem, 0, len(inventory))
	if inventory != nil {
		for _, p := range inventory {
			merch, err := s.merchRepo.GetByID(p.MerchID)
			if err != nil {
				continue
			}
			var item model.InventoryItem
			item.Type = merch.Name
			item.Quantity = p.Count
			info.Inventory = append(info.Inventory, item)

		}
	}

	recieved, _ := s.txRepo.GetListRecievedTransactionByID(user.ID)
	info.CoinHistory.Received = make([]model.CoinTransaction, 0, len(recieved))
	if recieved != nil {
		for _, t := range recieved {
			var trx model.CoinTransaction
			user, err := s.userRepo.GetByID(t.FromUser)

			if err != nil {
				return nil, fmt.Errorf("500")
			}
			trx.FromUser = user.Name
			trx.Amount = uint(t.Amount)
			info.CoinHistory.Received = append(info.CoinHistory.Received, trx)

		}
	}

	sent, err := s.txRepo.GetListSentTransactionByID(user.ID)
	info.CoinHistory.Sent = make([]model.CoinTransaction, 0, len(sent))
	if err == nil {
		for _, t := range sent {
			var trx model.CoinTransaction
			user, err := s.userRepo.GetByID(t.ToUser)
			if err != nil {
				return nil, fmt.Errorf("500")
			}
			trx.ToUser = user.Name
			trx.Amount = uint(t.Amount)
			info.CoinHistory.Sent = append(info.CoinHistory.Sent, trx)

		}
	}
	return &info, nil
}
