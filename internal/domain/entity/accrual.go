package entity

type AccrualStatus string

const (
	AccrualStatusRegistered AccrualStatus = "REGISTERED"
	AccrualStatusProcessing AccrualStatus = "PROCESSING"
	AccrualStatusInvalid    AccrualStatus = "INVALID"
	AccrualStatusProcessed  AccrualStatus = "PROCESSED"
)

type Accrual struct {
	OrderNumber   string
	Status        AccrualStatus
	AccrualAmount *float64
}
