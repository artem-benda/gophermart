package handler

import (
	"errors"
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
)

type uploadOrder struct {
	svc      *service.Order
	validate *validator.Validate
}

func NewUploadOrderHandler(svc *service.Order, v *validator.Validate) func(c fiber.Ctx) error {
	controller := uploadOrder{svc, v}
	return controller.upload
}

func (h uploadOrder) upload(ctx fiber.Ctx) error {
	ctx.Accepts("text/plain")

	request := dto.UploadOrderRequest{OrderNumber: string(ctx.Body())}
	err := h.validate.Struct(request)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return nil
	}

	userID := middleware.GetUserID(ctx)

	err = h.svc.Upload(ctx, userID, request.OrderNumber)
	if errors.Is(err, contract.ErrOrderAlreadyUploaded) {
		ctx.Response().SetStatusCode(http.StatusOK)
		return nil
	}
	if errors.Is(err, contract.ErrOrderUploadedByAnotherUser) {
		ctx.Response().SetStatusCode(http.StatusConflict)
		return nil
	}
	if err != nil {
		log.Debug(err)
		return fiber.ErrInternalServerError
	}

	ctx.Response().SetStatusCode(http.StatusAccepted)
	return nil
}
