package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userDetailDatastore struct {
	db *gorm.DB
}

// NewUserDetailDatastore : UserPersistenseを生成.
func NewUserDetailDatastore(db *gorm.DB) repository.UserDetailRepository {
	return &userDetailDatastore{db}
}
func (uD userDetailDatastore) FindAllUserDetail() ([]*dbmodel.UserDetail, error) {
	userDetails := []*dbmodel.UserDetail{}

	err := uD.db.Find(&userDetails).Error
	if err != nil {
		return nil, err
	}
	return userDetails, nil
}

func (uD userDetailDatastore) FindUserDetailByID(id string) (*dbmodel.UserDetail, error) {
	userDetails := dbmodel.UserDetail{UserDetailID: id}
	err := uD.db.Take(&userDetails).Error
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}

func (uD userDetailDatastore) InsertUserDetail(userDetail *dbmodel.UserDetail) error {
	return uD.db.Create(userDetail).Error
}

func (uD userDetailDatastore) UpdateUserDetail(userDetail *dbmodel.UserDetail) error {
	u := dbmodel.UserDetail{}
	// todo : 更新対象が見つからないとクラッシュ
	return uD.db.Model(&u).Where("user_id = ?", userDetail.UserID).Updates(dbmodel.UserDetail{
		UserName: userDetail.UserName,
		Icon:     userDetail.Icon,
		UpdateAt: userDetail.UpdateAt,
	}).Error
}

func (uD userDetailDatastore) DeleteUserDetail(userDetail *dbmodel.UserDetail) error {
	err := uD.db.Take(&userDetail).Error
	if err != nil {
		return err
	}
	return uD.db.Delete(userDetail).Error
}
