package repository

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/gofiber/fiber/v3"
)

type UserRepository struct {
	DAO dao.User
}

func (r *UserRepository) Register(ctx fiber.Ctx, login string, passwordHash string) (*int64, error) {
	return r.DAO.Insert(ctx, entity.User{Login: login, PasswordHash: passwordHash})
}

func (r *UserRepository) GetUserByLogin(ctx fiber.Ctx, login string) (*entity.User, error) {
	return nil, nil
}
