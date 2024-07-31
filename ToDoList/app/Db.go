package app

import (
	"database/sql"
	"time"
)

func NewDb() *sql.DB{
	db,err := sql.Open("mysql","root:Rayhan22@tcp(localhost:3306)/todo_list?parseTime=true")
	if err != nil{
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	
	return db
}