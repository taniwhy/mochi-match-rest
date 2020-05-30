package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type entryHistoryDatastore struct {
	db *gorm.DB
}

// NewEntryHistoryDatastore :
func NewEntryHistoryDatastore(db *gorm.DB) repository.EntryHistoryRepository {
	return &entryHistoryDatastore{db}
}

func (eD entryHistoryDatastore) FindAllEntryHistory() ([]*dbmodel.EntryHistory, error) {
	entryHistorys := []*dbmodel.EntryHistory{}

	err := eD.db.Find(&entryHistorys).Error
	if err != nil {
		return nil, err
	}
	return entryHistorys, nil
}

func (eD entryHistoryDatastore) InsertEntryHistory(entryHistory *dbmodel.EntryHistory) error {
	return eD.db.Create(entryHistory).Error
}

func (eD entryHistoryDatastore) UpdateEntryHistory(entryHistory *dbmodel.EntryHistory) error {
	return eD.db.Update(entryHistory).Error
}

func (eD entryHistoryDatastore) DeleteEntryHistory(entryHistory *dbmodel.EntryHistory) error {
	err := eD.db.Take(&entryHistory).Error
	if err != nil {
		return err
	}
	return eD.db.Delete(entryHistory).Error
}
