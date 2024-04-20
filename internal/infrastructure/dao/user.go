package dao

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	DB *pgxpool.Pool
}

func (dao User) GetByLogin(ctx fiber.Ctx, login string) (*entity.User, error) {
	user := entity.User{}

	row := dao.DB.QueryRow(ctx.UserContext(), "SELECT id, login, password_hash, points_balance FROM users WHERE login = $1", login)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.PointsBalance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao User) GetByID(ctx fiber.Ctx, userID int64) (*entity.User, error) {
	user := entity.User{}

	row := dao.DB.QueryRow(ctx.UserContext(), "SELECT id, login, password_hash, points_balance FROM users WHERE id = $1", userID)
	err := row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.PointsBalance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao User) Insert(ctx fiber.Ctx, user entity.User) (*int64, error) {
	userID := new(int64)
	row := dao.DB.QueryRow(ctx.UserContext(), "insert into users(login, password_hash) values($1, $2) returning id", user.Login, user.PasswordHash)
	err := row.Scan(userID)
	if err != nil {
		return nil, err
	}
	return userID, nil
}
