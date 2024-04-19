package handler

import (
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type withdraw struct {
	svc      *service.Withdrawal
	validate *validator.Validate
}

func NewWithdrawHandler(svc *service.Withdrawal, validate *validator.Validate) func(c fiber.Ctx) error {
	controller := withdraw{svc, validate}
	return controller.withdraw
}

func (h withdraw) withdraw(ctx fiber.Ctx) error {
	userID := middleware.GetUserID(ctx)

	withdrawDTO := new(dto.WithdrawRequest)
	var err error

	b := ctx.Bind()
	err = b.JSON(withdrawDTO)
	if err != nil {
		return err
	}

	err = h.validate.Struct(withdrawDTO)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return nil
	}

	return h.svc.Withdraw(ctx, userID, withdrawDTO.OrderNumber, withdrawDTO.Amount)
}
