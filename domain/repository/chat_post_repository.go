package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// ChatPostRepository : チャット投稿のリポジトリ
type ChatPostRepository interface {
	FindByRoomID(id string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimit(id string, limit int) ([]*models.ChatPost, error)
	FindByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*models.ChatPost, error)
	Insert(room *models.ChatPost) error
	Delete(room *models.ChatPost) error
}
