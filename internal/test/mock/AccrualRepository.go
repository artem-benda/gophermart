package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type AccrualRepository struct {
	mock.Mock
}

func (r *AccrualRepository) SyncOrderAccrual(ctx context.Context, orderNumber string) error {
	args := r.Called(ctx, orderNumber)
	return args.Error(0)
}
