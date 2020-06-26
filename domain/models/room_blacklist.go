package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// RoomBlacklist : room_blacklistテーブルモデル
type RoomBlacklist struct {
	RoomBlacklistID string
	RoomID          string
	BlackUser       string
	CreatedAt       time.Time
}

// NewBlacklist :
func NewBlacklist(rid, uid string) (*RoomBlacklist, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	r := &RoomBlacklist{
		RoomBlacklistID: id.String(),
		RoomID:          rid,
		BlackUser:       uid,
		CreatedAt:       time.Now(),
	}
	return r, nil
}
