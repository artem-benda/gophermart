package handler

import "github.com/gofiber/fiber/v3"

func RegisterUser(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	return nil
}
