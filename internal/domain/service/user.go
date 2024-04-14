package service

import (
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	UserRepository contract.UserRepository
}

func (s User) Register(ctx fiber.Ctx, login string, password string) (*int64, error) {
	return s.UserRepository.Register(ctx, login, password)
}
