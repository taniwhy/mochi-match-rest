package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type favorateGameDatastore struct {
	db *gorm.DB
}

// NewFavorateGameDatastore :
func NewFavorateGameDatastore(db *gorm.DB) repository.EntryHistoryRepository {
	return &entryHistoryDatastore{db}
}

func (eD favorateGameDatastore) FindFavorateGameByID(id string) ([]*models.FavorateGame, error) {
	favorateGames := []*models.FavorateGame{}

	err := eD.db.Find(&favorateGames).Error
	if err != nil {
		return nil, err
	}
	return favorateGames, nil
}

func (eD favorateGameDatastore) InsertFavorateGame(favgame *models.FavorateGame) error {
	return eD.db.Create(favgame).Error
}

func (eD favorateGameDatastore) DeleteFavorateGame(favgame *models.FavorateGame) error {
	err := eD.db.Take(&favgame).Error
	if err != nil {
		return err
	}
	return eD.db.Delete(favgame).Error
}
