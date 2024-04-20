package main

import (
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3/log"
)

func mustCreateValidator() *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("luhn", dto.LuhnStringValidator)
	if err != nil {
		log.Fatal(err)
	}
	return validate
}
