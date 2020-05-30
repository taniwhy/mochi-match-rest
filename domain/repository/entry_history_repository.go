package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// EntryHistoryRepository :
type EntryHistoryRepository interface {
	FindAllEntryHistory() ([]*dbmodel.EntryHistory, error)
	InsertEntryHistory(entryHistory *dbmodel.EntryHistory) error
	UpdateEntryHistory(entryHistory *dbmodel.EntryHistory) error
	DeleteEntryHistory(entryHistory *dbmodel.EntryHistory) error
}
