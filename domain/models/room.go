package models

import (
	"time"
)

// Room : roomテーブルモデル
type Room struct {
	RoomID    string
	RoomOwner string
	GameTitle string
	Capacity  int
	IsLock    bool
	CreatedAt time.Time
}
