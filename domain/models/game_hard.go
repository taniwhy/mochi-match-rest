package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
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
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &GameHard{
		GameHardID: id.String(),
		HardName:   hN,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}, nil
}
