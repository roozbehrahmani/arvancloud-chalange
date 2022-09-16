package wallet_service

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
)

type ServiceInterface interface {
	GetWalletWithPhone(phone string) (*models.Wallet, error)
	GetWalletWithUser(*models.User) (*models.Wallet, error)
	WalletCharging(wallet *models.Wallet, amount float64) (*models.Wallet, error)
}
