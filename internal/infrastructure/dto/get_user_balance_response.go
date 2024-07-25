package dto

type GetUserBalanceResponse struct {
	Current   float64  `json:"current"`
	Withdrawn *float64 `json:"withdrawn,omitempty"`
}
