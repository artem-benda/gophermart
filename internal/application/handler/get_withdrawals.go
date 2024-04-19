package handler

import (
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/gofiber/fiber/v3"
)

type getWithdrawals struct {
	svc *service.Withdrawal
}

func NewGetWithdrawalsHandler(svc *service.Withdrawal) func(c fiber.Ctx) error {
	controller := getWithdrawals{svc}
	return controller.getList
}

func (h getWithdrawals) getList(ctx fiber.Ctx) error {
	userID := middleware.GetUserID(ctx)

	withdrawals, err := h.svc.GetList(ctx, userID)

	if err != nil {
		return err
	}

	err = ctx.JSON(dto.MapWithdrawalsToDTO(withdrawals))

	if err != nil {
		return err
	}

	return nil
}
