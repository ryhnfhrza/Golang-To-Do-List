package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
)

type TasksRepository interface{
	CreateTask(ctx context.Context, tx *sql.Tx, task domain.Tasks) domain.Tasks
	UpdateTask(ctx context.Context, tx *sql.Tx,task domain.Tasks)domain.Tasks
	DeleteTask(ctx context.Context, tx *sql.Tx,task domain.Tasks)
	FindTaskById(ctx context.Context,tx *sql.Tx,idTask, idUser string) (domain.Tasks,error)
	FindAllTask(ctx context.Context,tx *sql.Tx,idUser,sortBy,order string)[]domain.Tasks
	SearchTask(ctx context.Context, tx *sql.Tx, keyword, idUser,sortBy,order string)([]domain.Tasks,error)
	FindTaskDueInOneDay(ctx context.Context,tx *sql.Tx)[]domain.UserTasks
	FindTaskDueInOneHour(ctx context.Context,tx *sql.Tx)[]domain.UserTasks
	UpdateTaskAfterNotification(ctx context.Context, tx *sql.Tx,task domain.Tasks)error
	CompletedTask(ctx context.Context, tx *sql.Tx,task domain.Tasks)
}