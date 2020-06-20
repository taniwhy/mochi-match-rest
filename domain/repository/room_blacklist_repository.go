package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// RoomBlacklistRepository : ルームブラックリストのリポジトリ
type RoomBlacklistRepository interface {
	FindAllBlacklist() ([]*models.RoomBlacklist, error)
	FindBlacklistByID(id int64) (*models.RoomBlacklist, error)
	InsertBlacklist(room *models.RoomBlacklist) error
	DeleteBlacklist(room *models.RoomBlacklist) error
}
