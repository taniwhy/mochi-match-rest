package models

import (
	"time"
)

// GameTitle : game_titleテーブルモデル
type GameTitle struct {
	GameTitleID string
	GameTitle   string `json:"game_title" binding:"required"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}
