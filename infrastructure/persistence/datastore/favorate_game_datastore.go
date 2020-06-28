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

func (eD favoriteGameDatastore) FindByID(id string) ([]*models.FavoriteGame, error) {
	f := []*models.FavoriteGame{}
	err := eD.db.Where("user_id = ?", id).Find(&f).Error
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
	err := eD.db.Where("user_id = ? AND game_title = ?", uID, gT).Delete(&f).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err}
	}
	return nil
}
