package handler

import (
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	svc      *service.User
	validate *validator.Validate
}

func Login(ctx fiber.Ctx) error {
	return nil
}
