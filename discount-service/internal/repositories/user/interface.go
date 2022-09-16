package user_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserWithPhoneWithTransaction(tx *gorm.DB, phone string) (user *models.User, err error)
}
