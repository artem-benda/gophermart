package main

func main() {
	cfg := mustReadConfig()
	mustRunDbMigrations(cfg.DatabaseDSN)

}
