package main

import (
	"encoding/base64"
	"flag"
	"github.com/gofiber/fiber/v3/log"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Endpoint        string `env:"RUN_ADDRESS"`
	DatabaseDSN     string `env:"DATABASE_URI"`
	AccrualEndpoint string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	LogLevel        string `env:"LOG_LEVEL"`
	Salt            string `env:"SALT"`
}

func mustReadConfig() Config {
	var config Config

	flag.StringVar(&config.Endpoint, "a", "localhost:8080", "address and port of server")
	flag.StringVar(&config.DatabaseDSN, "d", "postgres://ac_gophermart_backend:ac_gophermart_backend123@localhost:5432/gophermart?sslmode=disable", "Database connection URL in pgx format, for ex. postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10")
	flag.StringVar(&config.AccrualEndpoint, "r", "http://localhost:8080", "address and port of accrual service")
	flag.StringVar(&config.LogLevel, "l", "debug", "logging level: debug, info, warn, error, dpanic, panic, fatal")
	flag.StringVar(&config.Salt, "s", "BPjkLEqJfARvsYGW++WRcnCjxHyZsrnxXd/qdzpWIaE=", "salt in base64std format, using for hashing passwords, at least 8 bytes is recommended by the RFC")
	flag.Parse()

	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	return config
}

func (c Config) mustGetSalt() []byte {
	salt, err := base64.StdEncoding.DecodeString(c.Salt)

	if err != nil {
		log.Fatal(err)
	}

	return salt
}
