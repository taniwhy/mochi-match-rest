package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
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
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &EntryHistory{
		EntryHistoryID: id.String(),
		UserID:         uid,
		RoomID:         rid,
		IsLeave:        false,
		CreatedAt:      time.Now(),
		LeavedAt:       sql.NullTime{Time: time.Now(), Valid: false},
	}, nil
}
