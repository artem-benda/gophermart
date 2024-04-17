package handler

import (
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type loginUser struct {
	svc service.User
	v   *validator.Validate
}

func NewLoginHandler(svc *service.User, v *validator.Validate) func(c fiber.Ctx) error {
	controller := registerUser{svc, v}
	return controller.registerUser
}

func (h loginUser) login(ctx fiber.Ctx, login string, password string) error {
	ctx.Accepts("application/json")

	loginDTO := new(dto.LoginRequest)
	var err error

	b := ctx.Bind()
	err = b.JSON(loginDTO)
	if err != nil {
		return err
	}

	err = h.v.Struct(loginDTO)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return nil
	}

	err = h.svc.Login(ctx, login, password)
	if errors.Is(service.ErrUserNotFound, err) || errors.Is(service.ErrUnauthorized, err) {
		return fiber.ErrUnauthorized
	}
	return err
}
