package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"gopkg.in/guregu/null.v4"
)

// Room : roomテーブルモデル
type Room struct {
	RoomID     string
	UserID     string
	RoomText   string
	GameListID string
	GameHardID string
	Capacity   int
	Start      null.Time
	IsLock     bool
	CreatedAt  time.Time
	UpdateAt   time.Time
}

// NewRoom : roomテーブルのレコードモデル生成
func NewRoom(uid, text, glid, ghid string, cap int, s time.Time) (*Room, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	r := &Room{
		RoomID:     id.String(),
		UserID:     uid,
		RoomText:   text,
		GameListID: glid,
		GameHardID: ghid,
		Capacity:   cap,
		IsLock:     false,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}
	if s.IsZero() == true {
		r.Start = null.NewTime(s, false)
	} else {
		r.Start = null.NewTime(s, true)
	}
	return r, nil
}
