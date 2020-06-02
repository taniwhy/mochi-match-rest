package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// FavoriteGameRepository : ユーザーのリポジトリ
type FavoriteGameRepository interface {
	FindByID(id string) ([]*models.FavoriteGame, error)
	Insert(room *models.FavoriteGame) error
	Delete(uID, fID string) error
}
