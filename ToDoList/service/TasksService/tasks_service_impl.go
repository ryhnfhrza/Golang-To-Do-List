package service

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/exception"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
	repository "github.com/ryhnfhrza/Golang-To-Do-List-API/repository/TasksRepository"
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

func(service *TasksServiceImpl)CreateTask(ctx context.Context, request web.CreateTaskRequest) web.UserTasksResponses{
	err := service.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	idUser, username, err := helper.ExtractUserFromToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
	
	idStrNoHyphens := helper.GenerateTaskID()

	
	// handle optional field
	dueDate, err := helper.ParseDueDate(request.DueDate)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	
	description := request.Description
	if description == "" {
		description = "No description provided"
		}	

		//handle created_at and updated at
		now := time.Now()
		loc, err := time.LoadLocation("Asia/Makassar")
    helper.PanicIfError(err)
		witaTime := now.In(loc)
		

	//handle notified
	notifStatus := helper.CalculateNotificationStatus(dueDate)


	task := domain.Tasks{
		IdTasks: idStrNoHyphens, 
    UserId: idUser,
    Title: request.Title,
    Description: description,
    Completed: 0,
    DueDate: dueDate,
    Notified: notifStatus,
    CreatedAt: witaTime,
    UpdatedAt: witaTime,
	}

	defer exception.HandleSQLError()

	task = service.TaskRepository.CreateTask(ctx,tx,task)

	taskResponse := helper.ToTasksResponse(task)
	return web.UserTasksResponses{
		UserName: username,
		Tasks:    []web.TaskResponse{taskResponse},
	}
}

func(service *TasksServiceImpl)UpdateTask(ctx context.Context, request web.UpdateTaskRequest) web.UserTasksResponses{
	err := service.validate.Struct(request)
	helper.PanicIfError(err)
	
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	idUser, username, err := helper.ExtractUserFromToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	task,err := service.TaskRepository.FindTaskById(ctx,tx,request.IdTask,idUser)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	request.Title = helper.GetDefaultIfEmpty(request.Title,task.Title)
	request.Description = helper.GetDefaultIfEmpty(request.Description,task.Description)
	

	// handle optional field
	
	dueDate, err := helper.ParseDueDate(request.DueDate)
	if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
	}
	//handle notified
	notifStatus := helper.CalculateNotificationStatus(dueDate)

	task.DueDate = dueDate
	task.Title = request.Title
	task.Description = request.Description
	task.Notified = notifStatus
	
	defer exception.HandleSQLError()
	
	task = service.TaskRepository.UpdateTask(ctx,tx,task)

	taskResponse := helper.ToTasksResponse(task)
	return web.UserTasksResponses{
		UserName: username,
		Tasks:    []web.TaskResponse{taskResponse},
	}
}

func(service *TasksServiceImpl)DeleteTask(ctx context.Context, taskId string){

	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	idUser, _ , err := helper.ExtractUserFromToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	task,err := service.TaskRepository.FindTaskById(ctx,tx,taskId,idUser)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TaskRepository.DeleteTask(ctx,tx,task)
}

func(service *TasksServiceImpl)FindAllTask(ctx context.Context,sortBy,order string)web.UserTasksResponses{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	
	idUser, username , err := helper.ExtractUserFromToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	if sortBy == "" {
		sortBy = "created_at"
	}
	if order == "" {
			order = "DESC"
	}

	orderReq := strings.ToUpper(order)
	
	validSortBy,validOrder,err :=helper.ValidateSortParams(sortBy,orderReq)
	if err != nil{
		panic(exception.NewBadRequestError(err.Error()))
	}
	

	tasks := service.TaskRepository.FindAllTask(ctx,tx,idUser,validSortBy,validOrder)
	
	return web.UserTasksResponses{
		UserName: username,
		Tasks:    helper.ToTasksResponses(tasks),
	}
}

func(service *TasksServiceImpl)SearchTask(ctx context.Context, keyword,sortBy,order string) web.UserTasksResponses{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	
	idUser, username , err := helper.ExtractUserFromToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	if sortBy == "" {
		sortBy = "created_at"
	}
	if order == "" {
			order = "DESC"
	}

	orderReq := strings.ToUpper(order)
	
	validSortBy,validOrder,err := helper.ValidateSortParams(sortBy,orderReq)
	if err != nil{
		panic(exception.NewBadRequestError(err.Error()))
	}

	tasks,err := service.TaskRepository.SearchTask(ctx,tx,keyword,idUser,validSortBy,validOrder)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UserTasksResponses{
		UserName: username,
		Tasks:    helper.ToTasksResponses(tasks),
	}
}

func(service *TasksServiceImpl)SendDueDateReminders(ctx context.Context)error{
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userTasksDueOneDay := service.TaskRepository.FindTaskDueInOneDay(ctx,tx)
	for _, t := range userTasksDueOneDay {
		
		stringDueDate := t.DueDate.Time.Format(time.RFC3339)
		helper.SendGomail("emailTemplate.html",t.Email,t.Username,t.Title,t.Description,stringDueDate,"1 Day")
		
		taskToUpdate := domain.Tasks{
			IdTasks:   t.IdTasks,
			UserId:    t.UserId,
			Title:     t.Title,
			Description: t.Description,
			Completed: 0,  
			DueDate:   t.DueDate,
			Notified:  1,  
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}

		err := service.TaskRepository.UpdateTaskAfterNotification(ctx,tx,taskToUpdate)
		helper.PanicIfError(err)
	}

	userTasksDueOneHour := service.TaskRepository.FindTaskDueInOneHour(ctx,tx)
	for _, t := range userTasksDueOneHour {
		
		stringDueDate := t.DueDate.Time.Format(time.RFC3339)
		helper.SendGomail("emailTemplate.html",t.Email,t.Username,t.Title,t.Description,stringDueDate,"1 hour")
		
		taskToUpdate := domain.Tasks{
			IdTasks:   t.IdTasks,
			UserId:    t.UserId,
			Title:     t.Title,
			Description: t.Description,
			Completed: 0,  
			DueDate:   t.DueDate,
			Notified:  2,  
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}

		err := service.TaskRepository.UpdateTaskAfterNotification(ctx,tx,taskToUpdate)
		helper.PanicIfError(err)
	}
	return nil
}

func(service *TasksServiceImpl)CompletedTask(ctx context.Context, taskId string){
	tx,err := service.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	idUser, _ , err := helper.ExtractUserFromToken(ctx)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	task,err := service.TaskRepository.FindTaskById(ctx,tx,taskId,idUser)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TaskRepository.CompletedTask(ctx,tx,task)
}
