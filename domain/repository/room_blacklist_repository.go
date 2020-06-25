package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// RoomBlacklistRepository : ルームブラックリストのリポジトリ
type RoomBlacklistRepository interface {
	FindByRoomID(string) ([]*models.RoomBlacklist, error)
	Insert(*models.RoomBlacklist) error
	Delete(string) error
}
