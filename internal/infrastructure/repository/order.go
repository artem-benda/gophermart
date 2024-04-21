package repository

import (
	"context"
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type OrderRepository struct {
	DAO dao.Order
}

func (r *OrderRepository) Upload(ctx fiber.Ctx, userID int64, orderNumber string) error {
	err := r.DAO.Insert(ctx, userID, orderNumber)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.IntegrityConstraintViolation {
		return contract.ErrOrderAlreadyUploaded
	}
	return err
}

func (r *OrderRepository) GetByUserID(ctx fiber.Ctx, userID int64) ([]entity.Order, error) {
	return r.DAO.GetByUserID(ctx, userID)
}

func (r *OrderRepository) GetListToSyncAccruals(ctx context.Context) ([]entity.Order, error) {
	return r.DAO.FindByStatuses(ctx, entity.OrderStatusNew, entity.OrderStatusProcessing)
}
