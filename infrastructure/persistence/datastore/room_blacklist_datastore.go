package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type roomBlacklistDatastore struct {
	db *gorm.DB
}

// NewRoomBlacklistDatastore : ルームブラックリストデータストアの生成
func NewRoomBlacklistDatastore(db *gorm.DB) repository.IRoomBlacklistRepository {
	return &roomBlacklistDatastore{db}
}

func (d *roomBlacklistDatastore) FindByRoomID(roomID string) ([]*models.RoomBlacklist, error) {
	blacklists := []*models.RoomBlacklist{}
	err := d.db.Where("room_id = ?", roomID).Find(&blacklists).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return blacklists, nil
}

func (d *roomBlacklistDatastore) Insert(blacklist *models.RoomBlacklist) error {
	err := d.db.Create(blacklist).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (d *roomBlacklistDatastore) Delete(roomID, userID string) error {
	rB := models.RoomBlacklist{}
	err := d.db.Where("room_id = ? AND user_id = ?", roomID, userID).Delete(rB).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}
