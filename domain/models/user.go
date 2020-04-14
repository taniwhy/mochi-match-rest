package models

import (
	"time"
)

// User : usersテーブルモデル
type User struct {
	ID         int64
	Provider   string
	ProviderID string
	IsAdmin    bool
	IsFrozen   bool
	CreatedAt  time.Time
	UpdateAt   time.Time
}
