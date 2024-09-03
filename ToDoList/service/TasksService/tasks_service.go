package service

import (
	"context"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
)

type TasksService interface {
	CreateTask(ctx context.Context, request web.CreateTaskRequest) web.UserTasksResponses
	UpdateTask(ctx context.Context, request web.UpdateTaskRequest) web.UserTasksResponses
	DeleteTask(ctx context.Context, taskId string)
	FindAllTask(ctx context.Context,sortBy,order string)web.UserTasksResponses
	SearchTask(ctx context.Context, keyword,sortBy,order string) web.UserTasksResponses
	SendDueDateReminders(ctx context.Context)error
	CompletedTask(ctx context.Context, taskId string)
}