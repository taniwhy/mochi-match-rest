package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type gameTitleDatastore struct {
	db *gorm.DB
}

// NewGameTitleDatastore : .
func NewGameTitleDatastore(db *gorm.DB) repository.GameTitleRepository {
	return &gameTitleDatastore{db}
}

func (gD gameTitleDatastore) FindAllGameTitle() ([]*dbmodel.GameTitle, error) {
	gameTitle := []*dbmodel.GameTitle{}
	err := gD.db.Find(&gameTitle).Error
	if err != nil {
		return nil, err
	}
	return gameTitle, nil
}

func (gD gameTitleDatastore) InsertGameTitle(gameTitle *dbmodel.GameTitle) error {
	return gD.db.Create(gameTitle).Error
}

func (gD gameTitleDatastore) UpdateGameTitle(gT *dbmodel.GameTitle) error {
	return gD.db.Model(gT).Where("game_title_id = ?", gT.GameTitleID).Updates(gT).Error
}

func (gD gameTitleDatastore) DeleteGameTitle(gameTitle *dbmodel.GameTitle) error {
	err := gD.db.Take(&gameTitle).Error
	if err != nil {
		return err
	}
	return gD.db.Delete(gameTitle).Error
}
