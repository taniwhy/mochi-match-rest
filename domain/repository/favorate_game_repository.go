package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/response"
)

// FavoriteGameRepository : ユーザーのリポジトリ
type FavoriteGameRepository interface {
	FindFavoriteGameByID(id string) ([]*response.FavoriteGamesRes, error)
	InsertFavoriteGame(room *models.FavoriteGame) error
	DeleteFavoriteGame(uID, fID string) error
}
