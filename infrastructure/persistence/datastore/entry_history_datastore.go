package datastore

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type entryHistoryDatastore struct {
	db *gorm.DB
}

// NewEntryHistoryDatastore : 入室履歴データストアの生成
func NewEntryHistoryDatastore(db *gorm.DB) repository.IEntryHistoryRepository {
	return &entryHistoryDatastore{db}
}

func (d *entryHistoryDatastore) FindAll() ([]*models.EntryHistory, error) {
	historys := []*models.EntryHistory{}
	err := d.db.Find(&historys).Error
	if err != nil {
		return nil, err
	}
	return historys, nil
}

func (d *entryHistoryDatastore) FindNotLeave(userID string) (*models.EntryHistory, error) {
	history := &models.EntryHistory{}
	err := d.db.Where("user_id = ? AND is_leave = ?", userID, false).First(&history).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return history, nil
}

func (d *entryHistoryDatastore) FindNotLeaveByRoomID(userID, roomID string) (*models.EntryHistory, error) {
	history := &models.EntryHistory{}
	err := d.db.
		Table("entry_histories").
		Select(`
		entry_histories.game_title,
		entry_histories.user_id,
		entry_histories.room_id,
		entry_histories.is_leave,
		entry_histories.created_at,
		entry_histories.leaved_at
		`).
		Joins("LEFT JOIN rooms ON entry_histories.room_id = rooms.room_id").
		Where(`
			entry_histories.user_id = ? AND
			entry_histories.room_id = ? AND
			entry_histories.is_leave = ? AND
			rooms.is_lock = ?`,
			userID, roomID, false, false).
		First(&history).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return history, nil
}

func (d *entryHistoryDatastore) FindNotLeaveListByRoomID(roomID string) ([]*output.JoinUserRes, error) {
	users := []*output.JoinUserRes{}
	err := d.db.
		Table("entry_histories").
		Select(`
			entry_histories.user_id,
			user_details.user_name,
			user_details.icon
			`).
		Joins("LEFT JOIN user_details ON entry_histories.user_id = user_details.user_id").
		Where("entry_histories.room_id = ? AND entry_histories.is_leave = ?", roomID, false).Order("created_at asc").Scan(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return users, nil
}

func (d *entryHistoryDatastore) Insert(histroy *models.EntryHistory) error {
	return d.db.Create(histroy).Error
}

func (d *entryHistoryDatastore) Update(histroy *models.EntryHistory) error {
	return d.db.Update(histroy).Error
}

func (d *entryHistoryDatastore) Delete(histroy *models.EntryHistory) error {
	err := d.db.Take(&histroy).Error
	if err != nil {
		return err
	}
	return d.db.Delete(histroy).Error
}

func (d *entryHistoryDatastore) CountEntryUser(roomID string) (int, error) {
	var count int
	h := models.EntryHistory{}
	err := d.db.Model(&h).Where("room_id = ? AND is_leave = ?", roomID, false).Count(&count).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0, nil
	}
	if err != nil {
		return 0, errors.ErrDataBase{Detail: err}
	}
	return count, nil
}

func (d *entryHistoryDatastore) LeaveFlg(roomID, userID string) error {
	h := &models.EntryHistory{}
	err := d.db.Model(&h).
		Where("room_id = ? AND user_id = ? AND is_leave = ?", roomID, userID, false).
		Updates(models.EntryHistory{IsLeave: true, LeavedAt: sql.NullTime{Time: time.Now(), Valid: true}}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err}
	}
	return nil
}
