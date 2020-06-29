//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IRoomBlacklistUseCase : インターフェース
type IRoomBlacklistUseCase interface {
	GetByRoomID(c *gin.Context) ([]*models.RoomBlacklist, error)
	Insert(c *gin.Context) error
	Delete(c *gin.Context) error
}

type roomBlacklistUsecase struct {
	roomBlacklistRepository repository.IRoomBlacklistRepository
}

// NewRoomBlacklistUsecase : RoomBlacklistユースケースの生成
func NewRoomBlacklistUsecase(rR repository.IRoomBlacklistRepository) IRoomBlacklistUseCase {
	return &roomBlacklistUsecase{
		roomBlacklistRepository: rR,
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

func (u *roomBlacklistUsecase) Insert(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
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
	err := u.roomBlacklistRepository.Delete(roomID)
	if err != nil {
		return err
	}
	return nil
}
