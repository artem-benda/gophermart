package handler

import (
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type getUserBalance struct {
	svc *service.User
}

func NewGetUserBalanceHandler(svc *service.User) func(c fiber.Ctx) error {
	log.Debug("NewGetUserBalanceHandler...")
	controller := getUserBalance{svc}
	return controller.get
}

func (h getUserBalance) get(ctx fiber.Ctx) error {
	userID := middleware.GetUserID(ctx)

	user, err := h.svc.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	totalWithdrawals, err := h.svc.GetTotalWithdrawals(ctx, userID)
	if err != nil {
		return err
	}

	err = ctx.JSON(dto.GetUserBalanceResponse{Current: user.PointsBalance, Withdrawn: totalWithdrawals})
	if err != nil {
		return err
	}

	return nil
}
