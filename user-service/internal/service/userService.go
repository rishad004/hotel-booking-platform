package service

import (
	"context"
	"errors"

	hotel_pb "github.com/rishad004/hotel-booking-platform/user-service/proto/client/hotel"
)

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
	ListHotel() (*hotel_pb.RoomListResponse, error)
}

type userService struct {
	hotelClient hotel_pb.HotelServiceClient
	repo        UserRepo
}

func NewUserService(repo UserRepo, hotelClient hotel_pb.HotelServiceClient) *userService {
	return &userService{repo: repo, hotelClient: hotelClient}
}

func (s *userService) LogIn(email, password string) (*User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) SignUp(name, password, email string) (*User, error) {
	user, err := s.repo.CreateUser(name, password, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ListHotel() (*hotel_pb.RoomListResponse, error) {
	room, err := s.hotelClient.ListRooms(context.Background(), &hotel_pb.Empty{})
	if err != nil {
		return nil, err
	}
	return room, nil
}
