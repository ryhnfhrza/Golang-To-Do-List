package exception

import "github.com/go-sql-driver/mysql"

func HandleSQLError() {
	if r := recover(); r != nil {
		if sqlErr, ok := r.(*mysql.MySQLError); ok {
			if sqlErr.Number == 1644 {
				panic(NewBadRequestError(sqlErr.Message))
			}
		}
		panic(r)
	}
}
