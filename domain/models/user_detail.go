package models

import (
	"time"
)

// UserDetail : user_detailテーブルモデル
type UserDetail struct {
	UserDetailID string
	UserID       string
	UserName     string `json:"user_name" binding:"required"`
	Icon         int
	UpdateAt     time.Time
}
