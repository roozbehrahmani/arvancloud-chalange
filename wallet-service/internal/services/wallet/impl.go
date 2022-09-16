package wallet_service

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
)

func (s *service) GetWalletWithPhone(phone string) (*models.Wallet, error) {

	user, err := s.userRepo.GetUserWithPhone(&phone)
	if err != nil {
		return nil, err
	}
	return s.walletRepo.GetWalletWithUser(user)
}

func (s *service) GetWalletWithUser(user *models.User) (*models.Wallet, error) {
	return s.walletRepo.GetWalletWithUser(user)
}

func (s *service) WalletCharging(wallet *models.Wallet, amount float64) (*models.Wallet, error) {
	tx := s.db.Begin()
	err := s.walletRepo.IncreaseAmountWithTransaction(tx, wallet, amount)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = s.walletTransactionRepo.CreateWithTransaction(tx, wallet.ID, amount, true)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return wallet, nil
}
