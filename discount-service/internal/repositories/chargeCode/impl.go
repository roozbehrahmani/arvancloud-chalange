package charge_code_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/app/helpers"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func (r *ChargeCodeRepository) ValidateCode(code string) (*models.ChargeCode, error) {
	var chargeCode models.ChargeCode

	err := r.database.
		Where("code = ?", code).Where("(expire_at > ? OR expire_at is NULL)", time.Now()).
		First(&chargeCode).Error
	if err != nil {
		return nil, err
	}
	if chargeCode.MaxUsage <= chargeCode.CurrentUsage || chargeCode.Status == false {
		return nil, &helpers.CodeNotValidError{}
	}

	return &chargeCode, nil
}

func (r *ChargeCodeRepository) IncreaseChargeCodeUsageWithTransaction(tx *gorm.DB, chargeCode *models.ChargeCode) error {
	chargeCode.CurrentUsage += 1
	if chargeCode.CurrentUsage >= chargeCode.MaxUsage {
		chargeCode.Status = false
	}
	return tx.Clauses(clause.Locking{Strength: "UPDATE"}).Updates(&chargeCode).Error
}

func (r *ChargeCodeRepository) UserCanUseChargeCode(userId uint, chargeCodeId uint) bool {
	var count int64
	r.database.Table("charge_code_users").
		Where("user_id = ?", userId).
		Where("charge_code_id = ?", chargeCodeId).
		Count(&count)
	return count == 0
}

func (r *ChargeCodeRepository) CreateUserChargeCodeRecordWithTransaction(tx *gorm.DB, userId uint, chargeCodeId uint) error {
	return tx.Table("charge_code_users").Create(&models.ChargeCodeUser{
		UserId:       userId,
		ChargeCodeId: chargeCodeId,
	}).Error
}
