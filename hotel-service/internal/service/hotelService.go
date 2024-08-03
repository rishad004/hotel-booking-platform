package service

type Room struct {
	Room         int
	Availability bool
}

type HotelRepo interface {
	GetRoom() ([]*Room, error)
	CreateRoom(room int) (*Room, error)
	DeleteRoom(room int) error
	ChangeAvailability(room int, status string) error
}

type HotelService interface {
	ListRooms() ([]*Room, error)
	AddRoom(room int) (*Room, error)
	RemoveRoom(room int) error
	ChangeStatus(room int, status string) error
}

type hotelService struct {
	repo HotelRepo
}

func NewHotelService(repo HotelRepo) *hotelService {
	return &hotelService{repo: repo}
}

func (s *hotelService) ListRooms() ([]*Room, error) {
	rooms, err := s.repo.GetRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *hotelService) AddRoom(room int) (*Room, error) {
	roomDetail, err := s.repo.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	return roomDetail, nil
}

func (s *hotelService) RemoveRoom(room int) error {
	if err := s.repo.DeleteRoom(room); err != nil {
		return err
	}
	return nil
}

func (s *hotelService) ChangeStatus(room int, status string) error {
	if err := s.repo.ChangeAvailability(room, status); err != nil {
		return err
	}
	return nil
}
