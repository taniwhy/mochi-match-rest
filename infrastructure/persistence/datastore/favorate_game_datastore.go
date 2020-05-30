package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/models/response"
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

func (eD favoriteGameDatastore) FindFavoriteGameByID(id string) ([]*response.FavoriteGamesRes, error) {
	f := []*response.FavoriteGamesRes{}
	eD.db.Table("favorite_games").
		Select("favorite_games.favorite_game_id, game_titles.game_title_id, game_titles.game_title").
		Joins("left join game_titles on game_titles.game_title_id = favorite_games.game_title_id").
		Where("favorite_games.user_id = ?", id).
		Scan(&f)
	return f, nil
}

func (eD favoriteGameDatastore) InsertFavoriteGame(favgame *dbmodel.FavoriteGame) error {
	return eD.db.Create(favgame).Error
}

func (eD favoriteGameDatastore) DeleteFavoriteGame(uID, fID string) error {
	f := dbmodel.FavoriteGame{}
	recordNotFound := eD.db.Where("user_id = ? AND game_title_id = ?", uID, fID).Delete(&f).RecordNotFound()
	if recordNotFound {
		return fmt.Errorf("Record not found : %v", uID)
	}
	return nil
}
