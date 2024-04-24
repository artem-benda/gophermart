package service

import (
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
)

type Withdrawal struct {
	WithdrawalRepository contract.WithdrawalRepository
}

func (w Withdrawal) GetList(ctx fiber.Ctx, userID int64) ([]entity.Withdrawal, error) {
	return w.WithdrawalRepository.GetListByUserID(ctx, userID)
}

func (w Withdrawal) Withdraw(ctx fiber.Ctx, userID int64, orderNumber string, amount float64) error {
	return w.WithdrawalRepository.Withdraw(ctx, userID, orderNumber, amount)
}
