//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IEntryHistoryUseCase : インターフェース
type IEntryHistoryUseCase interface {
	GetByID(c *gin.Context) ([]*output.EntryHistoryResBody, error)
}

type entryHistoryUsecase struct {
	roomRepository         repository.IRoomRepository
	entryHistoryRepository repository.IEntryHistoryRepository
}

// NewEntryHistoryUsecase : 参加履歴ユースケースの生成
func NewEntryHistoryUsecase(
	rR repository.IRoomRepository,
	eR repository.IEntryHistoryRepository) IEntryHistoryUseCase {
	return &entryHistoryUsecase{
		roomRepository:         rR,
		entryHistoryRepository: eR,
	}
}

func (u *entryHistoryUsecase) GetByID(c *gin.Context) ([]*output.EntryHistoryResBody, error) {
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return nil, errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	histories, err := u.entryHistoryRepository.FindListByUserID(userID)
	if err != nil {
		return nil, err
	}
	var resBody []*output.EntryHistoryResBody
	for _, h := range histories {
		roomDetail, err := u.roomRepository.FindByID(h.RoomID)
		if roomDetail == nil {
			continue
		}
		joinUsers, err := u.entryHistoryRepository.FindListByRoomID(h.RoomID)
		if err != nil {
			return nil, err
		}
		b := output.EntryHistoryResBody{
			PlaydedDate: h.CreatedAt,
			HostName:    roomDetail.UserName,
			GameName:    roomDetail.GameTitle,
		}
		for _, g := range joinUsers {
			r := output.JoinUserRes{
				UserID:   g.UserID,
				UserName: g.UserName,
				Icon:     g.Icon,
			}
			b.JoinUsers = append(b.JoinUsers, r)
		}
		resBody = append(resBody, &b)
	}
	return resBody, nil
}
