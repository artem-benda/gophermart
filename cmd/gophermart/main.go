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
	userService := service.User{UserRepository: &userRepository, Salt: cfg.mustGetSalt()}

	app.Post("/api/user/register", handler.NewRegisterUserHandler(&userService, v))
	app.Post("/api/user/login", handler.NewLoginHandler(&userService, v))
	log.Fatal(app.Listen(cfg.Endpoint))
}
