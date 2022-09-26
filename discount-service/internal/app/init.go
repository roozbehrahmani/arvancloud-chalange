package application

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	charge_code_dispatcher "github.com/roozbehrahmani/abrarvan_test/internal/queue/dispatchers/chargeCode"
	charge_code_service "github.com/roozbehrahmani/abrarvan_test/internal/services/chargeCode"
)

type Application struct {
	Config            config.Config
	ChargeCodeService charge_code_service.ServiceInterface
	dispatcher        charge_code_dispatcher.ChargeWalletDispatcher
}

func Create(ctx context.Context, cnf config.Config, chargeCodeService charge_code_service.ServiceInterface, chargeWalletDispatcher charge_code_dispatcher.ChargeWalletDispatcher) *Application {
	return &Application{
		cnf,
		chargeCodeService,
		chargeWalletDispatcher,
	}
}
