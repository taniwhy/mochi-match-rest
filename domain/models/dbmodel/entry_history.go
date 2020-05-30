package dbmodel

import (
	"time"
)

// EntryHistory : entry_historyテーブルモデル
type EntryHistory struct {
	ID        int
	Room      int
	User      int
	IsLeave   bool
	EntryTime time.Time
	LeaveTime time.Time
	CreatedAt time.Time
}
