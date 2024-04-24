package worker

import (
	"context"
	"github.com/artem-benda/gophermart/internal/domain/service"
)

type accrualWorker struct {
	svc *service.Accrual
	ctx context.Context
}

func NewAccrualWorkerFunc(svc *service.Accrual, ctx context.Context) func() {
	return accrualWorker{svc: svc, ctx: ctx}.SynchronizeOrderAccruals
}

func (w accrualWorker) SynchronizeOrderAccruals() {
	w.svc.SyncOrdersAccrualsWorker(w.ctx)
}
