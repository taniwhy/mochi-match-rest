//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// IGameHardRepository : ゲームハードのリポジトリ
type IGameHardRepository interface {
	FindAll() ([]*models.GameHard, error)
	Insert(*models.GameHard) error
	Update(*models.GameHard) error
	Delete(*models.GameHard) error
}
