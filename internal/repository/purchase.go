package repository

import (
	"fmt"

	"avito-coin-service/internal/model"

	"gorm.io/gorm"
)

type PurchaseRepository interface {
	Create(tx *model.Purchase) error
	GetByUserAndMerch(userID uint, merchId uint) (*model.Purchase, error)
	GetListByUserID(userID uint) ([]*model.Purchase, error)
	Update(p *model.Purchase) error
	ProcessPurchase(user *model.User, merch *model.Merch) error
}

type purchaseRep struct {
	DB *gorm.DB
}

func NewPurchaseRepository(DB *gorm.DB) PurchaseRepository {
	return &purchaseRep{
		DB: DB,
	}
}

func (r *purchaseRep) Create(tx *model.Purchase) error {
	return r.DB.Create(tx).Error
}

func (r *purchaseRep) GetByUserAndMerch(userID uint, merchId uint) (*model.Purchase, error) {
	var purchase model.Purchase

	if err := r.DB.Where("user_id = ? AND merch_id = ?", userID, merchId).First(&purchase).Error; err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (r *purchaseRep) GetListByUserID(userID uint) ([]*model.Purchase, error) {
	var purchases []*model.Purchase

	if err := r.DB.Where("user_id = ?", userID).Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}

func (r *purchaseRep) Update(p *model.Purchase) error {
	result := r.DB.Model(&model.Purchase{}).
		Where("user_id = ? AND merch_id = ?",
			p.UserID,
			p.MerchID).
		Update("count", p.Count)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("покупка не найдена")
	}

	return nil
}

func (r *purchaseRep) ProcessPurchase(user *model.User, merch *model.Merch) error {
	tx := r.DB.Begin()

	// Updating user balance
	user.Balance -= merch.Price

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()

		return err
	}

	purchase, err := r.GetByUserAndMerch(user.ID, merch.ID)

	if err != nil {

		newPurchase := model.Purchase{
			UserID:  user.ID,
			MerchID: merch.ID,
			Count:   1,
		}

		if err := tx.Create(&newPurchase).Error; err != nil {
			tx.Rollback()

			return err
		}

	} else {

		purchase.Count += 1

		if err := tx.Save(purchase).Error; err != nil {
			tx.Rollback()

			return fmt.Errorf("couldn't update data")
		}
	}

	return tx.Commit().Error
}
