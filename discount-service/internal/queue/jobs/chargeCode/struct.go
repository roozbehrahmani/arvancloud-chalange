package charge_code_job

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	charge_code_service "github.com/roozbehrahmani/abrarvan_test/internal/services/chargeCode"
)

type ChargeWalletJob struct {
	Name              string
	ChargeCode        *models.ChargeCode
	User              *models.User
	ChargeCodeService charge_code_service.ServiceInterface
}
