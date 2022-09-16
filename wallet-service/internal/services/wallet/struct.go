package wallet_service

import (
	"github.com/roozbehrahmani/abrarvan_test/config"
	user_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/user"
	wallet_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/wallet"
	wallet_transaction_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/walletTransaction"
	"gorm.io/gorm"
)

type service struct {
	db                    *gorm.DB
	cnf                   *config.Config
	walletRepo            *wallet_repository.WalletRepository
	walletTransactionRepo *wallet_transaction_repository.WalletTransactionRepository
	userRepo              *user_repository.UserRepository
}
