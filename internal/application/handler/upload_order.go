package handler

import (
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/gofiber/fiber/v3"
)

type uploadOrder struct {
	svc *service.Order
}

func NewUploadOrderHandler(svc *service.Order) func(c fiber.Ctx) error {
	controller := uploadOrder{svc}
	return controller.upload
}

func (h uploadOrder) upload(ctx fiber.Ctx) error {
	ctx.Accepts("text/plain")

	request := dto.UploadOrderRequest{OrderID: dto.LuhnNumber(ctx.Body())}
	if !request.Valid() {
		return fiber.ErrBadRequest
	}

	userID := middleware.GetUserID(ctx)

	err := h.svc.Upload(ctx, userID, request.OrderNumberString())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
