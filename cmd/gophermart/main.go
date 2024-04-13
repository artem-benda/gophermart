package main

import (
	"github.com/artem-benda/gophermart/internal/application/handler"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/artem-benda/gophermart/internal/infrastructure/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	cfg := mustReadConfig()
	mustRunDbMigrations(cfg.DatabaseDSN)
	dbPool := mustCreateConnectionPool(cfg.DatabaseDSN)
	app := fiber.New()
	v := validator.New()
	userDAO := dao.User{DB: dbPool}
	userRepository := repository.UserRepository{DAO: userDAO}
	registerUserService := service.User{UserRepository: &userRepository}
	app.Post("/api/user/register", handler.NewRegisterUserHandler(&registerUserService, v))
	app.Post("/api/user/login", handler.Login)
	log.Fatal(app.Listen(cfg.Endpoint))
}
