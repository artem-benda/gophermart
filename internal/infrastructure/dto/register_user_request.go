package dto

type RegisterUserRequest struct {
	Login    string `json:"login" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}
