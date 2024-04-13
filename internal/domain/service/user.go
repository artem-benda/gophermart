package service

import (
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	UserRepository contract.UserRepository
}

func (s User) Register(ctx fiber.Ctx, login string, password string) error {
	_, err := s.UserRepository.Register(ctx, login, password)
	return err
}
