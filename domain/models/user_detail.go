package models

import (
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// UserDetail : user_detailテーブルモデル
type UserDetail struct {
	UserDetailID string
	UserID       string
	UserName     string `json:"user_name" binding:"required"`
	Icon         string `json:"icon" binding:"required"`
	UpdateAt     time.Time
}

// NewUserDetail : user_detailテーブルのレコードモデル生成
func NewUserDetail(uid, name string) (*UserDetail, error) {
	return &UserDetail{
		UserDetailID: uuid.UuID(),
		UserID:       uid,
		UserName:     name,
		Icon:         "icon",
		UpdateAt:     clock.Now(),
	}, nil
}
