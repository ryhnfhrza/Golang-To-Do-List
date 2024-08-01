package web

type AuthResponse struct {
	Id       string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}