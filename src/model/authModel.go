package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Pasword  string `json:"password"`
	Email    string `json:"email"`
}
