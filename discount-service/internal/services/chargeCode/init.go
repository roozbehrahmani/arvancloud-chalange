package charge_code_service

import (
	"github.com/roozbehrahmani/abrarvan_test/config"
	charge_code_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/chargeCode"
	user_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/user"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB, cnf *config.Config, userRepo *user_repository.UserRepository, chargeCodeRepo *charge_code_repository.ChargeCodeRepository) ServiceInterface {
	return &service{
		db:             db,
		cnf:            cnf,
		userRepo:       userRepo,
		ChargeCodeRepo: chargeCodeRepo,
	}
}
