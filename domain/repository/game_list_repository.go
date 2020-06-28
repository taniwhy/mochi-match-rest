//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// IGameListRepository : ゲームタイトルのリポジトリ
type IGameListRepository interface {
	FindAll() ([]*models.GameList, error)
	Insert(*models.GameList) error
	Update(*models.GameList) error
	Delete(*models.GameList) error
}
