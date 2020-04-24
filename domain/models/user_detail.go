package models

import (
	"time"
)

// UserDetail : user_detailテーブルモデル
type UserDetail struct {
	ID       int64
	User     int64
	UserName string
	Icon     int
	UpdateAt time.Time
}
