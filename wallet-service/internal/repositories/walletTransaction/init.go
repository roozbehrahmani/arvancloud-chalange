package wallet_transaction_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
)

func Initialize(conns *connections.Connections) *WalletTransactionRepository {

	WalletTransactionRepo := &WalletTransactionRepository{database: conns.MysqlDatabase}
	return WalletTransactionRepo
}
