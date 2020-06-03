package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type roomDatastore struct {
	db *gorm.DB
}

// NewRoomDatastore : UserPersistenseを生成.
func NewRoomDatastore(db *gorm.DB) repository.RoomRepository {
	return &roomDatastore{db}
}

func (rD roomDatastore) FindAllRoom() ([]*models.Room, error) {
	rooms := []*models.Room{}

	err := rD.db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rD roomDatastore) FindRoomByID(id string) (*models.Room, error) {
	room := models.Room{RoomID: id}
	err := rD.db.Take(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (rD roomDatastore) InsertRoom(room *models.Room) error {
	return rD.db.Create(room).Error
}

func (rD roomDatastore) UpdateRoom(room *models.Room) error {
	return rD.db.Updates(room).Error
}

func (rD roomDatastore) DeleteRoom(room *models.Room) error {
	err := rD.db.Take(&room).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(room).Error
}
