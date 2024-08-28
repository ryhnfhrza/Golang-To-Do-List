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

func(controller *TasksControllerImpl)DeleteTask(writer http.ResponseWriter, request *http.Request){
	_, ok := request.Context().Value(util.TokenKey).(string)
	if !ok {
		exception.WriteUnauthorizedError(writer, "Token not found in context")
		return
	}
	
	vars := mux.Vars(request)
	taskId := vars["taskId"]
	
	
	controller.TasksService.DeleteTask(request.Context(),taskId)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
	}
	
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *TasksControllerImpl)FindAllTask(writer http.ResponseWriter, request *http.Request){
	
	_, ok := request.Context().Value(util.TokenKey).(string)
	if !ok {
		exception.WriteUnauthorizedError(writer, "Token not found in context")
		return
	}

	sortBy := request.URL.Query().Get("sort_by")
	order := request.URL.Query().Get("order")
	
	TasksResponses := controller.TasksService.FindAllTask(request.Context(),sortBy,order)
	
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: TasksResponses,
	}
	
	helper.WriteToResponseBody(writer,webResponse)
}

func(controller *TasksControllerImpl)SearchTask(writer http.ResponseWriter, request *http.Request){
	_, ok := request.Context().Value(util.TokenKey).(string)
	if !ok {
		exception.WriteUnauthorizedError(writer, "Token not found in context")
		return
	}
	
	vars := mux.Vars(request)
	task := vars["keyword"]

	sortBy := request.URL.Query().Get("sort_by")
	order := request.URL.Query().Get("order")
	
	taskResponse := controller.TasksService.SearchTask(request.Context(),task,sortBy,order)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: taskResponse,
	}
	
	helper.WriteToResponseBody(writer,webResponse)
}
