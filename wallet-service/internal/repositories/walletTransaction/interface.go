package wallet_transaction_repository

import "gorm.io/gorm"

type WalletTransactionRepo interface {
	CreateWithTransaction(tx *gorm.DB, walletId uint, amount float64, IncrementOrDecrementType bool) error
}
