package repository

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/gofiber/fiber/v3"
)

type WithdrawalRepository struct {
	DAO dao.Withdrawal
}

func (r *WithdrawalRepository) GetTotalByUserID(ctx fiber.Ctx, userID int64) (*float64, error) {
	return r.DAO.GetSumByUserID(ctx.UserContext(), userID)
}

func (r *WithdrawalRepository) GetListByUserID(ctx fiber.Ctx, userID int64) ([]entity.Withdrawal, error) {
	return r.DAO.GetByUserID(ctx.UserContext(), userID)
}

func (r *WithdrawalRepository) Withdraw(ctx fiber.Ctx, userID int64, orderNumber string, amount float64) error {
	return r.DAO.Insert(ctx.UserContext(), userID, orderNumber, amount)
}
