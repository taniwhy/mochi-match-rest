package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type gameListDatastore struct {
	db *gorm.DB
}

// NewGameListDatastore : ゲームタイトルデータストアの生成
func NewGameListDatastore(db *gorm.DB) repository.IGameListRepository {
	return &gameListDatastore{db}
}

func (gD gameListDatastore) FindAll() ([]*models.GameList, error) {
	gameTitle := []*models.GameList{}
	err := gD.db.Find(&gameTitle).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return gameTitle, nil
}

func (gD gameListDatastore) Insert(gT *models.GameList) error {
	err := gD.db.Create(gT).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (gD gameListDatastore) Update(gT *models.GameList) error {
	err := gD.db.Model(gT).
		Where("game_list_id = ?", gT.GameListID).
		Updates(models.GameList{GameTitle: gT.GameTitle, UpdateAt: gT.UpdateAt}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (gD gameListDatastore) Delete(gT *models.GameList) error {
	err := gD.db.Where("game_list_id = ?", gT.GameListID).Delete(gT).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}
