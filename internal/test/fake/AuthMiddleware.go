package fake

import (
	"errors"
	"github.com/artem-benda/gophermart/internal/application/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"strconv"
	"strings"
)

func NewAuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeaders := c.GetReqHeaders()["Authorization"]

		if len(authHeaders) == 0 {
			c.Status(fiber.StatusUnauthorized)
			log.Debug("authHeaders len == 0 -> Unauthorized")
			return nil
		}

		authHeader := authHeaders[0]
		if authHeader == "" {
			c.Status(fiber.StatusUnauthorized)
			log.Debug("authHeader is empty -> Unauthorized")
			return nil
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		reqToken := splitToken[1]

		userID, err := strconv.ParseInt(reqToken, 10, 64)
		if err != nil {
			panic(errors.New("invalid value, use userID as token value, for ex. 'Bearer: 1'"))
		}

		middleware.SetUserID(c, userID)

		return c.Next()
	}
}
