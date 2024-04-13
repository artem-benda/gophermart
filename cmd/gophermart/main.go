package main

import (
	"github.com/artem-benda/gophermart/internal/application/handler"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	cfg := mustReadConfig()
	mustRunDbMigrations(cfg.DatabaseDSN)
	dbPool := mustCreateConnectionPool(cfg.DatabaseDSN)
	app := fiber.New()
	app.Post("/api/user/register", handler.RegisterUser)
	app.Post("/api/user/login", handler.Login)
	log.Fatal(app.Listen(cfg.Endpoint))
}
