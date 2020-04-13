package models

import (
	"time"
)

// EntryHistory : entry_historyテーブルモデル
type EntryHistory struct {
	ID        int
	Room      int
	User      int
	
	EntryTime time.Time
	UpdateAt  time.Time
}
