package helper

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

func ParseDueDate(dueDateStr string) (sql.NullTime, error) {
	if dueDateStr == "" {
			return sql.NullTime{Valid: false}, nil
	}
	defaultTime := " 23:59:59"
	if !strings.Contains(dueDateStr, ":") {
			dueDateStr += defaultTime
	}

	formattedDate := "2006-01-02 15:04:05"
	parsedDate, err := time.Parse(formattedDate, dueDateStr)
	if err != nil {
			return sql.NullTime{}, errors.New("Invalid due_date format: " + dueDateStr + " . Expected format: YYYY-MM-DD or YYYY-MM-DD HH:MM:SS")
	}
	
	return sql.NullTime{Time: parsedDate, Valid: true}, nil
}
