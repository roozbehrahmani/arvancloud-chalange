package charge_code_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"gorm.io/gorm"
)

type ChargeCodeRepo interface {
	ValidateCode(code string) (*models.ChargeCode, error)
	UserCanUseChargeCode(userId uint64, chargeCodeId uint64) bool
	IncreaseChargeCodeUsageWithTransaction(tx *gorm.DB, chargeCode *models.ChargeCode) error
	CreateUserChargeCodeRecordWithTransaction(tx *gorm.DB, userId uint64, chargeCodeId uint64) error
}
