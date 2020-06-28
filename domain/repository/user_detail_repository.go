//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// UserDetailRepository : ユーザー詳細のリポジトリ
type UserDetailRepository interface {
	FindByID(id string) (*models.UserDetail, error)
	Insert(userDetail *models.UserDetail) error
	Update(id, name, icon string) error
	Delete(userDetail *models.UserDetail) error
}
