package handler

import (
	"errors"
	"fmt"
	"github.com/artem-benda/gophermart/internal/application/jwt"
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
)

type registerUser struct {
	svc      *service.User
	validate *validator.Validate
}

func NewRegisterUserHandler(svc *service.User, v *validator.Validate) func(c fiber.Ctx) error {
	log.Debug("NewRegisterUserHandler...")
	controller := registerUser{svc, v}
	return controller.registerUser
}

func (h registerUser) registerUser(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	userDTO := new(dto.RegisterUserRequest)
	var err error

	b := ctx.Bind()
	err = b.JSON(userDTO)
	if err != nil {
		return err
	}

	err = h.validate.Struct(userDTO)
	if err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return nil
	}

	userID, err := h.svc.Register(ctx, userDTO.Login, userDTO.Password)

	if errors.Is(err, contract.ErrUserAlreadyRegistered) {
		ctx.Response().SetStatusCode(http.StatusConflict)
		return nil
	}

	if err != nil {
		log.Debug("unexpected error", ":", err)
		return fiber.ErrInternalServerError
	}

	log.Debug("registered user with id: ", *userID)

	token, err := jwt.BuildJWTString(*userID)

	if err != nil {
		return err
	}

	ctx.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return nil
}
