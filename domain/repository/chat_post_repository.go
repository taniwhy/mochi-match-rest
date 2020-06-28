//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// IChatPostRepository : チャット投稿のリポジトリ
type IChatPostRepository interface {
	FindByRoomID(id string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimit(id string, limit int) ([]*models.ChatPost, error)
	FindByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*models.ChatPost, error)
	Insert(room *models.ChatPost) error
	Delete(room *models.ChatPost) error
}
