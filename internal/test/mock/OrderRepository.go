package mock

import (
	"context"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/mock"
)

type OrderRepository struct {
	mock.Mock
}

func (r *OrderRepository) Upload(ctx fiber.Ctx, userID int64, orderNumber string) error {
	args := r.Called(ctx, userID, orderNumber)
	return args.Error(0)
}

func (r *OrderRepository) GetByUserID(ctx fiber.Ctx, userID int64) ([]entity.Order, error) {
	args := r.Called(ctx, userID)
	return args.Get(0).([]entity.Order), args.Error(1)
}

func (r *OrderRepository) GetListToSyncAccruals(ctx context.Context) ([]entity.Order, error) {
	args := r.Called(ctx)
	return args.Get(0).([]entity.Order), args.Error(1)
}
