package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userDatastore struct {
	db *gorm.DB
}

// NewUserDatastore : UserPersistenseを生成.
func NewUserDatastore(db *gorm.DB) repository.UserRepository {
	return &userDatastore{db}
}

func (uD userDatastore) FindByID(id string) (*models.User, error) {
	u := models.User{}
	err := uD.db.Where("user_id = ?", id).First(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return &u, nil
}

func (uD userDatastore) FindByProviderID(provider, id string) (*models.User, error) {
	u := models.User{}
	switch provider {
	case "google":
		err := uD.db.Where("google_id = ?", id).First(&u).Error
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		if err != nil {
			return nil, errors.ErrDataBase{Detail: err}
		}
		return nil, errors.ErrIDAlreadyExists{Provider: provider, ID: id}
	default:
		return nil, errors.ErrUnexpectedQueryProvider{Provider: provider}
	}
}

func (uD userDatastore) Insert(user *models.User) error {
	return uD.db.Create(user).Error
}

func (uD userDatastore) Update(user *models.User) error {
	return uD.db.Updates(user).Error
}

func (uD userDatastore) Delete(id string) error {
	user := models.User{}
	recordNotFound := uD.db.Where("user_id = ?", id).Take(&user).RecordNotFound()
	if recordNotFound {
		return fmt.Errorf("Record not found : %v", id)
	}
	return uD.db.Model(&user).Where("user_id = ?", id).Update("is_delete", true).Error
}
