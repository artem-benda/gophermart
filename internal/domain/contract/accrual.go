package contract

import "context"

type AccrualRepository interface {
	SyncOrderAccrual(ctx context.Context, orderNumber string) error
}
