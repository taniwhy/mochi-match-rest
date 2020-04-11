package models

import (
	"time"
)

// ChatPost : chat_postテーブルモデル
type ChatPost struct {
	ID        int
	Room      int
	User      int
	Message   string
	CreatedAt time.Time
}
