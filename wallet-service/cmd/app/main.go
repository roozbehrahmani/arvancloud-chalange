package main

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	"github.com/roozbehrahmani/abrarvan_test/internal/router"
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
	//conns, err := connections.Init(ctx, cnf)
	//if err != nil {
	//	log.Fatalf("can't init connections:%v", err)
	//	os.Exit(1)
	//}
	//
	//userRepo := user_repository.Initialize(conns)
	//walletRepo := wallet_repository.Initialize(conns)
	//walletTransactionRepo := wallet_transaction_repository.Initialize(conns)
	//
	//walletService := wallet_service.Initialize(conns.MysqlDatabase, cnf, walletRepo, walletTransactionRepo, userRepo)
	//
	//app := application.Create(ctx, *cnf, walletService)
	//
	server := router.Initialize(ctx, cnf, nil)

	if err := server.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
