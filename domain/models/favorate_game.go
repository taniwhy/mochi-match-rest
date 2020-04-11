package models

import (
	"time"
)

// FavorateGame : favorate_gameテーブルモデル
type FavorateGame struct {
	ID         int
	UserDetail int
	GameTitle  int
	CreatedAt  time.Time
}
