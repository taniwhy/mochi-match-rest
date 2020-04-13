package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// UserRepository : ユーザーのリポジトリ
type UserRepository interface {
	GetAllUser() ([]*models.User, error)
	GetUserByID(ID int) (*models.User, error)
	GetUserByProviderID(provider, providerID string) (*models.User, error)
	InsertUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}
