package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// FavorateGameRepository : ユーザーのリポジトリ
type FavorateGameRepository interface {
	FindFavorateGameByID(id string) (*models.FavorateGame, error)
	InsertFavorateGame(room *models.FavorateGame) error
	DeleteFavorateGame(room *models.FavorateGame) error
}
