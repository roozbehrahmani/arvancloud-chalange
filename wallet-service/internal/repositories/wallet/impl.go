package wallet_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"gorm.io/gorm"
)

func (r *WalletRepository) GetWalletWithUser(user *models.User) (*models.Wallet, error) {
	var wallet models.Wallet
	err := r.database.Preload("WalletTransaction").Where("user_id = ?", user.ID).First(&wallet).Error
	if err != nil {
		if err.Error() == noRecordError {
			newWallet := models.Wallet{
				UserId:        user.ID,
				Inventory:     0,
				Balance:       0,
				BlockedAmount: 0,
			}
			return r.createWallet(&newWallet)
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) createWallet(newWallet *models.Wallet) (*models.Wallet, error) {
	err := r.database.Create(newWallet).Error
	if err != nil {
		return nil, err
	}
	return newWallet, err
}

func (r *WalletRepository) IncreaseAmountWithTransaction(tx *gorm.DB, wallet *models.Wallet, amount float64) error {
	wallet.Inventory += amount
	wallet.Balance += amount
	return tx.Updates(&wallet).Error
}
