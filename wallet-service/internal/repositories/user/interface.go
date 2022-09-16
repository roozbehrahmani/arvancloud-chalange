package user_repository

import "github.com/roozbehrahmani/abrarvan_test/internal/models"

type UserRepo interface {
	GetUserWithPhone(phone string) (user *models.User, err error)
}
