package user_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
)

func Initialize(conns *connections.Connections) *UserRepository {

	UserRepo := &UserRepository{database: conns.MysqlDatabase}
	return UserRepo
}
