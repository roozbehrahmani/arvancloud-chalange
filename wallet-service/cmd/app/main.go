package main

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	application "github.com/roozbehrahmani/abrarvan_test/internal/app"
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
	user_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/user"
	wallet_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/wallet"
	wallet_transaction_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/walletTransaction"
	"github.com/roozbehrahmani/abrarvan_test/internal/router"
	wallet_service "github.com/roozbehrahmani/abrarvan_test/internal/services/wallet"
	"log"
	"os"
)

func main() {
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()
	cnf, err := config.Init(ctx)
	if err != nil {
		log.Fatalf("can't read configs:%v", err)
		os.Exit(1)
	}
	conns, err := connections.Init(ctx, cnf)
	if err != nil {
		log.Fatalf("can't init connections:%v", err)
		os.Exit(1)
	}

	userRepo := user_repository.Initialize(conns)
	walletRepo := wallet_repository.Initialize(conns)
	walletTransactionRepo := wallet_transaction_repository.Initialize(conns)

	walletService := wallet_service.Initialize(conns.MysqlDatabase, cnf, walletRepo, walletTransactionRepo, userRepo)

	app := application.Create(ctx, *cnf, walletService)

	server := router.Initialize(ctx, cnf, app)

	if err := server.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
