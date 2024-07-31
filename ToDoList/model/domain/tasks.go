package domain

import "time"

type Tasks struct{
		IdTasks string
    UserId string
    Title string
    Description string
    Completed string
    DueDate time.Time
    Notified string
    CreatedAt time.Time
    UpdatedAt time.Time
}