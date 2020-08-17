//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
)

// IEntryHistoryRepository : 入室履歴のリポジトリ
type IEntryHistoryRepository interface {
	FindAll() ([]*models.EntryHistory, error)
	FindNotLeave(userID string) (*models.EntryHistory, error)
	FindNotLeaveByRoomID(userID, roomID string) (*models.EntryHistory, error)
	FindListByRoomID(roomID string) ([]*output.JoinUserRes, error)
	FindListByUserID(userID string) ([]*models.EntryHistory, error)
	FindNotLeaveListByRoomID(roomID string) ([]*output.JoinUserRes, error)
	Insert(*models.EntryHistory) error
	Update(*models.EntryHistory) error
	Delete(*models.EntryHistory) error
	CountEntryUser(string) (int, error)
	LeaveFlg(rid, uid string) error
}
