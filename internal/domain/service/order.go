package service

import (
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
)

type Order struct {
	OrderRepository contract.OrderRepository
}

func (s Order) Upload(ctx fiber.Ctx, userID int64, orderNumber string) error {
	return s.OrderRepository.Upload(ctx, userID, orderNumber)
}

func (s Order) GetAll(ctx fiber.Ctx, userID int64) ([]entity.Order, error) {
	return s.OrderRepository.GetByUserID(ctx, userID)
}
