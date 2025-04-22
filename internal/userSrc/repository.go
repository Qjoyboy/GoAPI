package usersrc

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(u *User) error
	GetUsers() ([]User, error)
	GetUserById(id string) (User, error)
	UpdateUser(u User) error
	DeleteUser(id string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(u *User) error {
	return r.db.Create(&u).Error
}

func (r *userRepo) GetUsers() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepo) GetUserById(id string) (User, error) {
	var user User
	err := r.db.First(&user, "id=?", id).Error
	return user, err
}

func (r *userRepo) UpdateUser(u User) error {
	return r.db.Save(&u).Error
}

func (r *userRepo) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id=?", id).Error
}
