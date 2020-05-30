package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// RoomReservationUseCase :
type RoomReservationUseCase interface {
	FindAllRoomReservation() ([]*dbmodel.RoomReservation, error)
	FindRoomReservationByID(id int64) (*dbmodel.RoomReservation, error)
	InsertRoomReservation(roomReservation *dbmodel.RoomReservation) error
	UpdateRoomReservation(roomReservation *dbmodel.RoomReservation) error
	DeleteRoomReservation(roomReservation *dbmodel.RoomReservation) error
}

type roomReservationUsecase struct {
	roomReservationRepository repository.RoomReservationRepository
}

// NewRoomReservationUsecase :
func NewRoomReservationUsecase(rR repository.RoomReservationRepository) RoomReservationUseCase {
	return &roomReservationUsecase{
		roomReservationRepository: rR,
	}
}

func (rU roomReservationUsecase) FindAllRoomReservation() ([]*dbmodel.RoomReservation, error) {
	roomReservations, err := rU.roomReservationRepository.FindAllRoomReservation()
	if err != nil {
		return nil, err
	}
	return roomReservations, nil
}

func (rU roomReservationUsecase) FindRoomReservationByID(id int64) (*dbmodel.RoomReservation, error) {
	roomReservation, err := rU.roomReservationRepository.FindRoomReservationByID(id)
	if err != nil {
		return nil, err
	}
	return roomReservation, nil
}

func (rU roomReservationUsecase) InsertRoomReservation(roomReservation *dbmodel.RoomReservation) error {
	err := rU.roomReservationRepository.InsertRoomReservation(roomReservation)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomReservationUsecase) UpdateRoomReservation(room *dbmodel.RoomReservation) error {
	err := rU.roomReservationRepository.UpdateRoomReservation(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomReservationUsecase) DeleteRoomReservation(room *dbmodel.RoomReservation) error {
	err := rU.roomReservationRepository.DeleteRoomReservation(room)
	if err != nil {
		return err
	}
	return nil
}
