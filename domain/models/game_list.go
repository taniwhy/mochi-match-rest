package models

import (
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// GameList : game_titleテーブルモデル
type GameList struct {
	GameListID string    `json:"id" binding:"required"`
	GameTitle  string    `json:"game_title" binding:"required"`
	CreatedAt  time.Time `json:"created_at" binding:"required"`
	UpdateAt   time.Time `json:"update_at" binding:"required"`
}

// NewGameList : game_listsテーブルのレコードモデル生成
func NewGameList(gameTitle string) (*GameList, error) {
	return &GameList{
		GameListID: uuid.UuID(),
		GameTitle:  gameTitle,
		CreatedAt:  clock.Now(),
		UpdateAt:   clock.Now(),
	}, nil
}
