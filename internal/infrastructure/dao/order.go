package dao

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Order struct {
	DB *pgxpool.Pool
}

func (dao Order) Insert(ctx fiber.Ctx, userID int64, orderNumber string) error {
	_, err := dao.DB.Exec(ctx.UserContext(), "insert into user_orders(order_number, user_id, uploaded_at, status) values($1, $2)", orderNumber, userID, time.Now(), entity.OrderStatusNew)
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
