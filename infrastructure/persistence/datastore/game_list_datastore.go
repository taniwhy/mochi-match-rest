package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type gameListDatastore struct {
	db *gorm.DB
}

// NewGameListDatastore : ゲームタイトルデータストアの生成
func NewGameListDatastore(db *gorm.DB) repository.IGameListRepository {
	return &gameListDatastore{db}
}

func (d *gameListDatastore) FindAll() ([]*models.GameList, error) {
	gamelists := []*models.GameList{}
	err := d.db.Find(&gamelists).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	return gamelists, nil
}

func (d *gameListDatastore) FindHot() ([]*output.HotGameRes, error) {
	hotGames := []*output.HotGameRes{}
	err := d.db.
		Table("game_lists").
		Select(`
			game_lists.game_list_id,
			game_lists.game_title,
			(
				SELECT
					COUNT(rooms.room_id)
				FROM
					rooms
				WHERE
					rooms.is_lock = false AND
					rooms.game_list_id = game_lists.game_list_id
			) As room_count,
			(
				SELECT
					COUNT(entry_histories.entry_history_id)
				FROM
					entry_histories
				WHERE
					entry_histories.room_id = rooms.room_id AND
					entry_histories.is_leave = false
			) As player_count
			`).
		Joins(`LEFT JOIN rooms ON rooms.game_list_id = game_lists.game_list_id`).
		Joins(`LEFT JOIN entry_histories ON entry_histories.room_id = rooms.room_id`).
		Limit(10).
		Order("player_count desc").
		Scan(&hotGames).Error
	if err != nil {
		return nil, errors.ErrDataBase{Detail: err.Error()}
	}
	fmt.Println(hotGames[0])
	return hotGames, nil
}

func (d *gameListDatastore) Insert(gamelist *models.GameList) error {
	err := d.db.Create(gamelist).Error
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (d *gameListDatastore) Update(gamelist *models.GameList) error {
	err := d.db.Model(gamelist).
		Where("game_list_id = ?", gamelist.GameListID).
		Updates(models.GameList{GameTitle: gamelist.GameTitle, UpdateAt: gamelist.UpdateAt}).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}

func (d *gameListDatastore) Delete(gamelist *models.GameList) error {
	err := d.db.Where("game_list_id = ?", gamelist.GameListID).Delete(gamelist).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.ErrNotFound{Detail: err.Error()}
	}
	if err != nil {
		return errors.ErrDataBase{Detail: err.Error()}
	}
	return nil
}
