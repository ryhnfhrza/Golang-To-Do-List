package domain

import (
	"database/sql"
	"time"
)
type Tasks struct{
		IdTasks string
    UserId string
    Title string
    Description string
    Completed int
    DueDate sql.NullTime
    Notified int
    CreatedAt time.Time
    UpdatedAt time.Time
}