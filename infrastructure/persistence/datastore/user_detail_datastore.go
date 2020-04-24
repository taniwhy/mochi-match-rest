package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userDetailDatastore struct {
	db *gorm.DB
}

// NewUserDetailDatastore : UserPersistenseを生成.
func NewUserDetailDatastore(db *gorm.DB) repository.UserDetailRepository {
	return &userDetailDatastore{db}
}

func (uD userDetailDatastore) FindAllUserDetail() ([]*models.UserDetail, error) {
	userDetails := []*models.UserDetail{}

	err := uD.db.Find(&userDetails).Error
	if err != nil {
		return nil, err
	}
	return userDetails, nil
}

func (uD userDetailDatastore) FindUserDetailByID(id int64) (*models.UserDetail, error) {
	userDetails := models.UserDetail{ID: id}
	err := uD.db.Take(&userDetails).Error
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func (uD userDetailDatastore) InsertUserDetail(userDetail *models.UserDetail) error {
	return uD.db.Create(userDetail).Error
}

func (uD userDetailDatastore) UpdateUserDetail(userDetail *models.UserDetail) error {
	return uD.db.Updates(userDetail).Error
}

func (uD userDetailDatastore) DeleteUserDetail(userDetail *models.UserDetail) error {
	err := uD.db.Take(&userDetail).Error
	if err != nil {
		return err
	}
	return uD.db.Delete(userDetail).Error
}
