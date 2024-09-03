package domain

import (
	"database/sql"
	"time"
)
type UserTasks struct{
		IdTasks string
    UserId string
		Email string
		Username string
    Title string
    Description string
    DueDate sql.NullTime
		CreatedAt time.Time
    UpdatedAt time.Time
}