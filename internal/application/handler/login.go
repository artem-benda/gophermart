package handler

import (
	"errors"
	"fmt"
	"github.com/artem-benda/gophermart/internal/application/jwt"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
)

type loginUser struct {
	svc *service.User
	v   *validator.Validate
}

func NewLoginHandler(svc *service.User, v *validator.Validate) func(c fiber.Ctx) error {
	controller := loginUser{svc, v}
	return controller.login
}

func (h loginUser) login(ctx fiber.Ctx) error {
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

	id, err := h.svc.Login(ctx, loginDTO.Login, loginDTO.Password)
	if errors.Is(service.ErrUserNotFound, err) || errors.Is(service.ErrUnauthorized, err) {
		return fiber.ErrUnauthorized
	}

	if err != nil {
		ctx.Response().SetStatusCode(http.StatusInternalServerError)
	}

	log.Debug("logging in user with id: ", *id)

	token, err := jwt.BuildJWTString(*id)

	if err != nil {
		return err
	}

	ctx.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return nil
}
