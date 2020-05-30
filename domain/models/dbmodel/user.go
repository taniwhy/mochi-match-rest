package dbmodel

import (
	"database/sql"
	"time"
)

// User : usersテーブルモデル
type User struct {
	UserID     string
	GoogleID   sql.NullString
	FacebookID sql.NullString
	TwitterID  sql.NullString
	Email      string `json:"email" binding:"required"`
	IsAdmin    bool
	IsFreeze   bool
	IsDelete   bool
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeleteAt   sql.NullTime
}
