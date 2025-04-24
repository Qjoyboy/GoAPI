package usersrc

import "goapi/internal/src"

type User struct {
	ID       string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Tasks    []src.Task `json:"tasks"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
