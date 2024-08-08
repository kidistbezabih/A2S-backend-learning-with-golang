package models

type Task struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
