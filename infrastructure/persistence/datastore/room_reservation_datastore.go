package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type roomReservationDatastore struct {
	db *gorm.DB
}

// NewRoomReservationDatastore : UserPersistenseを生成.
func NewRoomReservationDatastore(db *gorm.DB) repository.RoomReservationRepository {
	return &roomReservationDatastore{db}
}

func (rD roomReservationDatastore) FindAllRoomReservation() ([]*dbmodel.RoomReservation, error) {
	blacklist := []*dbmodel.RoomReservation{}

	err := rD.db.Find(&blacklist).Error
	if err != nil {
		return nil, err
	}
	return blacklist, nil
}

func (rD roomReservationDatastore) FindRoomReservationByID(id int64) (*dbmodel.RoomReservation, error) {
	roomReservation := dbmodel.RoomReservation{ID: id}
	err := rD.db.Take(&roomReservation).Error
	if err != nil {
		return nil, err
	}
	return &roomReservation, nil
}

func (rD roomReservationDatastore) InsertRoomReservation(roomReservation *dbmodel.RoomReservation) error {
	return rD.db.Create(roomReservation).Error
}

func (rD roomReservationDatastore) UpdateRoomReservation(roomReservation *dbmodel.RoomReservation) error {
	return rD.db.Update(roomReservation).Error
}

func (rD roomReservationDatastore) DeleteRoomReservation(roomReservation *dbmodel.RoomReservation) error {
	err := rD.db.Take(&roomReservation).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(roomReservation).Error
}
