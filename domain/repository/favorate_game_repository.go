//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
)

// IFavoriteGameRepository : お気に入りゲームのリポジトリ
type IFavoriteGameRepository interface {
	FindByID(id string) ([]*output.FavoriteGamesRes, error)
	Insert(room *models.FavoriteGame) error
	Delete(uID, fID string) error
}
