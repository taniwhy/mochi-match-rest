package models

import (
	"time"
)

// ChatPost : chat_postテーブルモデル
type ChatPost struct {
	ID        int64
	Room      int64
	User      int64
	Message   string
	CreatedAt time.Time
}
