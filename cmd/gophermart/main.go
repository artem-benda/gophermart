package main

import (
	"github.com/artem-benda/gophermart/internal/application/handler"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/artem-benda/gophermart/internal/infrastructure/repository"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
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
	v := mustCreateValidator()

	withdrawalDAO := dao.Withdrawal{DB: dbPool}
	withdrawalRepository := repository.WithdrawalRepository{DAO: withdrawalDAO}
	withdrawalService := service.Withdrawal{WithdrawalRepository: &withdrawalRepository}

	userDAO := dao.User{DB: dbPool}
	userRepository := repository.UserRepository{DAO: userDAO}
	userService := service.User{UserRepository: &userRepository, WithdrawalRepository: &withdrawalRepository, Salt: cfg.mustGetSalt()}

	orderDAO := dao.Order{DB: dbPool}
	orderRepository := repository.OrderRepository{DAO: orderDAO}
	orderService := service.Order{OrderRepository: &orderRepository}

	app.Post("/api/user/register", handler.NewRegisterUserHandler(&userService, v))
	app.Post("/api/user/login", handler.NewLoginHandler(&userService, v))
	app.Post("/api/user/orders", handler.NewUploadOrderHandler(&orderService, v))
	app.Get("/api/user/orders", handler.NewGetUserOrdersHandler(&orderService))
	app.Get("/api/user/balance", handler.NewGetUserBalanceHandler(&userService))
	app.Post("/api/user/balance/withdraw", handler.NewWithdrawHandler(&withdrawalService, v))
	app.Get("/api/user/withdrawals", handler.NewGetWithdrawalsHandler(&withdrawalService))
	log.Fatal(app.Listen(cfg.Endpoint))
}
