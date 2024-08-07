package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/app"
	controller "github.com/ryhnfhrza/Golang-To-Do-List-API/controller/AuthController"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/exception"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	repository "github.com/ryhnfhrza/Golang-To-Do-List-API/repository/AuthRepository"
	service "github.com/ryhnfhrza/Golang-To-Do-List-API/service/AuthService"
)

func main() {
	validate := validator.New()
	db := app.NewDb()

	//Auth
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository , db ,validate )
	authController := controller.NewAuthController(authService)

	
	
	router := mux.NewRouter()

	//loginForm
	router.HandleFunc("/todolist/registration",authController.Registration).Methods("POST")
	router.HandleFunc("/todolist/login",authController.Login).Methods("POST")
	router.HandleFunc("/todolist/logout", authController.Logout).Methods("GET")

	/*todolist := router.PathPrefix("/mytodolist").Subrouter()
	todolist.HandleFunc("/create",).Methods("")*/

	router.Use(exception.ErrorHandler)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
