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

func (d *gameHardDatastore) FindAll() ([]*models.GameHard, error) {
	hards := []*models.GameHard{}
	err := d.db.Find(&hards).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return hards, nil
}

func (d *gameHardDatastore) Insert(hard *models.GameHard) error {
	err := d.db.Create(hard).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (d *gameHardDatastore) Update(hard *models.GameHard) error {
	err := d.db.Model(hard).
		Where("game_hard_id = ?", hard.GameHardID).
		Updates(models.GameHard{HardName: hard.HardName, UpdateAt: hard.UpdateAt}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (d *gameHardDatastore) Delete(hard *models.GameHard) error {
	err := d.db.Where("game_hard_id = ?", hard.GameHardID).Delete(hard).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}
