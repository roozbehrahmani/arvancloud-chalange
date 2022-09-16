package wallet_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"gorm.io/gorm"
)

type WalletRepo interface {
	GetWalletWithUser(user *models.User) (wallet *models.Wallet, err error)
	IncreaseAmountWithTransaction(tx *gorm.DB, wallet *models.Wallet, amount float64) error
}
