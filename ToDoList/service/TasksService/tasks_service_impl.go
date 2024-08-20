package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/exception"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
	repository "github.com/ryhnfhrza/Golang-To-Do-List-API/repository/TasksRepository"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/util"
)

type TasksServiceImpl struct {
	TaskRepository repository.TasksRepository
	Db                  *sql.DB
	validate            *validator.Validate
}

func NewTasksService(taskRepository repository.TasksRepository, db *sql.DB, Validate *validator.Validate) TasksService {
	return &TasksServiceImpl{
		TaskRepository: taskRepository,
		Db: db,
		validate:        Validate,
	}
}

func(service *TasksServiceImpl)CreateTask(ctx context.Context, request web.CreateTaskRequest) web.TaskResponse{
	err := service.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	tokenString, ok := ctx.Value(util.TokenKey).(string)
	if !ok {
		helper.PanicIfError(errors.New("token not found in context"))
	}

	claims := &util.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return util.JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		panic(exception.NewUnauthorizedError("invalid token"))
	}
	
	// get ID from klaim JWT as UserID
	idUser := claims.ID
	username:= claims.Username
	
	// id tasks maker
	id := uuid.New()
	idStr := id.String()
	idStrNoHyphens := strings.ReplaceAll(idStr, "-", "")
	
	// handle optional field
	var dueDate sql.NullTime
	if request.DueDate != "" {
			defaultTime := " 23:59:59"
			if !strings.Contains(request.DueDate, ":") {
					request.DueDate += defaultTime
			}

			formattedDate := "2006-01-02 15:04:05"
			parsedDate, err := time.Parse(formattedDate, request.DueDate)
			if err != nil {
					errorMessage := "Invalid due_date format: " + request.DueDate+ " . Expected format: YYYY-MM-DD or YYYY-MM-DD HH:MM:SS"
					panic(exception.NewBadRequestError(errorMessage))
			}

			dueDate = sql.NullTime{Time: parsedDate, Valid: true}
	} else {
			dueDate = sql.NullTime{Valid: false}
	}

		
	description := request.Description
	if description == "" {
		description = "No description provided"
	}

	

	task := domain.Tasks{
		IdTasks: idStrNoHyphens, 
    UserId: idUser,
    Title: request.Title,
    Description: description,
    Completed: 0,
    DueDate: dueDate,
    Notified: 0,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
	}

	defer func() {
		if r := recover(); r != nil {
			if sqlErr, ok := r.(*mysql.MySQLError); ok {
				if sqlErr.Number == 1644 { 
					panic(exception.NewBadRequestError(sqlErr.Message))
				}
			}
			panic(r)
		}
	}()

	task = service.TaskRepository.CreateTask(ctx,tx,task)

	var completedStatus string
	if task.Completed == 1{
		completedStatus = "Completed"
	}else{
		completedStatus = "Pending"
	}

	return helper.ToTasksResponse(task,username,completedStatus)
}