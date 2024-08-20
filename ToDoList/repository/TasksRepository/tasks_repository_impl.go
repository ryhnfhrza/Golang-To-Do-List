package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
)

type TasksRepositoryImpl struct{

}

func NewTasksRepository()TasksRepository{
	return &TasksRepositoryImpl{}
}

func(Repository *TasksRepositoryImpl)CreateTask(ctx context.Context, tx *sql.Tx, task domain.Tasks) domain.Tasks{
	SQL := "insert into tasks (id,user_id,title,description,completed,due_date,notified,created_at,updated_at) values (?,?,?,?,?,?,?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,task.IdTasks,task.UserId,task.Title,task.Description,task.Completed,task.DueDate,task.Notified,task.CreatedAt,task.UpdatedAt)
	helper.PanicIfError(err)

	return task
}