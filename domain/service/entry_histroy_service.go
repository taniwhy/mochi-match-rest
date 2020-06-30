//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package service

import (
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IEntryHistoryService : インターフェース
type IEntryHistoryService interface {
	CanJoin(userID string) (bool, error)
	CheckJoin(userID, roomID string) (bool, error)
}

type entryHistoryService struct {
	entryHistoryRepository repository.IEntryHistoryRepository
}

// NewEntryHistoryService : EntryHistoryServiceの生成
func NewEntryHistoryService(eR repository.IEntryHistoryRepository) IEntryHistoryService {
	return &entryHistoryService{
		entryHistoryRepository: eR,
	}
}

func (s *entryHistoryService) CanJoin(userID string) (bool, error) {
	res, err := s.entryHistoryRepository.FindNotLeave(userID)
	if err != nil {
		return false, err
	}
	if res != nil {
		return false, nil
	}
	return true, nil
}

func (s *entryHistoryService) CheckJoin(userID, roomID string) (bool, error) {
	res, err := s.entryHistoryRepository.FindNotLeaveByRoomID(userID, roomID)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}
	return true, nil
}
