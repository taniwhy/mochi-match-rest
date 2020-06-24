package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// GameList : game_titleテーブルモデル
type GameList struct {
	GameListID string
	GameTitle  string `json:"game_title" binding:"required"`
	CreatedAt  time.Time
	UpdateAt   time.Time
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
