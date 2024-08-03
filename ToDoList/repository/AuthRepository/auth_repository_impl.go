package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
)

type AuthRepositoryImpl struct{

}

func NewAuthRepository()AuthRepository{
	return &AuthRepositoryImpl{}
}

func(Repository *AuthRepositoryImpl)Registration(ctx context.Context, tx *sql.Tx, users domain.Users) domain.Users{
	SQL := "insert into users (id,username,email,password,created_at,Updated_at) values (?,?,?,?,?,?)"
	_,err := tx.ExecContext(ctx,SQL,users.Id,users.Username,users.Email,users.Password,users.CreatedAt,users.UpdatedAt)
	helper.PanicIfError(err)

	return users
}

func(Repository *AuthRepositoryImpl)Login(ctx context.Context, tx *sql.Tx, username, password string) (domain.Users,error){
	SQL := "select username,email,id from users where username = ? and password = ?"
	rows , err := tx.QueryContext(ctx, SQL, username,password)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.Users{}
	if rows.Next(){
		err := rows.Scan(&user.Username,&user.Email,&user.Id)
		helper.PanicIfError(err)
		return user,nil
	}else{
		return user,errors.New("invalid password")
	}
}

func(Repository *AuthRepositoryImpl)CheckEmail(ctx context.Context, tx *sql.Tx, email string) (string,error){
	SQL := "select email from users where email = ?"
	rows , err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next(){
		return "",errors.New("Can't use email " + email +" , because email is already registered")
	}else{
		return email,nil
	}
}






