//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
)

// IChatPostRepository : チャット投稿のリポジトリ
type IChatPostRepository interface {
	FindByRoomID(id string) ([]*output.ChatPostResBody, error)
	FindByRoomIDAndLimit(id string, limit int) ([]*output.ChatPostResBody, error)
	FindByRoomIDAndOffset(id, offset string) ([]*output.ChatPostResBody, error)
	FindByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*output.ChatPostResBody, error)
	Insert(room *models.ChatPost) (*output.ChatPostResBody, error)
	Delete(room *models.ChatPost) error
}
