package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// RoomRepository : ユーザーのリポジトリ
type RoomRepository interface {
	FindAllRoom() ([]*models.Room, error)
	FindRoomByID(id int64) (*models.Room, error)
	InsertRoomDetail(room *models.Room) error
	UpdateRoomDetail(room *models.Room) error
	DeleteRoomDetail(room *models.Room) error
}
