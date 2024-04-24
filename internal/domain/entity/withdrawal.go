package entity

import "time"

type Withdrawal struct {
	OrderNumber string
	UserID      int64
	Amount      float64
	CreatedAt   time.Time
	ProcessedAt time.Time
}
