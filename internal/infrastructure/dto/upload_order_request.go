package dto

type UploadOrderRequest struct {
	OrderNumber string `validate:"luhn"`
}
