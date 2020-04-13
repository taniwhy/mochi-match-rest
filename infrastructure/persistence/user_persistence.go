package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type userPersistence struct {
	db *gorm.DB
}

// NewUserPersistence : UserPersistenseを生成.
func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

func (uP userPersistence) GetAllUser() ([]*models.User, error) {
	users := []*models.User{}

	err := uP.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uP userPersistence) GetUserByID(ID int) (*models.User, error) {
	User := models.User{ID: ID}
	err := uP.db.Take(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (uP userPersistence) GetUserByProviderID(provider, providerID string) (*models.User, error) {
	User := models.User{
		Provider:   provider,
		ProviderID: providerID,
	}
	err := uP.db.Take(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (uP userPersistence) InsertUser(User *models.User) error {
	return uP.db.Create(User).Error
}

func (uP userPersistence) UpdateUser(User *models.User) error {
	return uP.db.Updates(User).Error
}

func (uP userPersistence) DeleteUser(User *models.User) error {
	err := uP.db.Take(&User).Error
	if err != nil {
		return err
	}
	return uP.db.Delete(User).Error
}
