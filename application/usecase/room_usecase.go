package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// RoomUseCase :
type RoomUseCase interface {
	FindAllRoom() ([]*dbmodel.Room, error)
	FindRoomByID(id string) (*dbmodel.Room, error)
	InsertRoom(room *dbmodel.Room) error
	UpdateRoom(room *dbmodel.Room) error
	DeleteRoom(room *dbmodel.Room) error
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

func (rU roomUsecase) FindAllRoom() ([]*dbmodel.Room, error) {
	rooms, err := rU.roomRepository.FindAllRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rU roomUsecase) FindRoomByID(id string) (*dbmodel.Room, error) {
	room, err := rU.roomRepository.FindRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (rU roomUsecase) InsertRoom(room *dbmodel.Room) error {
	err := rU.roomRepository.InsertRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomUsecase) UpdateRoom(room *dbmodel.Room) error {
	err := rU.roomRepository.UpdateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomUsecase) DeleteRoom(room *dbmodel.Room) error {
	err := rU.roomRepository.DeleteRoom(room)
	if err != nil {
		return err
	}
	return nil
}
