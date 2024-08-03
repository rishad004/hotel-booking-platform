package handler

import (
	"context"

	"github.com/rishad004/hotel-booking-platform/hotel-service/internal/service"
	pb "github.com/rishad004/hotel-booking-platform/hotel-service/proto"
)

type HotelHandler struct {
	pb.UnimplementedHotelServiceServer
	Service service.HotelService
}

func NewHotelHandler(svc service.HotelService) *HotelHandler {
	return &HotelHandler{Service: svc}
}

func (h *HotelHandler) AddRoom(c context.Context, req *pb.AddRoomRequest) (*pb.RoomResponse, error) {
	_, err := h.Service.AddRoom(int(req.Room))
	if err != nil {
		return nil, err
	}
	return &pb.RoomResponse{Message: "Room added successfully!"}, nil
}

func (h *HotelHandler) ListRooms(c context.Context, req *pb.Empty) (*pb.RoomListResponse, error) {
	room, err := h.Service.ListRooms()
	if err != nil {
		return nil, err
	}
	var rooms []*pb.Room
	for _, v := range room {
		rooms = append(rooms, &pb.Room{
			Room:        int32(v.Room),
			Availablity: v.Availability,
		})
	}
	return &pb.RoomListResponse{Rooms: rooms}, nil
}

func (h *HotelHandler) RemoveRoom(c context.Context, req *pb.RemoveRoomRequest) (*pb.RoomResponse, error) {
	if err := h.Service.RemoveRoom(int(req.RoomId)); err != nil {
		return nil, err
	}
	return &pb.RoomResponse{Message: "Room removed successfully!"}, nil
}

func (h *HotelHandler) ChangeStatus(c context.Context, req *pb.ChangeStatusRequest) (*pb.RoomResponse, error) {
	if err := h.Service.ChangeStatus(int(req.Room), req.Status); err != nil {
		return nil, err
	}
	return &pb.RoomResponse{Message: "Status changed to : " + req.Status}, nil
}
