//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// FavoriteGameRepository : お気に入りゲームのリポジトリ
type FavoriteGameRepository interface {
	FindByID(id string) ([]*models.FavoriteGame, error)
	Insert(room *models.FavoriteGame) error
	Delete(uID, fID string) error
}
