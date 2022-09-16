package charge_code_repository

import (
	"gorm.io/gorm"
)

type ChargeCodeRepository struct {
	database *gorm.DB
}

var database *gorm.DB
