package wallet_transaction_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"gorm.io/gorm"
)

func (r *WalletTransactionRepository) CreateWithTransaction(tx *gorm.DB, walletId uint, amount float64, IncrementOrDecrementType bool) error {
	walletTransaction := models.WalletTransaction{
		WalletId: walletId,
		Amount:   amount,
		Type:     IncrementOrDecrementType,
	}
	return tx.Create(&walletTransaction).Error
}
