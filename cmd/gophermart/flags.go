package main

import (
	"flag"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Endpoint        string `env:"RUN_ADDRESS"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
	AccrualEndpoint string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	LogLevel        string `env:"LOG_LEVEL"`
	Salt            string `env:"SALT"`
}

func mustReadConfig() Config {
	var config Config

	flag.StringVar(&config.Endpoint, "a", "localhost:8080", "address and port of metrics server")
	flag.StringVar(&config.DatabaseDSN, "d", "postgres://ac_gophermart_backend:ac_gophermart_backend123@localhost:5432/gophermart?sslmode=disable", "Database connection URL in pgx format, for ex. postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10")
	flag.StringVar(&config.AccrualEndpoint, "r", "localhost:8080", "address and port of accrual service")
	flag.StringVar(&config.LogLevel, "l", "debug", "logging level: debug, info, warn, error, dpanic, panic, fatal")
	flag.StringVar(&config.Salt, "s", "", "salt in base64std format, using for hashing passwords, at least 8 bytes is recommended by the RFC")
	flag.Parse()

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
