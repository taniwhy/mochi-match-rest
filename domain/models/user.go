package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
<<<<<<< HEAD
	"github.com/taniwhy/mochi-match-rest/domain/errors"
=======
>>>>>>> 9a70d396a44f4ff618d89b4dafe030a585b76c6f
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

// NewUser :
<<<<<<< HEAD
func NewUser(email string) (*User, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &User{
=======
func NewUser(email) (User, err error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return User{
>>>>>>> 9a70d396a44f4ff618d89b4dafe030a585b76c6f
		UserID:     uid.String(),
		GoogleID:   sql.NullString{String: "", Valid: false},
		FacebookID: sql.NullString{String: "", Valid: false},
		TwitterID:  sql.NullString{String: "", Valid: false},
		Email:      email,
		IsAdmin:    false,
		IsFreeze:   false,
		IsDelete:   false,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
		DeleteAt:   sql.NullTime{Time: time.Now(), Valid: false},
	}, nil
}
