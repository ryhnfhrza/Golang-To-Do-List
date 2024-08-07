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

func(Repository *AuthRepositoryImpl)Login(ctx context.Context, tx *sql.Tx, username string) (domain.Users,error){
	SQL := "select username,email,id,password from users where username = ? "
	rows , err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.Users{}
	if rows.Next(){
		err := rows.Scan(&user.Username,&user.Email,&user.Id,&user.Password)
		helper.PanicIfError(err)
		return user,nil
	}else{
		return user,errors.New("invalid username or password ")
	}
}

func(Repository *AuthRepositoryImpl)CheckEmail(ctx context.Context, tx *sql.Tx, email string) (bool,error){
	SQL := "SELECT COUNT(email) FROM users WHERE email = ?"
    var emailCount int
    err := tx.QueryRowContext(ctx, SQL, email).Scan(&emailCount)
    if err != nil {
        return false, err 
    }

    if emailCount > 0 {
        return true, nil 
    }
    return false, nil 
}

func(Repository *AuthRepositoryImpl)CheckUsername(ctx context.Context, tx *sql.Tx, username string) (bool,error){
	SQL := "SELECT COUNT(username) FROM users WHERE username = ?"
	var usernameCount int
	err := tx.QueryRowContext(ctx, SQL, username).Scan(&usernameCount)
	if err != nil {
			return false, err 
	}

	if usernameCount > 0 {
			return true, nil 
	}
	return false, nil 
}





