package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// RoomBlacklistRepository : ユーザーのリポジトリ
type RoomBlacklistRepository interface {
	FindAllBlacklist() ([]*dbmodel.RoomBlacklist, error)
	FindBlacklistByID(id int64) (*dbmodel.RoomBlacklist, error)
	InsertBlacklist(room *dbmodel.RoomBlacklist) error
	DeleteBlacklist(room *dbmodel.RoomBlacklist) error
}
