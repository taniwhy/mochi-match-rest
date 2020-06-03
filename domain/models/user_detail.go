package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// UserDetail : user_detailテーブルモデル
type UserDetail struct {
	UserDetailID string
	UserID       string
	UserName     string `json:"user_name" binding:"required"`
	Icon         string `json:"icon" binding:"required"`
	UpdateAt     time.Time
}

// NewUserDetail :
func NewUserDetail(uid, name string) (*UserDetail, error) {
	udid, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &UserDetail{
		UserDetailID: udid.String(),
		UserID:       uid,
		UserName:     name,
		Icon:         "",
		UpdateAt:     time.Now(),
	}, nil
}
