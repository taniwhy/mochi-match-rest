package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type roomBlacklistDatastore struct {
	db *gorm.DB
}

// NewRoomBlacklistDatastore : UserPersistenseを生成.
func NewRoomBlacklistDatastore(db *gorm.DB) repository.RoomBlacklistRepository {
	return &roomBlacklistDatastore{db}
}

func (rD roomBlacklistDatastore) FindAllBlacklist() ([]*dbmodel.RoomBlacklist, error) {
	blacklist := []*dbmodel.RoomBlacklist{}

	err := rD.db.Find(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return blacklist, nil
}

func (rD roomBlacklistDatastore) FindBlacklistByID(id int64) (*dbmodel.RoomBlacklist, error) {
	blacklist := dbmodel.RoomBlacklist{ID: id}
	err := rD.db.Take(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return &blacklist, nil
}

func (rD roomBlacklistDatastore) InsertBlacklist(blacklist *dbmodel.RoomBlacklist) error {
	return rD.db.Create(blacklist).Error
}

func (rD roomBlacklistDatastore) DeleteBlacklist(blacklist *dbmodel.RoomBlacklist) error {
	err := rD.db.Take(&blacklist).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(blacklist).Error
}
