//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IEntryHistoryUsecase : インターフェース
type IEntryHistoryUsecase interface {
	FindAllEntryHistory() ([]*models.EntryHistory, error)
	InsertEntryHistory(entryHistory *models.EntryHistory) error
	UpdateEntryHistory(entryHistory *models.EntryHistory) error
	DeleteEntryHistory(entryHistory *models.EntryHistory) error
}

type entryHistoryUsecase struct {
	entryHistoryRepository repository.IEntryHistoryRepository
}

// NewEntryHistoryUsecase : EntryHistoryユースケースの生成
func NewEntryHistoryUsecase(eR repository.IEntryHistoryRepository) IEntryHistoryUsecase {
	return &entryHistoryUsecase{entryHistoryRepository: eR}
}

func (u *entryHistoryUsecase) FindAllEntryHistory() ([]*models.EntryHistory, error) {
	entryHistories, err := u.entryHistoryRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return entryHistories, nil
}

func (u *entryHistoryUsecase) InsertEntryHistory(entryHistory *models.EntryHistory) error {
	err := u.entryHistoryRepository.Insert(entryHistory)
	if err != nil {
		return err
	}
	return nil
}

func (u *entryHistoryUsecase) UpdateEntryHistory(entryHistory *models.EntryHistory) error {
	err := u.entryHistoryRepository.Update(entryHistory)
	if err != nil {
		return err
	}
	return nil
}

func (u *entryHistoryUsecase) DeleteEntryHistory(entryHistory *models.EntryHistory) error {
	err := u.entryHistoryRepository.Delete(entryHistory)
	if err != nil {
		return err
	}
	return nil
}
