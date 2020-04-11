package models

import (
	"time"
)

// RoomBlacklist : room_blacklistテーブルモデル
type RoomBlacklist struct {
	ID        int
	Room      int
	BlackUser int
	CreatedAt time.Time
}
