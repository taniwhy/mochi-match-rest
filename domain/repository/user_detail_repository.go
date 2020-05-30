package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// UserDetailRepository : ユーザーのリポジトリ
type UserDetailRepository interface {
	FindAllUserDetail() ([]*dbmodel.UserDetail, error)
	FindUserDetailByID(id string) (*dbmodel.UserDetail, error)
	InsertUserDetail(userDetail *dbmodel.UserDetail) error
	UpdateUserDetail(userDetail *dbmodel.UserDetail) error
	DeleteUserDetail(userDetail *dbmodel.UserDetail) error
}
