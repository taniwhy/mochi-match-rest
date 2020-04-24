package models

import (
	"time"
)

// Room : roomテーブルモデル
type Room struct {
	ID        int64
	RoomOwner int64
	GameTitle int64
	Capacity  int
	IsLock    bool
	CreatedAt time.Time
}
