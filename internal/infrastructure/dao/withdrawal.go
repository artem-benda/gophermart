package dao

import (
	"database/sql"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Withdrawal struct {
	DB *pgxpool.Pool
}

func (dao Withdrawal) Insert(ctx fiber.Ctx, userID int64, orderNumber string, amount float64) error {
	tx, err := dao.DB.Begin(ctx.UserContext())
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx.UserContext())
	}()
	_, err = tx.Exec(ctx.UserContext(), "insert into order_withdrawals(order_number, user_id, amount, created_at, processed_at) values($1, $2, $3, $4, $5)", orderNumber, userID, amount, time.Now(), time.Now())
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx.UserContext(), "update users SET points_balance = points_balance + $1 WHERE id = $2", amount, userID)
	if err != nil {
		return err
	}
	return tx.Commit(ctx.UserContext())
}

func (dao Withdrawal) GetSumByUserID(ctx fiber.Ctx, userID int64) (*float64, error) {
	var sum sql.NullFloat64
	row := dao.DB.QueryRow(ctx.UserContext(), "select SUM(amount) FROM order_withdrawals WHERE user_id = $1", userID)
	err := row.Scan(&sum)
	if err != nil {
		return nil, err
	}
	if !sum.Valid {
		return nil, nil
	}
	return &sum.Float64, nil
}

func (dao Withdrawal) GetByUserID(ctx fiber.Ctx, userID int64) ([]entity.Withdrawal, error) {
	rows, err := dao.DB.Query(ctx.UserContext(), "SELECT order_number, user_id, amount, created_at, processed_at FROM order_withdrawals WHERE user_id = $1 ORDER BY processed_at", userID)
	if err != nil {
		return nil, err
	}

	withdrawals := make([]entity.Withdrawal, 0)

	for rows.Next() {
		withdrawal := entity.Withdrawal{}
		err := rows.Scan(&withdrawal.OrderNumber, &withdrawal.UserID, &withdrawal.Amount, &withdrawal.CreatedAt, &withdrawal.ProcessedAt)
		if err != nil {
			return nil, err
		}
		withdrawals = append(withdrawals, withdrawal)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return withdrawals, nil
}
