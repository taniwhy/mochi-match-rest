package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// UserRepository : ユーザーのリポジトリ
type UserRepository interface {
	FindAllUser() ([]*dbmodel.User, error)
	FindUserByID(id string) (*dbmodel.User, error)
	FindUserByProviderID(provider, id string) (*dbmodel.User, error)
	InsertUser(user *dbmodel.User) error
	UpdateUser(user *dbmodel.User) error
	DeleteUser(id string) error
}
