package models

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
