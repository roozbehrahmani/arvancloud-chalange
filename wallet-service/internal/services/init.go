package services

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
)

type Service struct {
	config     *config.Config
	connection *connections.Connections
}

func Init(ctx context.Context, cnf *config.Config, conns *connections.Connections) *Service {
	return &Service{
		cnf,
		conns,
	}
}
