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

const appName = "surface"

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
	userSeeder(conns)
	chargeCodeSeeder(conns)

}

func userSeeder(conns *connections.Connections) {
	user := models.User{}
	conns.MysqlDatabase.Where("deleted_at = NULL").Delete(&user)

	phoneNumber := "09123456789"
	user = models.User{
		Phone: &phoneNumber,
	}
	conns.MysqlDatabase.Create(&user)

}

func chargeCodeSeeder(conns *connections.Connections) {
	chargeCode := models.ChargeCode{
		Code:     "free_5000",
		Amount:   5000000,
		MaxUsage: 1000,
		Status:   true,
	}
	conns.MysqlDatabase.Create(&chargeCode)

}
