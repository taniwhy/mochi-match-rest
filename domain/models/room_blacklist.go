package models

import (
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// RoomBlacklist : room_blacklistテーブルモデル
type RoomBlacklist struct {
	RoomBlacklistID string
	RoomID          string
	BlackUser       string
	CreatedAt       time.Time
}

// NewBlacklist : ブラックリストレコードの生成
func NewBlacklist(roomID, userID string) (*RoomBlacklist, error) {
	blacklist := &RoomBlacklist{
		RoomBlacklistID: uuid.UuID(),
		RoomID:          roomID,
		BlackUser:       userID,
		CreatedAt:       clock.Now(),
	}
	return blacklist, nil
}
