package config

type Config struct {
	MysqlDatabase
	ServerPort int
	Secret     string
}

type PGDatabase struct {
	Host     string
	port     string
	Database string
	Username string
	Password string
}

type MysqlDatabase struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}
