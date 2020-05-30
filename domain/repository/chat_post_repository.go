package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// ChatPostRepository : ユーザーのリポジトリ
type ChatPostRepository interface {
	FindAllChatPost() ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomID(id string) ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomIDAndLimit(id string, limit int) ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomIDAndOffset(id, offset string) ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*dbmodel.ChatPost, error)
	InsertChatPost(room *dbmodel.ChatPost) error
	DeleteChatPost(room *dbmodel.ChatPost) error
}
