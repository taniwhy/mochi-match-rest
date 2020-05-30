package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// RoomBlacklistUseCase :
type RoomBlacklistUseCase interface {
	FindAllBlacklist() ([]*dbmodel.RoomBlacklist, error)
	FindBlacklistByID(id int64) (*dbmodel.RoomBlacklist, error)
	InsertBlacklist(roomReservation *dbmodel.RoomBlacklist) error
	DeleteBlacklist(roomReservation *dbmodel.RoomBlacklist) error
}

type roomBlacklistUsecase struct {
	roomBlacklistRepository repository.RoomBlacklistRepository
}

// NewRoomBlacklistUsecase :
func NewRoomBlacklistUsecase(rR repository.RoomBlacklistRepository) RoomBlacklistUseCase {
	return &roomBlacklistUsecase{
		roomBlacklistRepository: rR,
	}
}

func (rU roomBlacklistUsecase) FindAllBlacklist() ([]*dbmodel.RoomBlacklist, error) {
	blacklists, err := rU.roomBlacklistRepository.FindAllBlacklist()
	if err != nil {
		return nil, err
	}
	return blacklists, nil
}

func (rU roomBlacklistUsecase) FindBlacklistByID(id int64) (*dbmodel.RoomBlacklist, error) {
	blacklist, err := rU.roomBlacklistRepository.FindBlacklistByID(id)
	if err != nil {
		return nil, err
	}
	return blacklist, nil
}

func (rU roomBlacklistUsecase) InsertBlacklist(blacklist *dbmodel.RoomBlacklist) error {
	err := rU.roomBlacklistRepository.InsertBlacklist(blacklist)
	if err != nil {
		return err
	}
	return nil
}

func (rU roomBlacklistUsecase) DeleteBlacklist(blacklist *dbmodel.RoomBlacklist) error {
	err := rU.roomBlacklistRepository.DeleteBlacklist(blacklist)
	if err != nil {
		return err
	}
	return nil
}
