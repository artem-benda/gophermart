package contract

import "github.com/artem-benda/gophermart/internal/domain/entity"

type AccrualRepository interface {
	GetByOrderNumber(orderNumber string) (accrual entity.Accrual, err error)
}
