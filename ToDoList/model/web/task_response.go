package web

type TaskResponse struct {
	UserName    string `json:"user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   string `json:"status"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
}