package mock

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/mock"
)

type WithdrawalRepository struct {
	mock.Mock
}

func (r *WithdrawalRepository) GetTotalByUserID(ctx fiber.Ctx, userID int64) (*float64, error) {
	args := r.Called(ctx, userID)
	return args.Get(0).(*float64), args.Error(1)
}

func (r *WithdrawalRepository) GetListByUserID(ctx fiber.Ctx, userID int64) ([]entity.Withdrawal, error) {
	args := r.Called(ctx, userID)
	return args.Get(0).([]entity.Withdrawal), args.Error(1)
}

func (r *WithdrawalRepository) Withdraw(ctx fiber.Ctx, userID int64, orderNumber string, amount float64) error {
	args := r.Called(ctx, userID, orderNumber, amount)
	return args.Error(0)
}
