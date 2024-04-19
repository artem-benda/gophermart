package repository

import (
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	DAO dao.User
}

func (r *UserRepository) Register(ctx fiber.Ctx, login string, passwordHash string) (*int64, error) {
	id, err := r.DAO.Insert(ctx, entity.User{Login: login, PasswordHash: passwordHash})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.IntegrityConstraintViolation {
		return nil, contract.ErrUserAlreadyRegistered
	}
	return id, err
}

func (r *UserRepository) GetUserByLogin(ctx fiber.Ctx, login string) (*entity.User, error) {
	return r.DAO.GetByLogin(ctx, login)
}

func (r *UserRepository) GetUserById(ctx fiber.Ctx, userID int64) (*entity.User, error) {
	return r.DAO.GetByID(ctx, userID)
}
