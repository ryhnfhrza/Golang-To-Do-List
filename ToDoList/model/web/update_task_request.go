package web

type UpdateTaskRequest struct {
	IdTask      string `validate:"required" json:"id_task"`
	Title       string `validate:"max=255" json:"title"`
	Description string `validate:"max=65535" json:"description"`
	DueDate     string `json:"due_date"`
}