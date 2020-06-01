package models

import (
	"time"
)

// FavoriteGame : favorate_gameテーブルモデル
type FavoriteGame struct {
	FavoriteGameID string
	UserID         string
	GameTitleID    string
	CreatedAt      time.Time
}
