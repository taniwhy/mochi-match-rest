package models

import (
	"fmt"
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
	"gopkg.in/guregu/null.v4"
)

// Room : roomテーブルモデル
type Room struct {
	RoomID     string    `json:"id" binding:"required"`
	UserID     string    `json:"user_id" binding:"required"`
	RoomText   string    `json:"room_text" binding:"required"`
	GameListID string    `json:"game_list_id" binding:"required"`
	GameHardID string    `json:"game_hard_id" binding:"required"`
	Capacity   int       `json:"capacity" binding:"required"`
	Start      null.Time `json:"start_at" binding:"required"`
	IsLock     bool      `json:"is_lock" binding:"required"`
	CreatedAt  time.Time `json:"created_at" binding:"required"`
	UpdateAt   time.Time `json:"update_at" binding:"required"`
}

// NewRoom : roomテーブルのレコードモデル生成
func NewRoom(uid, text, glid, ghid string, cap int, s time.Time) (*Room, error) {
	r := &Room{
		RoomID:     uuid.UuID(),
		UserID:     uid,
		RoomText:   text,
		GameListID: glid,
		GameHardID: ghid,
		Capacity:   cap,
		IsLock:     false,
		CreatedAt:  clock.Now(),
		UpdateAt:   clock.Now(),
	}
	fmt.Println(clock.Now())
	if s.IsZero() == true {
		r.Start = null.NewTime(s, false)
	} else {
		r.Start = null.NewTime(s, true)
	}
	return r, nil
}
