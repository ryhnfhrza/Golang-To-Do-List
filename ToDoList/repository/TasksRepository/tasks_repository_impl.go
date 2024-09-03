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
	SQL := "update tasks set title = ?, description = ?, due_date = ?, notified = ? where id = ? and user_id = ?"  

	_,err := tx.ExecContext(ctx,SQL,task.Title,task.Description,task.DueDate,task.Notified,task.IdTasks,task.UserId)
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

func(Repository *TasksRepositoryImpl)FindAllTask(ctx context.Context,tx *sql.Tx,idUser,sortBy,order string)[]domain.Tasks{
	SQL := "select title,description,due_date,completed,created_at from tasks where user_id = ? order by completed ASC, " + sortBy + " " + order 
	rows,err := tx.QueryContext(ctx,SQL,idUser)
	helper.PanicIfError(err)
	defer rows.Close()

	var tasks []domain.Tasks
	for rows.Next(){
		task := domain.Tasks{}
		err := rows.Scan(&task.Title,&task.Description,&task.DueDate,&task.Completed,&task.CreatedAt)
		helper.PanicIfError(err)
		tasks = append(tasks, task)

	}
	return tasks
}

func(Repository *TasksRepositoryImpl)SearchTask(ctx context.Context, tx *sql.Tx, keyword, idUser,sortBy,order string)([]domain.Tasks,error){
	SQL := `
        select title,description,due_date,completed,created_at 
        FROM tasks
        WHERE user_id = ?
        AND (title LIKE ? OR description LIKE ?) ORDER BY completed ASC, ` + sortBy + ` ` + order
    
	rows,err := tx.QueryContext(ctx,SQL,idUser,"%"+keyword+"%","%"+keyword+"%")
	helper.PanicIfError(err)
	defer rows.Close()

	var tasks []domain.Tasks
	for rows.Next(){
		task := domain.Tasks{}
		err := rows.Scan(&task.Title,&task.Description,&task.DueDate,&task.Completed,&task.CreatedAt)
		if err!= nil {
			return nil,errors.New("task not found")
		}
		tasks = append(tasks, task)

	}

	return tasks,nil
}

func(Repository *TasksRepositoryImpl)FindTaskDueInOneDay(ctx context.Context,tx *sql.Tx)[]domain.UserTasks{
	SQL := `SELECT 
					u.id ,
					u.email , 
					u.username , 
					t.id,
					t.title , 
					t.description , 
					t.due_date ,
					t.created_at,
					t.updated_at
					FROM 
							tasks t 
					JOIN 
							users u 
					ON 
							t.user_id = u.id 
					WHERE 
							t.due_date BETWEEN NOW() AND NOW() + INTERVAL 1 DAY 
							AND t.notified = 0
							AND t.completed = 0
				` 
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var userTasks []domain.UserTasks
	for rows.Next(){
		task := domain.UserTasks{}
		err := rows.Scan(&task.UserId,&task.Email,&task.Username,&task.IdTasks,&task.Title,&task.Description,&task.DueDate,&task.CreatedAt,&task.UpdatedAt)
		helper.PanicIfError(err)
		userTasks = append(userTasks, task)

	}
	return userTasks
}

func(Repository *TasksRepositoryImpl)FindTaskDueInOneHour(ctx context.Context,tx *sql.Tx)[]domain.UserTasks{
	SQL := `SELECT 
					u.id ,
					u.email , 
					u.username , 
					t.id , 
					t.title , 
					t.description , 
					t.due_date,
					t.created_at,
					t.updated_at
					FROM 
							tasks t 
					JOIN 
							users u 
					ON 
							t.user_id = u.id 
					WHERE 
							t.due_date BETWEEN NOW() AND NOW() + INTERVAL 1 HOUR
							AND t.notified = 1
							AND t.completed = 0
							` 

	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var userTasks []domain.UserTasks
	for rows.Next(){
		task := domain.UserTasks{}
		err := rows.Scan(&task.UserId,&task.Email,&task.Username,&task.IdTasks,&task.Title,&task.Description,&task.DueDate,&task.CreatedAt,&task.UpdatedAt)
		helper.PanicIfError(err)
		userTasks = append(userTasks, task)

	}
	return userTasks
}

func(Repository *TasksRepositoryImpl)UpdateTaskAfterNotification(ctx context.Context, tx *sql.Tx,task domain.Tasks)error{
	SQL := "update tasks set notified = ? where id = ? and user_id = ?"  

	_,err := tx.ExecContext(ctx,SQL,task.Notified,task.IdTasks,task.UserId)
	helper.PanicIfError(err)

	return err
}

func(Repository *TasksRepositoryImpl)CompletedTask(ctx context.Context, tx *sql.Tx,task domain.Tasks){
	SQL := "update tasks set completed = 1 where id = ? and user_id = ?"  

	_,err := tx.ExecContext(ctx,SQL,task.IdTasks,task.UserId)
	helper.PanicIfError(err)
}