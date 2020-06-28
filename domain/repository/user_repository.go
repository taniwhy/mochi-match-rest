//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// IUserRepository : ユーザーのリポジトリ
type IUserRepository interface {
	FindByID(id string) (*models.User, error)
	FindByProviderID(provider, id string) (*models.User, error)
	Insert(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
}
