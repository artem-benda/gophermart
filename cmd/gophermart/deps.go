package main

import (
	"github.com/artem-benda/gophermart/internal/domain/service"
	"github.com/artem-benda/gophermart/internal/infrastructure/api"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
	"github.com/artem-benda/gophermart/internal/infrastructure/repository"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppDependencies struct {
	AccrualService    *service.Accrual
	OrderService      *service.Order
	UserService       *service.User
	WithdrawalService *service.Withdrawal
}

func mustSetupAppDependencies(dbPool *pgxpool.Pool, cfg Config) *AppDependencies {
	withdrawalDAO := dao.Withdrawal{DB: dbPool}
	withdrawalRepository := repository.WithdrawalRepository{DAO: withdrawalDAO}
	withdrawalService := service.Withdrawal{WithdrawalRepository: &withdrawalRepository}

	userDAO := dao.User{DB: dbPool}
	userRepository := repository.UserRepository{DAO: userDAO}
	userService := service.User{UserRepository: &userRepository, WithdrawalRepository: &withdrawalRepository, Salt: cfg.mustGetSalt()}

	orderDAO := dao.Order{DB: dbPool}
	orderRepository := repository.OrderRepository{DAO: orderDAO}
	orderService := service.Order{OrderRepository: &orderRepository}

	apiClient := resty.New()
	apiClient.SetBaseURL(cfg.AccrualEndpoint)
	apiClient.SetLogger(log.DefaultLogger())
	accrualAPI := api.AccrualAPI{Client: apiClient}

	accrualRepository := repository.AccrualRepository{DAO: orderDAO, API: accrualAPI}
	accrualService := service.Accrual{OrdersRepo: &orderRepository, AccrualRepo: &accrualRepository}

	return &AppDependencies{
		AccrualService:    &accrualService,
		OrderService:      &orderService,
		UserService:       &userService,
		WithdrawalService: &withdrawalService,
	}
}
