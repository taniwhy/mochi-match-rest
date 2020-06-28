package models

import (
	"database/sql"
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// EntryHistory : entry_historyテーブルモデル
type EntryHistory struct {
	EntryHistoryID string
	UserID         string
	RoomID         string
	IsLeave        bool
	CreatedAt      time.Time
	LeavedAt       sql.NullTime
}

// NewEntryHistory : entry_historyテーブルのレコードモデル生成
func NewEntryHistory(uid, rid string) (*EntryHistory, error) {
	return &EntryHistory{
		EntryHistoryID: uuid.UuID(),
		UserID:         uid,
		RoomID:         rid,
		IsLeave:        false,
		CreatedAt:      clock.Now(),
		LeavedAt:       sql.NullTime{Time: clock.Now(), Valid: false},
	}, nil
}
