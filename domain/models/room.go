package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// Room : roomテーブルモデル
type Room struct {
	RoomID      string
	UserID      string
	RoomText    string
	GameTitleID string
	GameHardID  string
	Capacity    int
	IsLock      bool
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// NewRoom :
func NewRoom(uid, text, gtid, ghid string, cap int) (*Room, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &Room{
		RoomID:      id.String(),
		UserID:      uid,
		RoomText:    text,
		GameTitleID: gtid,
		GameHardID:  ghid,
		Capacity:    cap,
		IsLock:      false,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}, nil
}
