package datastore

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userDetailDatastore struct {
	db *gorm.DB
}

// NewUserDetailDatastore : ユーザー詳細データストアの生成
func NewUserDetailDatastore(db *gorm.DB) repository.IUserDetailRepository {
	return &userDetailDatastore{db}
}

func (uD userDetailDatastore) FindByID(id string) (*models.UserDetail, error) {
	userDetails := models.UserDetail{}
	err := uD.db.Where("user_id = ?", id).First(&userDetails).Error
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func (uD userDetailDatastore) Insert(userDetail *models.UserDetail) error {
	return uD.db.Create(userDetail).Error
}

func (uD userDetailDatastore) Update(id, name, icon string) error {
	u := models.UserDetail{}
	fmt.Println(id, name, icon)
	err := uD.db.Model(&u).
		Where("user_id = ?", id).
		Update("user_name", name).
		Update("icon", icon).
		Update("update_at", time.Now()).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (uD userDetailDatastore) Delete(userDetail *models.UserDetail) error {
	err := uD.db.Take(&userDetail).Error
	if err != nil {
		return err
	}
	return uD.db.Delete(userDetail).Error
}
