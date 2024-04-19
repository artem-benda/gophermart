package main

import (
	"github.com/artem-benda/gophermart/internal/infrastructure/dto"
	"github.com/go-playground/validator/v10"
)

func mustCreateValidator() *validator.Validate {
	validate := validator.New()
	err := validate.RegisterValidation("luhn", dto.LuhnStringValidator)
	if err != nil {
		panic(err)
	}
	return validate
}
