package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// EntryHistoryUsecase :
type EntryHistoryUsecase interface {
	FindAllEntryHistory() ([]*models.EntryHistory, error)
	InsertEntryHistory(entryHistory *models.EntryHistory) error
	UpdateEntryHistory(entryHistory *models.EntryHistory) error
	DeleteEntryHistory(entryHistory *models.EntryHistory) error
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

func (eU entryHistoryUsecase) FindAllEntryHistory() ([]*models.EntryHistory, error) {
	entryHistories, err := eU.entryHistoryRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return entryHistories, nil
}

func (eU entryHistoryUsecase) InsertEntryHistory(entryHistory *models.EntryHistory) error {
	err := eU.entryHistoryRepository.Insert(entryHistory)
	if err != nil {
		return err
	}
	return nil
}

func (eU entryHistoryUsecase) UpdateEntryHistory(entryHistory *models.EntryHistory) error {
	err := eU.entryHistoryRepository.Update(entryHistory)
	if err != nil {
		return err
	}
	return nil
}

func (eU entryHistoryUsecase) DeleteEntryHistory(entryHistory *models.EntryHistory) error {
	err := eU.entryHistoryRepository.Delete(entryHistory)
	if err != nil {
		return err
	}
	return nil
}
