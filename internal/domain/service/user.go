package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	UserRepository       contract.UserRepository
	WithdrawalRepository contract.WithdrawalRepository
	Salt                 []byte
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUnauthorized = errors.New("unauthorized")
)

func (s User) Register(ctx fiber.Ctx, login string, password string) (*int64, error) {
	passwordHash, err := computeHash(password, s.Salt)

	if err != nil {
		return nil, err
	}

	return s.UserRepository.Register(ctx, login, *passwordHash)
}

func (s User) Login(ctx fiber.Ctx, login string, password string) (*int64, error) {
	passwordHashString, err := computeHash(password, s.Salt)
	if err != nil {
		return nil, err
	}
	user, err := s.UserRepository.GetUserByLogin(ctx, login)

	if err != nil {
		log.Debug(err)
		return nil, err
	}

	if user == nil {
		log.Debug("user not found")
		return nil, ErrUserNotFound
	}

	if user.PasswordHash != *passwordHashString {
		log.Debug("hash mismatch")
		return nil, ErrUnauthorized
	}

	return &user.ID, nil
}

func (s User) GetUserByID(ctx fiber.Ctx, userID int64) (*entity.User, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s User) GetTotalWithdrawals(ctx fiber.Ctx, userID int64) (*float64, error) {
	return s.WithdrawalRepository.GetTotalByUserID(ctx, userID)
}

func computeHash(password string, salt []byte) (*string, error) {
	pwPbkdf2 := pbkdf2.Key([]byte(password), salt, 10240, 32, sha256.New)
	encodedHash := hex.EncodeToString(pwPbkdf2)

	return &encodedHash, nil
}
