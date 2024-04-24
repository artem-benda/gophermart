package dto

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"time"
)

type GetWithdrawalsResponse []GetWithdrawalsItem

type GetWithdrawalsItem struct {
	OrderNumber string    `json:"order"`
	Sum         float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

func MapWithdrawalsToDTO(withdrawals []entity.Withdrawal) GetWithdrawalsResponse {
	res := make(GetWithdrawalsResponse, 0)
	for _, item := range withdrawals {
		res = append(res, GetWithdrawalsItem{OrderNumber: item.OrderNumber, Sum: item.Amount, ProcessedAt: item.ProcessedAt})
	}
	return res
}
