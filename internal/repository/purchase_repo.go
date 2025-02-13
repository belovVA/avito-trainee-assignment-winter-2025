package repository

import (
	"avito-coin-service/internal/database"
	"avito-coin-service/internal/model"
	"fmt"
)

type PurchaseRepository struct{}

func NewPurchaseRepository() *PurchaseRepository {
	return &PurchaseRepository{}
}

func (r *PurchaseRepository) Create(tx *model.Purchase) error {
	return database.DB.Create(tx).Error
}

func (r *PurchaseRepository) GetByUserAndMerch(userID uint, merchId uint) (*model.Purchase, error) {
	var purchase model.Purchase

	if err := database.DB.Where("user_id = ? AND merch_id = ?", userID, merchId).First(&purchase).Error; err != nil {

		return nil, err
	}

	return &purchase, nil
}

func (r *PurchaseRepository) GetListByUserID(userID uint) ([]*model.Purchase, error) {
	var purchases []*model.Purchase

	if err := database.DB.Where("user_id = ?", userID).Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}

func (r *PurchaseRepository) Update(p model.Purchase) error {
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
