package service

import (
	"context"
	"github.com/artem-benda/gophermart/internal/domain/contract"
	"github.com/gofiber/fiber/v3/log"
	"time"
)

type Accrual struct {
	OrdersRepo  contract.OrderRepository
	AccrualRepo contract.AccrualRepository
}

func (s Accrual) SyncOrdersAccrualsWorker(ctx context.Context) {
mainLoop:
	for {
		orders, err := s.OrdersRepo.GetListToSyncAccruals(ctx)
		if err != nil {
			log.Debug("could not get orders to sync", ":", err)
			time.Sleep(10 * time.Second)
			continue mainLoop
		}

		for _, order := range orders {
			err := s.AccrualRepo.SyncOrderAccrual(ctx, order.Number)
			if err != nil {
				log.Debug("could not get sync order", ":", order.Number)
				time.Sleep(10 * time.Second)
				continue mainLoop
			}
		}

		time.Sleep(time.Second)
	}
}
