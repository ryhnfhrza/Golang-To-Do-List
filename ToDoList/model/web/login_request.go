package web

type LoginRequest struct {
	Username string `validate:"required,min=4,max=30" json:"username"`
	Password string `validate:"required,min=4,max=64" json:"password"`
}