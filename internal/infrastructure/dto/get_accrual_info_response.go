package dto

type GetAccrualInfoResponse struct {
	Number  string   `json:"number"`
	Status  string   `json:"status"`
	Accrual *float64 `json:"accrual,omitempty"`
}
