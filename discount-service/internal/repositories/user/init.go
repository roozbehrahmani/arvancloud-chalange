package user_repository

import (
	"gorm.io/gorm"
)

func Initialize(MysqlDatabase *gorm.DB) *UserRepository {

	UserRepo := &UserRepository{database: MysqlDatabase}
	return UserRepo
}
