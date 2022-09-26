package charge_code_job

import "github.com/roozbehrahmani/abrarvan_test/internal/models"

type chargeWalletJobDo interface {
	doWork(chargeCode *models.ChargeCode, user *models.User) (map[string]interface{}, error)
}
