package usersrc

type User struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
