package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type roomDatastore struct {
	db *gorm.DB
}

// NewRoomDatastore : UserPersistenseを生成.
func NewRoomDatastore(db *gorm.DB) repository.RoomRepository {
	return &roomDatastore{db}
}

func (rD roomDatastore) FindAllRoom() ([]*dbmodel.Room, error) {
	rooms := []*dbmodel.Room{}

	err := rD.db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rD roomDatastore) FindRoomByID(id string) (*dbmodel.Room, error) {
	room := dbmodel.Room{RoomID: id}
	err := rD.db.Take(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (rD roomDatastore) InsertRoom(room *dbmodel.Room) error {
	return rD.db.Create(room).Error
}

func (rD roomDatastore) UpdateRoom(room *dbmodel.Room) error {
	return rD.db.Updates(room).Error
}

func (rD roomDatastore) DeleteRoom(room *dbmodel.Room) error {
	err := rD.db.Take(&room).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(room).Error
}
