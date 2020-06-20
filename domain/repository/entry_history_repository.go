package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// EntryHistoryRepository : 入室履歴のリポジトリ
type EntryHistoryRepository interface {
	FindAll() ([]*models.EntryHistory, error)
	Insert(*models.EntryHistory) error
	Update(*models.EntryHistory) error
	Delete(*models.EntryHistory) error
	CountEntryUser(string) (int, error)
	CheckEntry(rid, uid string) (bool, error)
	LeaveFlg(rid, uid string) error
}
