package models

import (
	"time"
)

// GameTitle : game_titleテーブルモデル
type GameTitle struct {
	ID        int
	GameTitle string
	CreatedAt time.Time
	UpdateAt  time.Time
}
