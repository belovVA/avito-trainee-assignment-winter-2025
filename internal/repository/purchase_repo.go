package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
	"fmt"
)

type IPurchaseRepository interface {
	Create(tx *model.Purchase) error
	GetByUserAndMerch(userID uint, merchId uint) (*model.Purchase, error)
	GetListByUserID(userID uint) ([]*model.Purchase, error)
	Update(p model.Purchase) error
}

type purchaseRepository struct{}

func NewPurchaseRepository() IPurchaseRepository {
	return &purchaseRepository{}
}

func (r *purchaseRepository) Create(tx *model.Purchase) error {
	return database.DB.Create(tx).Error
}

func (r *purchaseRepository) GetByUserAndMerch(userID uint, merchId uint) (*model.Purchase, error) {
	var purchase model.Purchase

	if err := database.DB.Where("user_id = ? AND merch_id = ?", userID, merchId).First(&purchase).Error; err != nil {

		return nil, err
	}

	return &purchase, nil
}

func (r *purchaseRepository) GetListByUserID(userID uint) ([]*model.Purchase, error) {
	var purchases []*model.Purchase

	if err := database.DB.Where("user_id = ?", userID).Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}

func (r *purchaseRepository) Update(p model.Purchase) error {
	result := database.DB.Model(&model.Purchase{}).
		Where("user_id = ? AND merch_id = ?", p.UserID, p.MerchID).
		Update("count", p.Count)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("покупка не найдена")
	}

	return nil
}
