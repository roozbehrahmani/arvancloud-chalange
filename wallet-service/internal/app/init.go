package application

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	wallet_service "github.com/roozbehrahmani/abrarvan_test/internal/services/wallet"
)

type Application struct {
	Config        config.Config
	WalletService wallet_service.ServiceInterface
}

func Create(ctx context.Context, cnf config.Config, walletService wallet_service.ServiceInterface) *Application {
	return &Application{
		cnf,
		walletService,
	}
}
