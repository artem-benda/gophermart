package main

import (
	"context"
	"errors"
	"github.com/artem-benda/gophermart"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func mustRunDBMigrations(dbURL string) {
	d, err := iofs.New(gophermart.FS, "db/migrations")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
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
