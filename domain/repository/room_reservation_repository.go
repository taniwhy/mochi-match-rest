package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// RoomReservationRepository : ユーザーのリポジトリ
type RoomReservationRepository interface {
	FindAllRoomReservation() ([]*models.RoomReservation, error)
	FindRoomReservationByID(id int64) (*models.RoomReservation, error)
	InsertRoomReservation(room *models.RoomReservation) error
	UpdateRoomReservation(room *models.RoomReservation) error
	DeleteRoomReservation(room *models.RoomReservation) error
}
