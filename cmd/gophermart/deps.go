package main

import "github.com/jackc/pgx/v5/pgxpool"

type AppDependencies struct {
	Config Config
	DB     pgxpool.Pool
}
