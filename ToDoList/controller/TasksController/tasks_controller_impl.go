package controller

import (
	"net/http"

	"github.com/gorilla/mux"
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
func(controller *TasksControllerImpl)UpdateTask(writer http.ResponseWriter,request *http.Request){
	taskUpdateRequest := web.UpdateTaskRequest{}
	helper.ReadFromRequestBody(request,&taskUpdateRequest)

	_, ok := request.Context().Value(util.TokenKey).(string)
	if !ok {
		exception.WriteUnauthorizedError(writer, "Token not found in context")
		return
	}
	
	vars := mux.Vars(request)
	taskId := vars["taskId"]
	taskUpdateRequest.IdTask = taskId
	
	TasksResponse := controller.TasksService.UpdateTask(request.Context(),taskUpdateRequest)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: TasksResponse,
	}
	
	helper.WriteToResponseBody(writer,webResponse)
}