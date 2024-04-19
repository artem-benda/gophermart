package dto

type WithdrawRequest struct {
	OrderNumber string  `json:"order" validate:"luhn"`
	Amount      float64 `json:"sum" validate:"required"`
}
