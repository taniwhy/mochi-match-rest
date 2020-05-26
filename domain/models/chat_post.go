package models

import (
	"time"
)

// ChatPost : chat_postテーブルモデル
type ChatPost struct {
	ChatPostID string    `json:"id"`
	Room       int64     `json:"room"`
	UserID     int64     `json:"user"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}
