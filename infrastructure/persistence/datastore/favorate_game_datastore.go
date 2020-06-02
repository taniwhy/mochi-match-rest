package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type favoriteGameDatastore struct {
	db *gorm.DB
}

type usersFavoriteGamesResBody struct {
	UserID        string
	FavoriteGames []favoriteGameDatastore
}

// NewFavoriteGameDatastore :
func NewFavoriteGameDatastore(db *gorm.DB) repository.FavoriteGameRepository {
	return &favoriteGameDatastore{db}
}

func (eD favoriteGameDatastore) FindByID(id string) ([]*models.FavoriteGame, error) {
	f := []*models.FavoriteGame{}
	err := eD.db.Table("favorite_games").
		Where("user_id = ?", id).
		Scan(&f).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return f, nil
}

func (eD favoriteGameDatastore) Insert(favgame *models.FavoriteGame) error {
	return eD.db.Create(favgame).Error
}

func (eD favoriteGameDatastore) Delete(uID, gT string) error {
	f := models.FavoriteGame{}
	recordNotFound := eD.db.Where("user_id = ? AND game_title = ?", uID, gT).Delete(&f).RecordNotFound()
	if recordNotFound {
		return fmt.Errorf("Record not found : %v", uID)
	}
	return nil
}
