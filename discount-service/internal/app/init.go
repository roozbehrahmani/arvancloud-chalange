package application

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	charge_code_service "github.com/roozbehrahmani/abrarvan_test/internal/services/chargeCode"
)

type Application struct {
	Config            config.Config
	ChargeCodeService charge_code_service.ServiceInterface
}

func Create(ctx context.Context, cnf config.Config, chargeCodeService charge_code_service.ServiceInterface) *Application {
	return &Application{
		cnf,
		chargeCodeService,
	}
}
