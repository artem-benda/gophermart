package main

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func mustRunDbMigrations(dbURL string) {
	m, err := migrate.New(
		"file://db/migrations",
		dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func mustCreateConnectionPool(databaseDSN string) *pgxpool.Pool {
	dbPool, err := pgxpool.New(context.Background(), databaseDSN)
	if err != nil {
		panic(err)
	}
	return dbPool
}
