package repository

import (
	"errors"

	"github.com/rishad004/hotel-booking-platform/hotel-service/internal/service"
)

type HotelRepo interface {
	GetRoom() ([]*service.Room, error)
	CreateRoom(room int) (*service.Room, error)
	DeleteRoom(room int) error
	ChangeAvailability(room int, status string) error
}

type hotelRepo struct {
	rooms map[int]*service.Room
}

func NewHotelRepo() HotelRepo {
	return &hotelRepo{rooms: make(map[int]*service.Room)}
}

func (r *hotelRepo) GetRoom() ([]*service.Room, error) {
	var rooms []*service.Room
	for _, v := range r.rooms {
		rooms = append(rooms, v)
	}
	return rooms, nil
}

func (r *hotelRepo) CreateRoom(room int) (*service.Room, error) {
	_, is := r.rooms[room]
	if is {
		return nil, errors.New("room already exists")
	}
	r.rooms[room] = &service.Room{Room: room, Availability: true}
	return r.rooms[room], nil
}

func (r *hotelRepo) DeleteRoom(room int) error {
	_, is := r.rooms[room]
	if !is {
		return errors.New("room not found")
	}
	delete(r.rooms, room)
	return nil
}

func (r *hotelRepo) ChangeAvailability(room int, status string) error {
	if _, is := r.rooms[room]; !is {
		return errors.New("room not found")
	}
	if status == "available" {
		r.rooms[room].Availability = true
		return nil
	}
	r.rooms[room].Availability = false
	return nil
}
