package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
)

type TasksRepository interface{
	CreateTask(ctx context.Context, tx *sql.Tx, task domain.Tasks) domain.Tasks
}