package config

type Config struct {
	MysqlDatabase
	WalletService
	ServerIP   string
	ServerPort int
}

type WalletService struct {
	WalletServiceSecret string
	WalletServiceHost   string
	WalletServicePORT   int
}

type MysqlDatabase struct {
	MysqlHost         string
	MysqlPort         int
	MysqlDatabase     string
	MysqlUsername     string
	MysqlPassword     string
	MysqlRootPassword string
}
