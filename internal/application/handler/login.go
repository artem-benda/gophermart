package handler

import (
	"github.com/artem-benda/gophermart/internal/domain/service/user"
	"github.com/gofiber/fiber/v3"
)

type deps struct {
	svc user.Login
}

func Login(ctx fiber.Ctx) error {

}
