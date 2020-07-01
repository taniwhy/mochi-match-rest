//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IRoomBlacklistUseCase : インターフェース
type IRoomBlacklistUseCase interface {
	GetByRoomID(c *gin.Context) ([]*models.RoomBlacklist, error)
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
}

type roomBlacklistUsecase struct {
	roomBlacklistRepository repository.IRoomBlacklistRepository
	roomService             service.IRoomService
}

// NewRoomBlacklistUsecase : RoomBlacklistユースケースの生成
func NewRoomBlacklistUsecase(rR repository.IRoomBlacklistRepository, rS service.IRoomService) IRoomBlacklistUseCase {
	return &roomBlacklistUsecase{
		roomBlacklistRepository: rR,
		roomService:             rS,
	}
}

func (u *roomBlacklistUsecase) GetByRoomID(c *gin.Context) ([]*models.RoomBlacklist, error) {
	roomID := c.Params.ByName("id")
	blacklist, err := u.roomBlacklistRepository.FindByRoomID(roomID)
	if err != nil {
		return nil, err
	}
	return blacklist, nil
}

func (u *roomBlacklistUsecase) Create(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err := u.roomService.IsOwner(userID, roomID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrNotRoomOwner{RoomID: userID}
	}
	blacklist, err := models.NewBlacklist(roomID, userID)
	if err != nil {
		return err
	}
	if err := u.roomBlacklistRepository.Insert(blacklist); err != nil {
		return err
	}
	return nil
}

func (u *roomBlacklistUsecase) Delete(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err := u.roomService.IsOwner(userID, roomID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrNotRoomOwner{RoomID: userID}
	}
	err = u.roomBlacklistRepository.Delete(roomID, userID)
	if err != nil {
		return err
	}
	return nil
}
