package mock

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (r *UserRepository) Register(ctx fiber.Ctx, login string, passwordHash string) (*int64, error) {
	args := r.Called(ctx, login, passwordHash)
	return args.Get(0).(*int64), args.Error(1)
}

func (r *UserRepository) GetUserByLogin(ctx fiber.Ctx, login string) (*entity.User, error) {
	args := r.Called(ctx, login)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (r *UserRepository) GetUserByID(ctx fiber.Ctx, userID int64) (*entity.User, error) {
	args := r.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}
