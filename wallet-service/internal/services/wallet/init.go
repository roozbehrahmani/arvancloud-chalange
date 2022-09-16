package wallet_service

import (
	"github.com/roozbehrahmani/abrarvan_test/config"
	user_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/user"
	"github.com/roozbehrahmani/abrarvan_test/internal/repositories/wallet"
	wallet_transaction_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/walletTransaction"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB, cnf *config.Config, walletRepo *wallet_repository.WalletRepository, walletTransactionRepo *wallet_transaction_repository.WalletTransactionRepository, userRepo *user_repository.UserRepository) ServiceInterface {
	return &service{
		db:                    db,
		cnf:                   cnf,
		walletRepo:            walletRepo,
		userRepo:              userRepo,
		walletTransactionRepo: walletTransactionRepo,
	}
}
