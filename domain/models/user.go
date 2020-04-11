package models

import (
	"time"
)

// User : usersテーブルモデル
type User struct {
	ID        int
	IsAdmin   bool
	IsFrozen  bool
	CreatedAt time.Time
	UpdateAt  time.Time
}
