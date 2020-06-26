package models

import (
	"time"

	"github.com/google/uuid"
)

// ChatPost : chat_postテーブルモデル
type ChatPost struct {
	ChatPostID string    `json:"id"`
	RoomID     string    `json:"room"`
	UserID     string    `json:"user"`
	Message    string    `json:"message" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
}

// NewChatPost :
func NewChatPost(rID, uID, m string) (*ChatPost, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &ChatPost{
		ChatPostID: id.String(),
		RoomID:     rID,
		UserID:     uID,
		Message:    m,
		CreatedAt:  time.Now(),
	}, nil
}
