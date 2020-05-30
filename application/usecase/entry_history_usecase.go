package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// EntryHistoryUsecase :
type EntryHistoryUsecase interface {
	FindAllEntryHistory() ([]*dbmodel.EntryHistory, error)
	InsertEntryHistory(entryHistory *dbmodel.EntryHistory) error
	UpdateEntryHistory(entryHistory *dbmodel.EntryHistory) error
	DeleteEntryHistory(entryHistory *dbmodel.EntryHistory) error
}

type entryHistoryUsecase struct {
	entryHistoryRepository repository.EntryHistoryRepository
}

// NewEntryHistoryUsecase :
func NewEntryHistoryUsecase(eR repository.EntryHistoryRepository) EntryHistoryUsecase {
	return &entryHistoryUsecase{
		entryHistoryRepository: eR,
	}
}

func (eU entryHistoryUsecase) FindAllEntryHistory() ([]*dbmodel.EntryHistory, error) {
	entryHistories, err := eU.entryHistoryRepository.FindAllEntryHistory()
	if err != nil {
		return nil, err
	}
	return entryHistories, nil
}

func (eU entryHistoryUsecase) InsertEntryHistory(entryHistory *dbmodel.EntryHistory) error {
	err := eU.entryHistoryRepository.InsertEntryHistory(entryHistory)
	if err != nil {
		return err
	}
	return nil
}

func (eU entryHistoryUsecase) UpdateEntryHistory(entryHistory *dbmodel.EntryHistory) error {
	err := eU.entryHistoryRepository.InsertEntryHistory(entryHistory)
	if err != nil {
		return err
	}
	return nil
}

func (eU entryHistoryUsecase) DeleteEntryHistory(entryHistory *dbmodel.EntryHistory) error {
	err := eU.entryHistoryRepository.DeleteEntryHistory(entryHistory)
	if err != nil {
		return err
	}
	return nil
}
