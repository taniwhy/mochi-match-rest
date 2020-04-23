package models

import (
	"time"
)

// Room : roomテーブルモデル
type Room struct {
	ID        int64
	RoomOwner int
	GameTitle int
	Capacity  int
	IsLock    bool
	CreatedAt time.Time
}
