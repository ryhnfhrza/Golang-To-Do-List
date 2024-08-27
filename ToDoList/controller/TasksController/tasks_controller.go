package controller

import (
	"net/http"
)

type TasksController interface {
	CreateTask(writer http.ResponseWriter, request *http.Request)
	UpdateTask(writer http.ResponseWriter, request *http.Request)
	DeleteTask(writer http.ResponseWriter, request *http.Request)
}