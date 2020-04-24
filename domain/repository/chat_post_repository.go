package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// ChatPostRepository : ユーザーのリポジトリ
type ChatPostRepository interface {
	FindAllChatPost() ([]*models.ChatPost, error)
	FindChatPostByRoomID(id int64) ([]*models.ChatPost, error)
	InsertChatPost(room *models.ChatPost) error
	DeleteChatPost(room *models.ChatPost) error
}
