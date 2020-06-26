package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// GameList : game_titleテーブルモデル
type GameList struct {
	GameListID string    `json:"id" binding:"required"`
	GameTitle  string    `json:"game_title" binding:"required"`
	CreatedAt  time.Time `json:"created_at" binding:"required"`
	UpdateAt   time.Time `json:"update_at" binding:"required"`
}

// NewGameList : game_listsテーブルのレコードモデル生成
func NewGameList(gT string) (*GameList, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &GameList{
		GameListID: id.String(),
		GameTitle:  gT,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}, nil
}
