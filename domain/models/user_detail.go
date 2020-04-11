package models

import (
	"time"
)

// UserDetail : user_detailテーブルモデル
type UserDetail struct {
	ID       int
	User     int
	UserName string
	Icon     int
	UpdateAt time.Time
}
