package dbmodel

import (
	"time"
)

// ChatPost : chat_postテーブルモデル
type ChatPost struct {
	ChatPostID string    `json:"id"`
	RoomID     string    `json:"room"`
	UserID     string    `json:"user"`
	Message    string    `json:"message" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
}
