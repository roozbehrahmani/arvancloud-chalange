package wallet_transaction_repository

import (
	"gorm.io/gorm"
)

type WalletTransactionRepository struct {
	database *gorm.DB
}
