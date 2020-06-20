package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// GameTitleRepository : ゲームタイトルのリポジトリ
type GameTitleRepository interface {
	FindAllGameTitle() ([]*models.GameTitle, error)
	InsertGameTitle(gameTitle *models.GameTitle) error
	UpdateGameTitle(gameTitle *models.GameTitle) error
	DeleteGameTitle(gameTitle *models.GameTitle) error
}
