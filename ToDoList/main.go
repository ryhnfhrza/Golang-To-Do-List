package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/app"
	controllerAuth "github.com/ryhnfhrza/Golang-To-Do-List-API/controller/AuthController"
	controllerTasks "github.com/ryhnfhrza/Golang-To-Do-List-API/controller/TasksController"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/exception"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/middlewares"
	repositoryAuth "github.com/ryhnfhrza/Golang-To-Do-List-API/repository/AuthRepository"
	repositoryTasks "github.com/ryhnfhrza/Golang-To-Do-List-API/repository/TasksRepository"
	serviceAuth "github.com/ryhnfhrza/Golang-To-Do-List-API/service/AuthService"
	serviceTasks "github.com/ryhnfhrza/Golang-To-Do-List-API/service/TasksService"
)

func main() {
	validate := validator.New()
	db := app.NewDb()

	//Auth
	authRepository := repositoryAuth.NewAuthRepository()
	authService := serviceAuth.NewAuthService(authRepository , db ,validate )
	authController := controllerAuth.NewAuthController(authService)

	//Tasks
	tasksRepository := repositoryTasks.NewTasksRepository()
	tasksService := serviceTasks.NewTasksService(tasksRepository,db,validate)
	tasksController := controllerTasks.NewTasksController(tasksService)
	
	
	router := mux.NewRouter()

	//loginForm
	router.HandleFunc("/todolist/registration",authController.Registration).Methods("POST")
	router.HandleFunc("/todolist/login",authController.Login).Methods("POST")
	router.HandleFunc("/todolist/logout", authController.Logout).Methods("GET")

	//task
	todolist := router.PathPrefix("/mytodolist").Subrouter()
	todolist.Use(middlewares.JWTMiddleware)
	todolist.HandleFunc("/create",tasksController.CreateTask).Methods("POST")
	todolist.HandleFunc("/update/{taskId}",tasksController.UpdateTask).Methods("PATCH")
	todolist.HandleFunc("/delete/{taskId}",tasksController.DeleteTask).Methods("DELETE")

	router.Use(exception.ErrorHandler) 

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

