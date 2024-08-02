package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rishad004/hotel-booking-platform/user-service/internal/service"
)

type UserRepo interface {
	CreateUser(name, password, email string) (*service.User, error)
	FindUserByEmail(email string) (*service.User, error)
}

type userRepo struct {
	users map[string]*service.User
}

func NewUserRepo() UserRepo {
	return &userRepo{users: make(map[string]*service.User)}
}

func (r *userRepo) CreateUser(name, password, email string) (*service.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return nil, errors.New("user already exists")
		}
	}
	userID := uuid.New().String()
	user := &service.User{
		ID:       userID,
		Name:     name,
		Password: password,
		Email:    email,
	}
	r.users[userID] = user
	return user, nil
}

func (r *userRepo) FindUserByEmail(email string) (*service.User, error) {

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
