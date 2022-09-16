package charge_code_repository

import (
	"gorm.io/gorm"
)

func Initialize(MysqlDatabase *gorm.DB) *ChargeCodeRepository {

	CahrgeCodeRepo := &ChargeCodeRepository{database: MysqlDatabase}
	return CahrgeCodeRepo
}
