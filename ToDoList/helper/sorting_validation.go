package helper

import (
	"errors"
)

func ValidateSortParams(sortBy, order string) (string, string, error) {
    validSortBy := map[string]bool{"title": true, "due_date": true, "created_at": true}
    validOrder := map[string]bool{"ASC": true, "DESC": true}

    if !validSortBy[sortBy] || !validOrder[order] {
        return "", "", errors.New("bad Request")
    }

    return sortBy, order, nil
}
