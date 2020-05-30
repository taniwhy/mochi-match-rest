package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userDatastore struct {
	db *gorm.DB
}

// NewUserDatastore : UserPersistenseを生成.
func NewUserDatastore(db *gorm.DB) repository.UserRepository {
	return &userDatastore{db}
}

func (uD userDatastore) FindAllUser() ([]*dbmodel.User, error) {
	users := []*dbmodel.User{}

	err := uD.db.Find(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uD userDatastore) FindUserByID(id string) (*dbmodel.User, error) {
	User := dbmodel.User{UserID: id}
	err := uD.db.Take(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (uD userDatastore) FindUserByProviderID(provider, id string) (*dbmodel.User, error) {
	var err error
	user := dbmodel.User{}

	switch provider {
	case "google":
		err = uD.db.Where("google_id = ?", id).Take(&user).Error
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uD userDatastore) InsertUser(user *dbmodel.User) error {
	return uD.db.Create(user).Error
}

func (uD userDatastore) UpdateUser(user *dbmodel.User) error {
	return uD.db.Updates(user).Error
}

func (uD userDatastore) DeleteUser(id string) error {
	user := dbmodel.User{}
	recordNotFound := uD.db.Where("user_id = ?", id).Take(&user).RecordNotFound()
	if recordNotFound {
		return fmt.Errorf("Record not found : %v", id)
	}
	return uD.db.Model(&user).Where("user_id = ?", id).Update("is_delete", true).Error
}
