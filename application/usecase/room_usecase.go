package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
)

// RoomUseCase :
type RoomUseCase interface {
	GetList(*gin.Context) ([]*models.Room, error)
	GetByID(*gin.Context) (*models.Room, error)
	Create(*gin.Context) error
	Update(*gin.Context) error
}

type roomUsecase struct {
	roomRepository repository.RoomRepository
	roomService    service.IRoomService
}

// NewRoomUsecase :
func NewRoomUsecase(rR repository.RoomRepository, rS service.IRoomService) RoomUseCase {
	return &roomUsecase{
		roomRepository: rR,
		roomService:    rS,
	}
}

func (rU roomUsecase) GetList(c *gin.Context) ([]*models.Room, error) {
	return nil, nil
}

func (rU roomUsecase) GetByID(c *gin.Context) (*models.Room, error) {
	return nil, nil
}

func (rU roomUsecase) Create(c *gin.Context) error {
	b := input.RoomCreateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrRoomCreateReqBinding{
			RoomText:    b.RoomText,
			GameTitleID: b.GameTitleID,
			GameHardID:  b.GameHardID,
			Capacity:    b.Capacity,
		}
	}
	claims, err := auth.GetTokenClaims(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	claimsID := claims["sub"].(string)
	//  既に開いてるルームが無いか確認します
	ok, err := rU.roomService.CanInsert(claimsID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyExists{}
	}

	r, err := models.NewRoom(claimsID, b.RoomText, b.GameTitleID, b.GameHardID, b.Capacity)
	if err != nil {
		return err
	}
	if err := rU.roomRepository.Insert(r); err != nil {
		return err
	}
	return nil
}

func (rU roomUsecase) Update(c *gin.Context) error {
	return nil
}
