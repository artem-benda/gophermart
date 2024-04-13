package handler

import (
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	svc service.User
}

func Login(ctx fiber.Ctx) error {
	return nil
}
