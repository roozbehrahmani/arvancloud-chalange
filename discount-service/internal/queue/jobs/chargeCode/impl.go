package charge_code_job

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
)

func (r *ChargeWalletJob) DoWork(chargeCode *models.ChargeCode, user *models.User) (map[string]interface{}, error) {
	response, err := r.ChargeCodeService.ChargeWalletJob(chargeCode, user)
	if err != nil {
		return nil, err
	}
	return response, err
}
