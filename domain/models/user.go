package models

import (
	"time"
)

// User : usersテーブルモデル
type User struct {
	ID         int64
	GoogleID   string
	FacebookID string
	TwitterID  string
	IsAdmin    bool
	IsFrozen   bool
	IsDelete   bool
	CreatedAt  time.Time
	UpdateAt   time.Time
}
