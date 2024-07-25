package contract

import (
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

type UserRepository interface {
	Register(ctx fiber.Ctx, login string, passwordHash string) (*int64, error)
	GetUserByLogin(ctx fiber.Ctx, login string) (*entity.User, error)
	GetUserByID(ctx fiber.Ctx, userID int64) (*entity.User, error)
}
