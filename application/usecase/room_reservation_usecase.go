package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IRoomReservationUseCase : インターフェース
type IRoomReservationUseCase interface {
	FindAllRoomReservation() ([]*models.RoomReservation, error)
	FindRoomReservationByID(id int64) (*models.RoomReservation, error)
	InsertRoomReservation(roomReservation *models.RoomReservation) error
	UpdateRoomReservation(roomReservation *models.RoomReservation) error
	DeleteRoomReservation(roomReservation *models.RoomReservation) error
}

type roomReservationUsecase struct {
	roomReservationRepository repository.RoomReservationRepository
}

// NewRoomReservationUsecase : RoomReservationユースケースの生成
func NewRoomReservationUsecase(rR repository.RoomReservationRepository) IRoomReservationUseCase {
	return &roomReservationUsecase{
		roomReservationRepository: rR,
	}
}

func (rU roomReservationUsecase) FindAllRoomReservation() ([]*models.RoomReservation, error) {
	roomReservations, err := rU.roomReservationRepository.FindAllRoomReservation()
	if err != nil {
		return nil, err
	}
	return roomReservations, nil
}

func (rU roomReservationUsecase) FindRoomReservationByID(id int64) (*models.RoomReservation, error) {
	roomReservation, err := rU.roomReservationRepository.FindRoomReservationByID(id)
	if err != nil {
		return nil, err
	}
	return roomReservation, nil
}

func (rU roomReservationUsecase) InsertRoomReservation(roomReservation *models.RoomReservation) error {
	err := rU.roomReservationRepository.InsertRoomReservation(roomReservation)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomReservationUsecase) UpdateRoomReservation(room *models.RoomReservation) error {
	err := rU.roomReservationRepository.UpdateRoomReservation(room)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomReservationUsecase) DeleteRoomReservation(room *models.RoomReservation) error {
	err := rU.roomReservationRepository.DeleteRoomReservation(room)
	if err != nil {
		return err
	}
	return nil
}
