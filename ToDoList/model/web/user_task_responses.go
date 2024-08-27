package web

type UserTasksResponses struct {
	UserName string         `json:"username"`
	Tasks    []TaskResponse `json:"tasks"`
}
