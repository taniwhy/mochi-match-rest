package models

import (
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// GameHard : game_hardsテーブルモデル
type GameHard struct {
	GameHardID string    `json:"id" binding:"required"`
	HardName   string    `json:"hard_name" binding:"required"`
	CreatedAt  time.Time `json:"created_at" binding:"required"`
	UpdateAt   time.Time `json:"update_at" binding:"required"`
}

// NewGameHard : game_hardsテーブルのレコードモデル生成
func NewGameHard(hN string) (*GameHard, error) {
	return &GameHard{
		GameHardID: uuid.UuID(),
		HardName:   hN,
		CreatedAt:  clock.Now(),
		UpdateAt:   clock.Now(),
	}, nil
}
