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
	FindTaskById(ctx context.Context,tx *sql.Tx, idTask, idUser string) (domain.Tasks,error)
	FindAllTask(ctx context.Context,tx *sql.Tx,idUser string)[]domain.Tasks
}