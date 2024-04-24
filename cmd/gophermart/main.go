package main

import (
	"context"
	"github.com/artem-benda/gophermart/internal/application/handler"
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/application/worker"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"os"
)

func main() {
	log.SetLevel(log.LevelDebug)
	log.SetOutput(os.Stdout)
	log.Debug("starting application")
	cfg := mustReadConfig()
	mustRunDBMigrations(cfg.DatabaseDSN)
	dbPool := mustCreateConnectionPool(cfg.DatabaseDSN)

	app := fiber.New()
	app.Use(logger.New())
	v := mustCreateValidator()

	deps := mustSetupAppDependencies(dbPool, cfg)

	workerFn := worker.NewAccrualWorkerFunc(deps.AccrualService, context.Background())
	go workerFn()

	app.Post("/api/user/register", handler.NewRegisterUserHandler(deps.UserService, v))
	app.Post("/api/user/login", handler.NewLoginHandler(deps.UserService, v))

	auth := middleware.NewAuthMiddleware()
	// Не используем .Use для middleware, т.к. нет общего пути для авторизоавнных и неавт. запросов
	app.Post("/api/user/orders", handler.NewUploadOrderHandler(deps.OrderService, v), auth)
	app.Get("/api/user/orders", handler.NewGetUserOrdersHandler(deps.OrderService), auth)
	app.Get("/api/user/balance", handler.NewGetUserBalanceHandler(deps.UserService), auth)
	app.Post("/api/user/balance/withdraw", handler.NewWithdrawHandler(deps.WithdrawalService, v), auth)
	app.Get("/api/user/withdrawals", handler.NewGetWithdrawalsHandler(deps.WithdrawalService), auth)
	log.Fatal(app.Listen(cfg.Endpoint))
}
