package handler

import (
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/gofiber/fiber/v3"
)

type RegisterUser struct {
	svc *service.User
}

func NewRegisterUserHandler(svc *service.User) func(c fiber.Ctx) error {
	controller := RegisterUser{svc}
	return controller.RegisterUser
}

func (h RegisterUser) RegisterUser(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	userDTO := new(dto.RegisterUserRequest)
	var err error

	b := ctx.Bind()
	err = b.JSON(userDTO)
	if err != nil {
		return err
	}

	err = h.svc.Register(ctx, userDTO.Login, userDTO.Password)
	if err != nil {
		return err
	}

	return nil
}
