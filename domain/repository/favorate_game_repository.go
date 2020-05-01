package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// FavorateGameRepository : ユーザーのリポジトリ
type FavorateGameRepository interface {
	FindAllRoom() ([]*models.Room, error)
	FindRoomByID(id int64) (*models.Room, error)
	InsertRoom(room *models.Room) error
	UpdateRoom(room *models.Room) error
	DeleteRoom(room *models.Room) error
}
