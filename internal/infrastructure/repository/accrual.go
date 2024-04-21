package repository

import (
	"context"
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"github.com/artem-benda/gophermart/internal/infrastructure/api"
	"github.com/artem-benda/gophermart/internal/infrastructure/dao"
)

type AccrualRepository struct {
	DAO dao.Order
	API api.AccrualAPI
}

func (r *AccrualRepository) SyncOrderAccrual(ctx context.Context, orderNumber string) error {
	accrual, err := r.API.GetAccrualInfo(orderNumber)
	if err != nil {
		return err
	}
	var status entity.OrderStatus
	switch accrual.Status {
	case entity.AccrualStatusRegistered:
		status = entity.OrderStatusNew
	case entity.AccrualStatusInvalid:
		status = entity.OrderStatusInvalid
	case entity.AccrualStatusProcessing:
		status = entity.OrderStatusProcessing
	case entity.AccrualStatusProcessed:
		status = entity.OrderStatusProcessed
	}
	return r.DAO.UpdateOrder(ctx, orderNumber, accrual.AccrualAmount, status)
}
