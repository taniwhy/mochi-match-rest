package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
)

// RoomRepository : ルームのリポジトリ
type RoomRepository interface {
	FindList() ([]*output.RoomResBody, error)
	FindByLimitAndOffset(limit, offset int) ([]*output.RoomResBody, error)
	FindByID(id string) (*models.Room, error)
	FindByUserID(id string) ([]*models.Room, error)
	FindUnlockByID(id string) (*models.Room, error)
	Insert(room *models.Room) error
	Update(room *models.Room) error
	Delete(room *models.Room) error
}
