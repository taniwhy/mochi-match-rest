package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// UserRepository : ユーザーのリポジトリ
type UserRepository interface {
	FindAllUser() ([]*models.User, error)
	FindUserByID(id string) (*models.User, error)
	FindUserByProviderID(provider, id string) (*models.User, error)
	InsertUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}
