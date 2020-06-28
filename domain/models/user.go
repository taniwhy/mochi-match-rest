package models

import (
	"database/sql"
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
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

// NewUser : usersテーブルのレコードモデル生成
func NewUser(email string) (*User, error) {
	return &User{
		UserID:     uuid.UuID(),
		GoogleID:   sql.NullString{String: "", Valid: false},
		FacebookID: sql.NullString{String: "", Valid: false},
		TwitterID:  sql.NullString{String: "", Valid: false},
		Email:      email,
		IsAdmin:    false,
		IsFreeze:   false,
		IsDelete:   false,
		CreatedAt:  clock.Now(),
		UpdateAt:   clock.Now(),
		DeleteAt:   sql.NullTime{Time: clock.Now(), Valid: false},
	}, nil
}
