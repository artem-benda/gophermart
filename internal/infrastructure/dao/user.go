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
	user := new(entity.User)

	row := dao.DB.QueryRow(ctx.UserContext(), "SELECT * FROM users WHERE login = $1", login)
	err := row.Scan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dao User) Insert(ctx fiber.Ctx, user entity.User) (*int64, error) {
	userId := new(int64)
	row := dao.DB.QueryRow(ctx.UserContext(), "insert into users(login, password_hash) values($1, $2) returning id", user.Login, user.PasswordHash)
	err := row.Scan(userId)
	if err != nil {
		return nil, err
	}
	return userId, nil
}
