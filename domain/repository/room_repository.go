//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
)

// IRoomRepository : ルームのリポジトリ
type IRoomRepository interface {
	FindList() ([]*output.RoomResBody, error)
	FindByLimitAndOffset(limit, offset int) ([]*output.RoomResBody, error)
	FindByID(roomID string) (*output.RoomResBody, error)
	FindByUserID(userID string) ([]*models.Room, error)
	FindUnlockByID(roomID string) (*models.Room, error)
	FindUnlockCountByID() (*int, error)
	Insert(room *models.Room) error
	Update(room *models.Room) error
	Delete(room *models.Room) error
	LockFlg(string, string) error
}
