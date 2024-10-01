package config

import (
	"flag"
	"os"
	"sync"
)

var once sync.Once

var config = struct {
	address        string
	accrualAddress string
	databaseDSN    string
	logLevel       string
}{}

func init() {
	once.Do(func() {
		flagParse()
		envParse()
	})
}

func flagParse() {
	flag.StringVar(&config.address, "address", "localhost:9000", "HTTP server address")
	flag.StringVar(&config.accrualAddress, "accrual_address", ":8080", "Accrual server address")
	flag.StringVar(&config.databaseDSN, "database_dsn", "host=localhost dbname=gophermart sslmode=disable", "Database DSN")
	flag.StringVar(&config.logLevel, "log_level", "info", "Log level")
}

func envParse() {
	envAddress := os.Getenv("address")
	if envAddress != "" {
		config.address = envAddress
	}

	envAccrualAddress := os.Getenv("accrual_address")
	if envAccrualAddress != "" {
		config.accrualAddress = envAccrualAddress
	}

	envDatabaseDSN := os.Getenv("database_dsn")
	if envDatabaseDSN != "" {
		config.databaseDSN = envDatabaseDSN
	}

	envLogLevel := os.Getenv("log_level")
	if envLogLevel != "" {
		config.logLevel = envLogLevel
	}
}

func GetAddress() string {
	return config.address
}

func GetAccrualAddress() string {
	return config.accrualAddress
}

func GetDatabaseDSN() string {
	return config.databaseDSN
}

func GetLogLevel() string {
	return config.logLevel
}
