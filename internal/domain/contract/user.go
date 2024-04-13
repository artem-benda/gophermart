package contract

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
)

type UserRepository interface {
	Register(ctx fiber.Ctx, login string, password string) (*int64, error)
	GetUserByLogin(ctx fiber.Ctx, login string) (*entity.User, error)
}
