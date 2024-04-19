package service

import (
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/gofiber/fiber/v3"
)

type balance struct {
	UserRepository       contract.UserRepository
	WithdrawalRepository contract.WithdrawalRepository
}

func (s balance) GetBalance(ctx fiber.Ctx, userID int64) (*float64, error) {
	user, err := s.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &user.PointsBalance, nil
}

func (s balance) GetWithdrawalsTotal(ctx fiber.Ctx, userID int64) (*float64, error) {
	withdrawals, err := s.WithdrawalRepository.GetTotalByUserID(ctx, userID)
	if err != nil {
		return err
	}
}
