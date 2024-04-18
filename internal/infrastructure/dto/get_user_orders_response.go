package dto

import (
	"github.com/artem-benda/gophermart/internal/domain/entity"
	"time"
)

type GetUserOrdersResponse []GetUserOrdersItem

type GetUserOrdersItem struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    *float64  `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func MapToDTO(c []entity.Order) GetUserOrdersResponse {
	res := make(GetUserOrdersResponse, 0)
	for _, item := range c {
		res = append(res, GetUserOrdersItem{Number: item.Number, Status: string(item.Status), Accrual: item.AccrualAmount, UploadedAt: item.UploadedAt})
	}
	return res
}
