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

func (eD entryHistoryDatastore) FindAll() ([]*models.EntryHistory, error) {
	entryHistorys := []*models.EntryHistory{}
	err := eD.db.Find(&entryHistorys).Error
	if err != nil {
		return nil, err
	}
	return entryHistorys, nil
}

func (eD entryHistoryDatastore) FindNotLeave(userID string) (*models.EntryHistory, error) {
	history := &models.EntryHistory{}
	err := eD.db.Where("user_id = ? AND is_leave = ?", userID, false).First(&history).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return history, nil
}

func (eD entryHistoryDatastore) FindNotLeaveByRoomID(userID, roomID string) (*models.EntryHistory, error) {
	history := &models.EntryHistory{}
	err := eD.db.Where("user_id = ? AND room_id = ? AND is_leave = ?", userID, roomID, false).First(&history).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return history, nil
}

func (eD entryHistoryDatastore) FindNotLeaveListByRoomID(roomID string) ([]*output.JoinUserRes, error) {
	users := []*output.JoinUserRes{}
	err := eD.db.
		Table("entry_histories").
		Select(`
			entry_histories.user_id,
			user_details.user_name,
			user_details.icon
			`).
		Joins("LEFT JOIN user_details ON entry_histories.user_id = user_details.user_id").
		Where("entry_histories.is_leave = ?", false).Order("created_at asc").Scan(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err}
	}
	return users, nil
}

func (eD entryHistoryDatastore) Insert(entryHistory *models.EntryHistory) error {
	return eD.db.Create(entryHistory).Error
}

func (eD entryHistoryDatastore) Update(entryHistory *models.EntryHistory) error {
	return eD.db.Update(entryHistory).Error
}

func (eD entryHistoryDatastore) Delete(entryHistory *models.EntryHistory) error {
	err := eD.db.Take(&entryHistory).Error
	if err != nil {
		return err
	}
	return eD.db.Delete(entryHistory).Error
}

func (eD entryHistoryDatastore) CountEntryUser(rid string) (int, error) {
	var count int
	h := models.EntryHistory{}
	err := eD.db.Model(&h).Where("room_id = ? AND is_leave = ?", rid, false).Count(&count).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0, nil
	}
	if err != nil {
		return 0, errors.ErrDataBase{Detail: err}
	}
	return count, nil
}

func (eD entryHistoryDatastore) LeaveFlg(rid, uid string) error {
	h := &models.EntryHistory{}
	err := eD.db.Model(&h).
		Where("room_id = ? AND user_id = ? AND is_leave = ?", rid, uid, false).
		Updates(models.EntryHistory{IsLeave: true, LeavedAt: sql.NullTime{Time: time.Now(), Valid: true}}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err}
	}
	return nil
}
