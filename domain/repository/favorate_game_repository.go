package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/models/response"
)

// FavoriteGameRepository : ユーザーのリポジトリ
type FavoriteGameRepository interface {
	FindFavoriteGameByID(id string) ([]*response.FavoriteGamesRes, error)
	InsertFavoriteGame(room *dbmodel.FavoriteGame) error
	DeleteFavoriteGame(uID, fID string) error
}
