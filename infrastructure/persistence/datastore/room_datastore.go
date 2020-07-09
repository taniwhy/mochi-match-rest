package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type roomDatastore struct {
	db *gorm.DB
}

// NewRoomDatastore : ルームデータストアの生成
func NewRoomDatastore(db *gorm.DB) repository.IRoomRepository {
	return &roomDatastore{db}
}

func (d *roomDatastore) FindList() ([]*output.RoomResBody, error) {
	rooms := []*output.RoomResBody{}
	err := d.db.
		Table("rooms").
		Select(`rooms.room_id,
			rooms.user_id,
			user_details.icon,
			game_hards.hard_name,
			game_lists.game_title,
			rooms.capacity,
			rooms.room_text,
			user_details.user_name,
			(
				SELECT
					COUNT(entry_histories.entry_history_id)
				FROM entry_histories
				WHERE rooms.room_id = entry_histories.room_id
			) As count,
			rooms.created_at,
			rooms.start`).
		Joins("LEFT JOIN user_details ON rooms.user_id = user_details.user_id").
		Joins("LEFT JOIN game_hards ON rooms.game_hard_id = game_hards.game_hard_id").
		Joins("LEFT JOIN game_lists ON rooms.game_list_id = game_lists.game_list_id").
		Order("created_at desc").
		Scan(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *roomDatastore) FindByLimitAndOffset(limit, offset int) ([]*output.RoomResBody, error) {
	rooms := []*output.RoomResBody{}
	err := d.db.
		Table("rooms").
		Select(`rooms.room_id,
				rooms.user_id,
				user_details.icon,
				game_hards.hard_name,
				game_lists.game_title,
				rooms.capacity,
				rooms.room_text,
				user_details.user_name,
				(
					SELECT
						COUNT(entry_histories.entry_history_id)
					FROM entry_histories
					WHERE rooms.room_id = entry_histories.room_id AND entry_histories.is_leave = false
				) As count,
				rooms.is_lock,
				rooms.created_at,
				rooms.start`).
		Joins("LEFT JOIN user_details ON rooms.user_id = user_details.user_id").
		Joins("LEFT JOIN game_hards ON rooms.game_hard_id = game_hards.game_hard_id").
		Joins("LEFT JOIN game_lists ON rooms.game_list_id = game_lists.game_list_id").
		Where("rooms.is_lock = ?", false).
		Limit(limit).Offset(offset).Order("created_at desc").Scan(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *roomDatastore) FindByID(id string) (*output.RoomResBody, error) {
	room := &output.RoomResBody{}
	err := d.db.
		Table("rooms").
		Select(`rooms.room_id,
				rooms.user_id,
				user_details.icon,
				game_hards.hard_name,
				game_lists.game_title,
				rooms.capacity,
				rooms.room_text,
				user_details.user_name,
				(
					SELECT
						COUNT(entry_histories.entry_history_id)
					FROM entry_histories
					WHERE rooms.room_id = entry_histories.room_id AND entry_histories.is_leave = false
				) As count,
				rooms.is_lock,
				rooms.created_at,
				rooms.start`).
		Joins("LEFT JOIN user_details ON rooms.user_id = user_details.user_id").
		Joins("LEFT JOIN game_hards ON rooms.game_hard_id = game_hards.game_hard_id").
		Joins("LEFT JOIN game_lists ON rooms.game_list_id = game_lists.game_list_id").
		Where("rooms.is_lock = ?", false).Order("created_at desc").Scan(&room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (d *roomDatastore) FindByUserID(userID string) ([]*models.Room, error) {
	rooms := []*models.Room{}
	err := d.db.Where("user_id = ?", userID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *roomDatastore) FindUnlockByID(userID string) (*models.Room, error) {
	rooms := &models.Room{}
	err := d.db.Where("user_id = ? AND is_lock = ?", userID, false).First(&rooms).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return rooms, nil
}

func (d *roomDatastore) Insert(room *models.Room) error {
	err := d.db.Create(room).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (d *roomDatastore) Update(room *models.Room) error {
	return d.db.Updates(room).Error
}

func (d *roomDatastore) Delete(room *models.Room) error {
	err := d.db.Take(&room).Error
	if err != nil {
		return err
	}
	return d.db.Delete(room).Error
}

// todo ロック時間は保存する？
func (d *roomDatastore) LockFlg(uid, rid string) error {
	h := &models.Room{}
	err := d.db.Model(&h).
		Where("room_id = ? AND user_id = ? AND is_lock = ?", rid, uid, false).
		Updates(models.Room{IsLock: true}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrRecordNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err}
	}
	return nil
}
