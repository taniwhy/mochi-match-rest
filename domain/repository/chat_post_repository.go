package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// ChatPostRepository : ユーザーのリポジトリ
type ChatPostRepository interface {
	FindAllChatPost() ([]*models.ChatPost, error)
	FindChatPostByRoomID(id string) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndLimit(id string, limit int) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*models.ChatPost, error)
	InsertChatPost(room *models.ChatPost) error
	DeleteChatPost(room *models.ChatPost) error
}
