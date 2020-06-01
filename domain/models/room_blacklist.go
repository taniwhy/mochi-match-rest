package models

import (
	"time"
)

// RoomBlacklist : room_blacklistテーブルモデル
type RoomBlacklist struct {
	ID        int64
	Room      int64
	BlackUser int64
	CreatedAt time.Time
}
