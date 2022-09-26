package main

import (
	"context"
	"fmt"
	"github.com/roozbehrahmani/abrarvan_test/config"
	"github.com/roozbehrahmani/abrarvan_test/internal/connections"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
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
	if conns == nil {
		log.Printf("nile")
		panic(fmt.Sprintf("can't init connections:%v", err))
	}
	err = conns.MysqlDatabase.AutoMigrate(models.User{}, models.Wallet{}, models.WalletTransaction{}, models.ChargeCode{})
	checkHasError(err)

}

func checkHasError(err error) {
	if err != nil {
		panic(fmt.Sprintf("can't init automigrate:%v", err))
	}

}
