package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
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

	helper.RunMigrations(db)

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
	todolist.HandleFunc("/list-tasks",tasksController.FindAllTask).Methods("GET")
	todolist.HandleFunc("/search/{keyword}",tasksController.SearchTask).Methods("GET")
	todolist.HandleFunc("/completed/{taskId}",tasksController.ComplatedTask).Methods("PATCH")
	
	router.Use(exception.ErrorHandler) 

	c := cron.New()
    _, errCron := c.AddFunc("@every 1m", func() {
        ctx := context.Background()
        if err := tasksService.SendDueDateReminders(ctx); err != nil {
            log.Println("Error sending due date reminders:", err)
        }
    })
    if errCron != nil {
        log.Fatalf("Error setting up cron job: %v", errCron)
    }
    c.Start()

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}


