package user_repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}
