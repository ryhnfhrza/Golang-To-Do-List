package service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
)

type AuthService interface {
	Registration(ctx context.Context, request web.RegistrationRequest) web.AuthResponse
	Login(ctx context.Context, request web.LoginRequest) (web.AuthResponse,*jwt.Token,error)
	
}