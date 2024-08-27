package repository

import (
	"context"
	"database/sql"
	"errors"

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

func(Repository *TasksRepositoryImpl)UpdateTask(ctx context.Context, tx *sql.Tx, task domain.Tasks) domain.Tasks{
	SQL := "update tasks set title = ?, description = ?, due_date = ? where id = ? and user_id = ?"  

	_,err := tx.ExecContext(ctx,SQL,task.Title,task.Description,task.DueDate,task.IdTasks,task.UserId)
	helper.PanicIfError(err)

	return task
}

func(Repository *TasksRepositoryImpl)DeleteTask(ctx context.Context, tx *sql.Tx,task domain.Tasks){
	SQL := "delete from tasks where id = ? and user_id = ?"  

	_,err := tx.ExecContext(ctx,SQL,task.IdTasks,task.UserId)
	helper.PanicIfError(err)
}

func(Repository *TasksRepositoryImpl)FindTaskById(ctx context.Context,tx *sql.Tx, idTask, idUser string) (domain.Tasks,error){
	SQL := "select id,user_id,title,description,completed,due_date,notified,created_at,updated_at from tasks where id = ? and user_id = ?"
	rows,err := tx.QueryContext(ctx,SQL,idTask,idUser)
	helper.PanicIfError(err)
	defer rows.Close()

	task := domain.Tasks{}
	if rows.Next(){
		err := rows.Scan(&task.IdTasks,&task.UserId,&task.Title,&task.Description,&task.Completed,&task.DueDate,&task.Notified,&task.CreatedAt,&task.UpdatedAt)
		helper.PanicIfError(err)
		return task , nil

	}else{
		return task,errors.New("task not found")
	}
}