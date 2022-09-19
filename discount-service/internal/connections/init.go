package connections

import (
	"context"
	"fmt"
	"github.com/roozbehrahmani/abrarvan_test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Connections struct {
	MysqlDatabase *gorm.DB
}

func Init(ctx context.Context, cnf *config.Config) (*Connections, error) {
	log.Printf(cnf.MysqlDatabase.MysqlHost)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cnf.MysqlDatabase.MysqlUsername, cnf.MysqlDatabase.MysqlPassword, cnf.MysqlDatabase.MysqlHost, cnf.MysqlDatabase.MysqlPort, cnf.MysqlDatabase.MysqlDatabase)

	//log.Printf("Connecting to Postgres: ", dsn.String())
	log.Printf("Connecting to Mysql: ", dsn)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//"root:@tcp(127.0.0.1:3306)/milyoonex?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Printf("Connected to mysql: ", dsn)
	return &Connections{db}, nil
}
