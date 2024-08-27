package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateTaskID() string {
	id := uuid.New()
	return strings.ReplaceAll(id.String(), "-", "")
}
