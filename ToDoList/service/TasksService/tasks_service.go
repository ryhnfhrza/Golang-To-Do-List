package service

import (
	"context"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
)

type TasksService interface {
	CreateTask(ctx context.Context, request web.CreateTaskRequest) web.TaskResponse
	UpdateTask(ctx context.Context, request web.UpdateTaskRequest) web.TaskResponse
}