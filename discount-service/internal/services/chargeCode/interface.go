package charge_code_service

import "github.com/roozbehrahmani/abrarvan_test/internal/models"

type ServiceInterface interface {
	ChargeWalletJob(chargeCode *models.ChargeCode, user *models.User) (map[string]interface{}, error)
	Validation(code string, phone string) (*models.User, *models.ChargeCode, error)
}
