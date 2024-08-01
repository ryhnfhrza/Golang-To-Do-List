package web

type RegistrationRequest struct {
	Username        string `validate:"required,min=6,max=30" json:"username"`
	Password        string `validate:"required,min=6,max=64" json:"password"`
	Email           string `validate:"required,email" json:"email"`
	ConfirmPassword string `validate:"required,eqfield=Password" json:"confirm_password"`
}