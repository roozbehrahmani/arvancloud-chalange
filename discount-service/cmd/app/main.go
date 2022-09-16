package main

import (
	"context"
	"github.com/roozbehrahmani/abrarvan_test/config"
	application "github.com/roozbehrahmani/abrarvan_test/internal/app"
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
	charge_code_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/chargeCode"
	user_repository "github.com/roozbehrahmani/abrarvan_test/internal/repositories/user"
	"github.com/roozbehrahmani/abrarvan_test/internal/router"
	charge_code_service "github.com/roozbehrahmani/abrarvan_test/internal/services/chargeCode"
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

	userRepo := user_repository.Initialize(conns.MysqlDatabase)
	chargeCodeRepo := charge_code_repository.Initialize(conns.MysqlDatabase)

	chargeCodeService := charge_code_service.Initialize(conns.MysqlDatabase, cnf, userRepo, chargeCodeRepo)

	app := application.Create(ctx, *cnf, chargeCodeService)

	//server := router.Initialize()
	server := router.Initialize(ctx, cnf, app)

	if err := server.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
