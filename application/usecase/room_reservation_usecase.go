package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// RoomReservationUseCase :
type RoomReservationUseCase interface {
	FindAllRoomReservation() ([]*models.RoomReservation, error)
	FindRoomReservationByID(id int64) (*models.RoomReservation, error)
	InsertRoomReservation(roomReservation *models.RoomReservation) error
	UpdateRoomReservation(roomReservation *models.RoomReservation) error
	DeleteRoomReservation(roomReservation *models.RoomReservation) error
}

type roomReservationUsecase struct {
	roomRepository repository.RoomRepository
}

// NewRoomReservationUsecase :
func NewRoomReservationUsecase(rR repository.RoomRepository) RoomUseCase {
	return &roomUsecase{
		roomRepository: rR,
	}
}

func (rU roomReservationUsecase) FindAllRoomReservation() ([]*models.Room, error) {
	rooms, err := rU.roomRepository.FindAllRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rU roomReservationUsecase) FindRoomReservationByID(id int64) (*models.Room, error) {
	room, err := rU.roomRepository.FindRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (rU roomReservationUsecase) InsertRoomReservation(room *models.Room) error {
	err := rU.roomRepository.InsertRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomReservationUsecase) UpdateRoomReservation(room *models.Room) error {
	err := rU.roomRepository.UpdateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomReservationUsecase) DeleteRoomReservation(room *models.Room) error {
	err := rU.roomRepository.DeleteRoom(room)
	if err != nil {
		return err
	}
	return nil
}
