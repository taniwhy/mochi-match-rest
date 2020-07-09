package datastore

import (
	"github.com/jinzhu/gorm"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type favoriteGameDatastore struct {
	db *gorm.DB
}

// NewFavoriteGameDatastore : お気に入りゲームデータストアの生成
func NewFavoriteGameDatastore(db *gorm.DB) repository.IFavoriteGameRepository {
	return &favoriteGameDatastore{db}
}

func (d *favoriteGameDatastore) FindByID(id string) ([]*models.FavoriteGame, error) {
	games := []*models.FavoriteGame{}
	err := d.db.Where("user_id = ?", id).Find(&games).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return games, nil
}

func (d *favoriteGameDatastore) Insert(favgame *models.FavoriteGame) error {
	return d.db.Create(favgame).Error
}

func (d *favoriteGameDatastore) Delete(uID, gT string) error {
	g := models.FavoriteGame{}
	err := d.db.Where("user_id = ? AND game_title = ?", uID, gT).Delete(&g).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err}
	}
	return nil
}
