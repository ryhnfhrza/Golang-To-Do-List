package web

type CreateTaskRequest struct {
	Title       string `validate:"required,min=1,max=255" json:"title"`
	Description string `validate:"max=65535,omitempty" json:"description"`
	DueDate     string `json:"due_date"`
}