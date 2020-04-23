package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// UserDetailRepository : ユーザーのリポジトリ
type UserDetailRepository interface {
	FindAllUserDetail() ([]*models.UserDetail, error)
	FindUserDetailByID(id int64) (*models.UserDetail, error)
	InsertUserDetail(userDetail *models.UserDetail) error
	UpdateUserDetail(userDetail *models.UserDetail) error
	DeleteUserDetail(userDetail *models.UserDetail) error
}
