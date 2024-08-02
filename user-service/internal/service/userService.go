package service

import "errors"

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepo interface {
	CreateUser(name, password, email string) (*User, error)
	FindUserByEmail(email string) (*User, error)
}

type UserService interface {
	SignUp(name, password, email string) (*User, error)
	LogIn(email, password string) (*User, error)
}

type userService struct {
	repo UserRepo
}

func (s *userService) LogIn(email, password string) (*User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func NewUserService(repo UserRepo) *userService {
	return &userService{repo: repo}
}

func (s *userService) SignUp(name, password, email string) (*User, error) {
	user, err := s.repo.CreateUser(name, password, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
