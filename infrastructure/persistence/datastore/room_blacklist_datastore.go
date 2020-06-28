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

func (rD roomBlacklistDatastore) FindByRoomID(rid string) ([]*models.RoomBlacklist, error) {
	rB := []*models.RoomBlacklist{}
	err := rD.db.Where("room_id = ?", rid).Find(&rB).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return rB, nil
}

func (rD roomBlacklistDatastore) Insert(rB *models.RoomBlacklist) error {
	err := rD.db.Create(rB).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (rD roomBlacklistDatastore) Delete(rid string) error {
	rB := models.RoomBlacklist{}
	err := rD.db.Where("room_id = ?", rid).Delete(rB).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}
