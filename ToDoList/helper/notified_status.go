package helper

import (
	"database/sql"
	"time"
)

func CalculateNotificationStatus(dueDate sql.NullTime) int {
		if !dueDate.Valid{
			return 0
		}

    now := time.Now()
    duration := dueDate.Time.Sub(now)

    
    if duration <= 23*time.Hour + 59*time.Minute + 59*time.Second {
        return 1
    }
    
    if duration > 24*time.Hour {
        return 0
    }
    return 0
}
