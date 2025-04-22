package usersrc

import "fmt"

type UserService interface {
	CreateUser(email, password string) (User, error)
	GetUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id, email, password string) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) emailValidate(email string) error {
	for _, v := range email {
		if v == '@' {
			return nil
		}
	}
	return fmt.Errorf("email hasnt '@'")
}

// CreateUser implements UserService.
func (s *userService) CreateUser(email string, password string) (User, error) {

	if err := s.emailValidate(email); err != nil {
		return User{}, nil
	}

	user := User{
		Email:    email,
		Password: password,
	}

	if err := s.repo.CreateUser(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

// GetUserByID implements UserService.
func (s *userService) GetUserByID(id string) (User, error) {
	return s.repo.GetUserById(id)
}

// GetUsers implements UserService.
func (s *userService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

// UpdateUser implements UserService.
func (s *userService) UpdateUser(id string, email string, password string) (User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return User{}, err
	}

	if err := s.emailValidate(email); err != nil {
		return User{}, nil
	}

	user.Email = email
	user.Password = password

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}
	return user, nil
}

// DeleteUser implements UserService.
func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
