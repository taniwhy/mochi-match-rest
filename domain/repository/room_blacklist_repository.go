//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

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
