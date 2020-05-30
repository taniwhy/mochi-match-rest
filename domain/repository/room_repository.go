package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// RoomRepository : ユーザーのリポジトリ
type RoomRepository interface {
	FindAllRoom() ([]*dbmodel.Room, error)
	FindRoomByID(id string) (*dbmodel.Room, error)
	InsertRoom(room *dbmodel.Room) error
	UpdateRoom(room *dbmodel.Room) error
	DeleteRoom(room *dbmodel.Room) error
}
