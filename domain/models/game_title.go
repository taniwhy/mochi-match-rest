package models

import (
	"time"
)

// GameTitle : game_titleテーブルモデル
type GameTitle struct {
	ID        int
	GameTitle int
	CreatedAt time.Time
	UpdateAt  time.Time
}
