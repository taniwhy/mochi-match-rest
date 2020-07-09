package datastore

import (
	"github.com/jinzhu/gorm"

	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type chatPostDatastore struct {
	db *gorm.DB
}

// NewChatPostDatastore : チャット投稿データストアの生成
func NewChatPostDatastore(db *gorm.DB) repository.IChatPostRepository {
	return &chatPostDatastore{db}
}

func (d *chatPostDatastore) FindByRoomID(roomID string) ([]*models.ChatPost, error) {
	chatposts := []*models.ChatPost{}
	err := d.db.Order("created_at desc").Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (d *chatPostDatastore) FindByRoomIDAndLimit(roomID string, limit int) ([]*models.ChatPost, error) {
	chatposts := []*models.ChatPost{}
	err := d.db.Order("created_at desc").Limit(limit).Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (d *chatPostDatastore) FindByRoomIDAndOffset(roomID, offset string) ([]*models.ChatPost, error) {
	chatposts := []*models.ChatPost{}
	err := d.db.Order("created_at desc").Where("created_at < ?", offset).Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (d *chatPostDatastore) FindByRoomIDAndLimitAndOffset(roomID, offset string, limit int) ([]*models.ChatPost, error) {
	chatposts := []*models.ChatPost{}
	err := d.db.Order("created_at desc").Limit(limit).Where("created_at < ?", offset).Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (d *chatPostDatastore) Insert(chatpost *models.ChatPost) (*output.ChatPostResBody, error) {
	err := d.db.Create(chatpost).Error
	if err != nil {
		return nil, err
	}
	chatpostRes := &output.ChatPostResBody{}
	err = d.db.
		Table("chat_posts").
		Select(`
			chat_posts.chat_post_id,
			chat_posts.room_id,
			chat_posts.user_id,
			user_details.user_name,
			user_details.icon,
			chat_posts.message,
			chat_posts.created_at
			`).
		Joins("LEFT JOIN user_details ON chat_posts.user_id = user_details.user_id").
		Where("chat_posts.chat_post_id = ?", chatpost.ChatPostID).Scan(&chatpostRes).Error
	return chatpostRes, nil
}

func (d *chatPostDatastore) Delete(chatpost *models.ChatPost) error {
	err := d.db.Take(&chatpost).Error
	if err != nil {
		return err
	}
	return d.db.Delete(chatpost).Error
}
