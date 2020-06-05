package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
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

func (rD roomDatastore) FindList() ([]*models.Room, error) {
	rooms := []*models.Room{}

	err := rD.db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rD roomDatastore) FindByID(id string) ([]*models.Room, error) {
	rooms := []*models.Room{}
	err := rD.db.Where("user_id = ?", id).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rD roomDatastore) FindUnlockByID(id string) (*models.Room, error) {
	rooms := &models.Room{}
	err := rD.db.Where("user_id = ? AND is_lock = ?", id, false).First(&rooms).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return rooms, nil
}

func (rD roomDatastore) Insert(room *models.Room) error {
	err := rD.db.Create(room).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (rD roomDatastore) Update(room *models.Room) error {
	return rD.db.Updates(room).Error
}

func (rD roomDatastore) Delete(room *models.Room) error {
	err := rD.db.Take(&room).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(room).Error
}
