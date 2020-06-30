//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package service

import (
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IRoomService : インターフェース
type IRoomService interface {
	CanInsert(id string) (bool, error)
	IsLock(id string) (bool, error)
	IsOwner(uid, rid string) (bool, error)
}

type roomService struct {
	roomRepository repository.IRoomRepository
}

// NewRoomService : RoomServiceの生成
func NewRoomService(rR repository.IRoomRepository) IRoomService {
	return &roomService{
		roomRepository: rR,
	}
}

func (s *roomService) CanInsert(id string) (bool, error) {
	r, err := s.roomRepository.FindUnlockByID(id)
	if err != nil {
		return false, err
	}
	if r != nil {
		return false, nil
	}
	return true, nil
}

func (s *roomService) IsLock(id string) (bool, error) {
	r, err := s.roomRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	if r.IsLock == true {
		return false, nil
	}
	return true, nil
}

func (s *roomService) IsOwner(uid, rid string) (bool, error) {
	r, err := s.roomRepository.FindByID(rid)
	if err != nil {
		return false, err
	}
	if r.UserID == uid {
		return true, nil
	}
	return false, nil
}
