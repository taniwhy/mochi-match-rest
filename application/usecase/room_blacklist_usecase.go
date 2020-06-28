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
	GetByRoomID(*gin.Context) ([]*models.RoomBlacklist, error)
	Insert(*gin.Context) error
	Delete(*gin.Context) error
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

func (rU roomBlacklistUsecase) GetByRoomID(c *gin.Context) ([]*models.RoomBlacklist, error) {
	rid := c.Params.ByName("id")
	blacklist, err := rU.roomBlacklistRepository.FindByRoomID(rid)
	if err != nil {
		return nil, err
	}
	return blacklist, nil
}

func (rU roomBlacklistUsecase) Insert(c *gin.Context) error {
	rid := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	uid := claims["sub"].(string)
	rB, err := models.NewBlacklist(rid, uid)
	if err != nil {
		return err
	}
	if err := rU.roomBlacklistRepository.Insert(rB); err != nil {
		return err
	}
	return nil
}

func (rU roomBlacklistUsecase) Delete(c *gin.Context) error {
	rid := c.Params.ByName("id")
	err := rU.roomBlacklistRepository.Delete(rid)
	if err != nil {
		return err
	}
	return nil
}
