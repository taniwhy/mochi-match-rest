package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

//UserRepository : ユーザーのリポジトリ
type UserRepository interface {
	FindAllUser() ([]*models.User, error)
	FindByID(userID string) (*models.User, error)
	StoreUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}
