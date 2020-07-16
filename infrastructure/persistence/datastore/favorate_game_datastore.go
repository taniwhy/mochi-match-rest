package datastore

import (
	"github.com/jinzhu/gorm"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type favoriteGameDatastore struct {
	db *gorm.DB
}

// NewFavoriteGameDatastore : お気に入りゲームデータストアの生成
func NewFavoriteGameDatastore(db *gorm.DB) repository.IFavoriteGameRepository {
	return &favoriteGameDatastore{db}
}

func (d *favoriteGameDatastore) FindByID(userID string) ([]*output.FavoriteGamesRes, error) {
	games := []*output.FavoriteGamesRes{}
	err := d.db.
		Table("favorite_games").
		Select(`
		favorite_games.game_title,,
		game_lists.game_list_id,
		favorite_games.created_at
		`).
		Joins("LEFT JOIN game_lists ON game_lists.game_title = favorite_games.game_title").
		Where(`favorite_games.user_id = ?`, userID).
		First(&games).Error
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
