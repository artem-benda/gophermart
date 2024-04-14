package middleware

import (
	"github.com/artem-benda/gophermart/internal/application/jwt"
	"github.com/gofiber/fiber/v3"
	"strings"
)

func NewAuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeaders := c.GetReqHeaders()["Authorization"]

		if len(authHeaders) == 0 {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}

		authHeader := authHeaders[0]
		if authHeader == "" {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		reqToken := splitToken[1]

		userID := jwt.GetUserID(reqToken)
		if userID == -1 {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}

		c.Context().SetUserValue("userID", userID)
		return c.Next()
	}
}
