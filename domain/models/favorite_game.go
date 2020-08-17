package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// FavoriteGame : favorate_gameテーブルモデル
type FavoriteGame struct {
	FavoriteGameID string
	UserID         string
	GameTitle      string
	CreatedAt      time.Time
}

// NewFavoriteGame : favorate_gameテーブルのレコードモデル生成
func NewFavoriteGame(userID, gameTitle string) (*FavoriteGame, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &FavoriteGame{
		FavoriteGameID: id.String(),
		UserID:         userID,
		GameTitle:      gameTitle,
		CreatedAt:      time.Now(),
	}, nil
}
