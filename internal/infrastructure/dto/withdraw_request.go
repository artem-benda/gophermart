package dto

type WithdrawRequest struct {
	OrderNumber string  `json:"order"`
	Amount      float64 `json:"sum"`
}
