package controller

import (
	"net/http"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/exception"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
	service "github.com/ryhnfhrza/Golang-To-Do-List-API/service/TasksService"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/util"
)

type TasksControllerImpl struct {
	TasksService service.TasksService
}

func NewTasksController(tasksService service.TasksService)TasksController{
	return &TasksControllerImpl{
		TasksService: tasksService,
	}
}

func(controller *TasksControllerImpl)CreateTask(writer http.ResponseWriter,request *http.Request){
	tasksRequest := web.CreateTaskRequest{}
	helper.ReadFromRequestBody(request,&tasksRequest)

	_, ok := request.Context().Value(util.TokenKey).(string)
	if !ok {
		exception.WriteUnauthorizedError(writer, "Token not found in context")
		return
	}
	
	TasksResponse := controller.TasksService.CreateTask(request.Context(),tasksRequest)

	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Data: TasksResponse,
	}
	
	helper.WriteToResponseBody(writer,webResponse)
}