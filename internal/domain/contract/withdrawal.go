package contract

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
)

type WithdrawalRepository interface {
	GetTotalByUserID(ctx fiber.Ctx, userID int64) (*float64, error)
	GetListByUserID(ctx fiber.Ctx, userID int64) ([]entity.Withdrawal, error)
	Withdraw(ctx fiber.Ctx, userID int64, orderNumber string, amount float64) error
}
