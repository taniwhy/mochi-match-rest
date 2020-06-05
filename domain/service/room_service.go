package service

import (
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IRoomService : インターフェース
type IRoomService interface {
	CanInsert(id string) (bool, error)
}

type roomService struct {
	roomRepository repository.RoomRepository
}

// NewRoomService : UserServiceの生成
func NewRoomService(rR repository.RoomRepository) IRoomService {
	return &roomService{
		roomRepository: rR,
	}
}

func (rS roomService) CanInsert(id string) (bool, error) {
	r, err := rS.roomRepository.FindUnlockByID(id)
	if err != nil {
		return false, err
	}
	if r != nil {
		return false, nil
	}
	return true, nil
}
