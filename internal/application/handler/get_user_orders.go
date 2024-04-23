package handler

import (
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type getUserOrders struct {
	svc *service.Order
}

func NewGetUserOrdersHandler(svc *service.Order) func(c fiber.Ctx) error {
	log.Debug("NewGetUserOrdersHandler...")
	controller := getUserOrders{svc}
	return controller.getList
}

func (h getUserOrders) getList(ctx fiber.Ctx) error {
	userID := middleware.GetUserID(ctx)

	orders, err := h.svc.GetAll(ctx, userID)

	if err != nil {
		return err
	}

	err = ctx.JSON(dto.MapOrdersToDTO(orders))

	if err != nil {
		return err
	}

	return nil
}
