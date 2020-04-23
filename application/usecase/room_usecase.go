package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// RoomUseCase :
type RoomUseCase interface {
	FindAllRoom() ([]*models.Room, error)
	FindRoomByID(id int64) (*models.Room, error)
	InsertRoom(room *models.Room) error
	UpdateRoom(room *models.Room) error
	DeleteRoom(room *models.Room) error
}

type roomUsecase struct {
	roomRepository repository.RoomRepository
}

// NewRoomUsecase :
func NewRoomUsecase(rR repository.RoomRepository) RoomUseCase {
	return &roomUsecase{
		roomRepository: rR,
	}
}

func (rU roomUsecase) FindAllRoom() ([]*models.Room, error) {
	rooms, err := rU.roomRepository.FindAllRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rU roomUsecase) FindRoomByID(id int64) (*models.Room, error) {
	room, err := rU.roomRepository.FindRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (rU roomUsecase) InsertRoom(room *models.Room) error {
	err := rU.roomRepository.InsertRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomUsecase) UpdateRoom(room *models.Room) error {
	err := rU.roomRepository.UpdateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomUsecase) DeleteRoom(room *models.Room) error {
	err := rU.roomRepository.DeleteRoom(room)
	if err != nil {
		return err
	}
	return nil
}
