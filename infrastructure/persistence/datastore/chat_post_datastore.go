package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type chatPostDatastore struct {
	db *gorm.DB
}

// NewChatPostDatastore : UserPersistenseを生成.
func NewChatPostDatastore(db *gorm.DB) repository.ChatPostRepository {
	return &chatPostDatastore{db}
}

func (cD chatPostDatastore) FindAllChatPost() ([]*dbmodel.ChatPost, error) {
	chatposts := []*dbmodel.ChatPost{}

	err := cD.db.Find(&chatposts).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cD chatPostDatastore) FindChatPostByRoomID(roomID string) ([]*dbmodel.ChatPost, error) {
	chatposts := []*dbmodel.ChatPost{}
	err := cD.db.Order("created_at desc").Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cD chatPostDatastore) FindChatPostByRoomIDAndLimit(roomID string, limit int) ([]*dbmodel.ChatPost, error) {
	chatposts := []*dbmodel.ChatPost{}
	err := cD.db.Order("created_at desc").Limit(limit).Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cD chatPostDatastore) FindChatPostByRoomIDAndOffset(roomID, offset string) ([]*dbmodel.ChatPost, error) {
	chatposts := []*dbmodel.ChatPost{}
	err := cD.db.Order("created_at desc").Where("created_at < ?", offset).Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cD chatPostDatastore) FindChatPostByRoomIDAndLimitAndOffset(roomID, offset string, limit int) ([]*dbmodel.ChatPost, error) {
	chatposts := []*dbmodel.ChatPost{}
	err := cD.db.Order("created_at desc").Limit(limit).Where("created_at < ?", offset).Find(&chatposts, "room_id=?", roomID).Error
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cD chatPostDatastore) InsertChatPost(chatpost *dbmodel.ChatPost) error {
	return cD.db.Create(chatpost).Error
}

func (cD chatPostDatastore) DeleteChatPost(chatpost *dbmodel.ChatPost) error {
	err := cD.db.Take(&chatpost).Error
	if err != nil {
		return err
	}
	return cD.db.Delete(chatpost).Error
}
