package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
	repository "github.com/ryhnfhrza/Golang-To-Do-List-API/repository/AuthRepository"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/util"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Db                  *sql.DB
	validate            *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, db *sql.DB, Validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		Db: db,
		validate:        Validate,
	}
}

func(service *AuthServiceImpl)Registration(ctx context.Context, request web.RegistrationRequest) web.AuthResponse{
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//id maker 
	id := uuid.New()
	idStr := id.String()
	idStrNoHyphens := strings.ReplaceAll(idStr, "-", "")
	

	email,err := service.AuthRepository.CheckEmail(ctx,tx,request.Email)
	helper.PanicIfError(err)

	//hashing password handle
	hashPassword,err := util.HashPassword(request.Password)
	helper.PanicIfError(err)

	user := domain.Users{
		Id: idStrNoHyphens,
		Username: request.Username,
		Email: email,
		Password: hashPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user = service.AuthRepository.Registration(ctx,tx,user)

	return helper.ToAuthResponse(user)
}

func(service *AuthServiceImpl)Login(ctx context.Context, request web.LoginRequest) (web.AuthResponse,*jwt.Token){
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	
	
	login,err := service.LoginFormRepository.Login(ctx,tx,request.Username)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	
	err = bcrypt.CompareHashAndPassword([]byte(login.Password),[]byte(request.Password))
	if err != nil{
		panic(exception.NewUnauthorizedError("Password is incorrect"))
	}
	
	expTime := time.Now().Add(time.Hour * 1)
	claims := &util.JWTClaim{
		Username: request.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "github.com/ryhnfhrza",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	
	
	return helper.ToLoginFormResponse(login),tokenAlgo
}

func(service *AuthServiceImpl)Logout(ctx context.Context){

}
