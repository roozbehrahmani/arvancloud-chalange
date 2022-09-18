package config

type Config struct {
	MysqlDatabase
	DiscountService
	ServerIP   string
	ServerPort int
	Secret     string
}

type DiscountService struct {
	DiscountServiceHost string
	DiscountServicePORT int
}

type MysqlDatabase struct {
	MysqlHost         string
	MysqlPort         int
	MysqlDatabase     string
	MysqlUsername     string
	MysqlPassword     string
	MysqlRootPassword string
}
