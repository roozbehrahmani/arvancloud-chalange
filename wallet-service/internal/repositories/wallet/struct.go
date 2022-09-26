package wallet_repository

import (
	"gorm.io/gorm"
)

type WalletRepository struct {
	database *gorm.DB
}
