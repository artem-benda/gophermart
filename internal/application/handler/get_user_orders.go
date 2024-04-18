package handler

import (
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/gofiber/fiber/v3"
)

type getUserOrders struct {
	svc *service.Order
}

func NewGetUserOrdersHandler(svc *service.Order) func(c fiber.Ctx) error {
	controller := getUserOrders{svc}
	return controller.getList
}

func (h getUserOrders) getList(ctx fiber.Ctx) error {
	userID := middleware.GetUserID(ctx)

	orders, err := h.svc.GetAll(ctx, userID)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
