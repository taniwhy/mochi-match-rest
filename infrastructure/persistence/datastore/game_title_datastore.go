package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type gameTitleDatastore struct {
	db *gorm.DB
}

// NewGameTitleDatastore : .
func NewGameTitleDatastore(db *gorm.DB) repository.ChatPostRepository {
	return &chatPostDatastore{db}
}

func (gD gameTitleDatastore) FindAllGameTitle() ([]*models.GameTitle, error) {
	gameTitle := []*models.GameTitle{}

	err := gD.db.Find(&gameTitle).Error
	if err != nil {
		return nil, err
	}
	return gameTitle, nil
}

func (gD gameTitleDatastore) InsertGameTitle(gameTitle *models.GameTitle) error {
	return gD.db.Create(gameTitle).Error
}

func (gD roomDatastore) UpdateGameTitle(gameTitle *models.GameTitle) error {
	return gD.db.Updates(gameTitle).Error
}

func (gD roomDatastore) DeleteGameTitle(gameTitle *models.GameTitle) error {
	err := gD.db.Take(&gameTitle).Error
	if err != nil {
		return err
	}
	return gD.db.Delete(gameTitle).Error
}
