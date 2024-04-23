package contract

import (
	"context"
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
)

var (
	ErrOrderAlreadyUploaded       = errors.New("order already uploaded")
	ErrOrderUploadedByAnotherUser = errors.New("order uploaded by another user")
)

type OrderRepository interface {
	Upload(ctx fiber.Ctx, userID int64, orderNumber string) error
	GetByUserID(ctx fiber.Ctx, userID int64) ([]entity.Order, error)
	GetListToSyncAccruals(ctx context.Context) ([]entity.Order, error)
}
