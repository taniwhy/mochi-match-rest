package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// EntryHistoryRepository :
type EntryHistoryRepository interface {
	FindAllEntryHistory() ([]*models.EntryHistory, error)
	InsertEntryHistory(entryHistory *models.EntryHistory) error
	UpdateEntryHistory(entryHistory *models.EntryHistory) error
	DeleteEntryHistory(entryHistory *models.EntryHistory) error
}
