//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// GameHardRepository : ゲームハードのリポジトリ
type GameHardRepository interface {
	FindAll() ([]*models.GameHard, error)
	Insert(*models.GameHard) error
	Update(*models.GameHard) error
	Delete(*models.GameHard) error
}
