package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type gameHardDatastore struct {
	db *gorm.DB
}

// NewGameHardDatastore : ゲームハードデータストアの生成
func NewGameHardDatastore(db *gorm.DB) repository.IGameHardRepository {
	return &gameHardDatastore{db}
}

func (gD gameHardDatastore) FindAll() ([]*models.GameHard, error) {
	gameHards := []*models.GameHard{}
	err := gD.db.Find(&gameHards).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return gameHards, nil
}

func (gD gameHardDatastore) Insert(gH *models.GameHard) error {
	err := gD.db.Create(gH).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (gD gameHardDatastore) Update(gH *models.GameHard) error {
	err := gD.db.Model(gH).
		Where("game_hard_id = ?", gH.GameHardID).
		Updates(models.GameHard{HardName: gH.HardName, UpdateAt: gH.UpdateAt}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (gD gameHardDatastore) Delete(gH *models.GameHard) error {
	err := gD.db.Where("game_hard_id = ?", gH.GameHardID).Delete(gH).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}
