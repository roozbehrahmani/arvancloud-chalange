package wallet_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
)

func Initialize(conns *connections.Connections) *WalletRepository {

	WalletRepo := &WalletRepository{database: conns.MysqlDatabase}
	return WalletRepo
}
