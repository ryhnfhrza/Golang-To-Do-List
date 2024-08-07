package repository

import (
	"context"
	"database/sql"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
)

type AuthRepository interface{
	Registration(ctx context.Context, tx *sql.Tx, users domain.Users) domain.Users
	Login(ctx context.Context, tx *sql.Tx, username string ) (domain.Users,error)
	CheckEmail (ctx context.Context, tx *sql.Tx, email string ) (bool,error)
	CheckUsername (ctx context.Context, tx *sql.Tx, username string ) (bool,error)
}