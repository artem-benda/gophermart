package entity

type User struct {
	ID            int64
	Login         string
	PasswordHash  string
	PointsBalance float64
}
