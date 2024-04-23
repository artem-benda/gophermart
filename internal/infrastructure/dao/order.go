package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Order struct {
	DB *pgxpool.Pool
}

func (dao Order) Insert(ctx fiber.Ctx, userID int64, orderNumber string) error {
	_, err := dao.DB.Exec(ctx.UserContext(), "insert into user_orders(order_number, user_id, uploaded_at, status) values($1, $2, $3, $4)", orderNumber, userID, time.Now(), entity.OrderStatusNew)
	if err != nil {
		return err
	}
	return nil
}

func (dao Order) GetByUserID(ctx fiber.Ctx, userID int64) ([]entity.Order, error) {
	rows, err := dao.DB.Query(ctx.UserContext(), "SELECT order_number, user_id, uploaded_at, status, accrual_amount FROM user_orders WHERE user_id = $1 ORDER BY uploaded_at", userID)
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Order, 0)

	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(&order.Number, &order.UserID, &order.UploadedAt, &order.Status, &order.AccrualAmount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return orders, nil
}

func (dao Order) GetByOrderNumber(ctx fiber.Ctx, orderNumber string) (*entity.Order, error) {
	row := dao.DB.QueryRow(ctx.UserContext(), "SELECT order_number, user_id, uploaded_at, status, accrual_amount FROM user_orders WHERE order_number = $1", orderNumber)

	order := entity.Order{}
	err := row.Scan(&order.Number, &order.UserID, &order.UploadedAt, &order.Status, &order.AccrualAmount)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (dao Order) UpdateOrder(ctx context.Context, orderNumber string, accrual *float64, status entity.OrderStatus) error {
	var accrualNullable sql.NullFloat64
	if accrual != nil {
		accrualNullable = sql.NullFloat64{Float64: *accrual, Valid: true}
	} else {
		accrualNullable = sql.NullFloat64{Valid: false}
	}
	_, err := dao.DB.Exec(ctx, "update user_orders SET accrual_amount = $1, status = $2 WHERE order_number = $2", accrualNullable, string(status), orderNumber)
	if err != nil {
		return err
	}
	return nil
}

func (dao Order) FindByStatuses(ctx context.Context, statuses ...entity.OrderStatus) ([]entity.Order, error) {
	statusesStr := make([]string, 0)
	for _, status := range statuses {
		statusesStr = append(statusesStr, string(status))
	}
	rows, err := dao.DB.Query(ctx, "SELECT order_number, user_id, uploaded_at, status, accrual_amount FROM user_orders WHERE status = any($1) ORDER BY uploaded_at", statusesStr)
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Order, 0)

	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(&order.Number, &order.UserID, &order.UploadedAt, &order.Status, &order.AccrualAmount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return orders, nil
}
