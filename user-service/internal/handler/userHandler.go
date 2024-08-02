package handler

import (
	"context"

	"github.com/rishad004/hotel-booking-platform/user-service/internal/service"
	pb "github.com/rishad004/hotel-booking-platform/user-service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	Service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

func (h *UserHandler) SignUp(c context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	user, err := h.Service.SignUp(req.Name, req.Password, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{Message: "User signed up successfully", Userid: user.ID}, nil
}

func (h *UserHandler) LogIn(c context.Context, req *pb.LogInRequest) (*pb.LogInResponse, error) {
	user, err := h.Service.LogIn(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LogInResponse{Token: user.ID}, nil
}
