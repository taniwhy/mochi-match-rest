package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// GameTitleRepository : ユーザーのリポジトリ
type GameTitleRepository interface {
	FindAllGameTitle() ([]*dbmodel.GameTitle, error)
	InsertGameTitle(gameTitle *dbmodel.GameTitle) error
	UpdateGameTitle(gameTitle *dbmodel.GameTitle) error
	DeleteGameTitle(gameTitle *dbmodel.GameTitle) error
}
