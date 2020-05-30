package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// RoomReservationRepository : ユーザーのリポジトリ
type RoomReservationRepository interface {
	FindAllRoomReservation() ([]*dbmodel.RoomReservation, error)
	FindRoomReservationByID(id int64) (*dbmodel.RoomReservation, error)
	InsertRoomReservation(room *dbmodel.RoomReservation) error
	UpdateRoomReservation(room *dbmodel.RoomReservation) error
	DeleteRoomReservation(room *dbmodel.RoomReservation) error
}
