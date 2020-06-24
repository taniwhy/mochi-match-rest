package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IRoomUseCase : インターフェース
type IRoomUseCase interface {
	GetList(*gin.Context) ([]*models.Room, error)
	GetByID(*gin.Context) (*models.Room, error)
	Create(*gin.Context) error
	Update(*gin.Context) error
	Delete(*gin.Context) error
	Join(*gin.Context) error
	Leave(*gin.Context) error
}

type roomUsecase struct {
	roomRepository         repository.RoomRepository
	entryHistoryRepository repository.EntryHistoryRepository
	roomService            service.IRoomService
}

// NewRoomUsecase : Roomユースケースの生成
func NewRoomUsecase(
	rR repository.RoomRepository,
	eHR repository.EntryHistoryRepository,
	rS service.IRoomService) IRoomUseCase {
	return &roomUsecase{
		roomRepository:         rR,
		entryHistoryRepository: eHR,
		roomService:            rS,
	}
}

func (rU roomUsecase) GetList(c *gin.Context) ([]*models.Room, error) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.ErrParams{Need: "page", Got: pageStr}
	}
	limit := 8
	offset := 8 * (page - 1)
	if page == 1 {
		offset = 0
	}
	r, err := rU.roomRepository.FindByLimitAndOffset(limit, offset)
	return r, nil
}

func (rU roomUsecase) GetByID(c *gin.Context) (*models.Room, error) {
	return nil, nil
}

func (rU roomUsecase) Create(c *gin.Context) error {
	b := input.RoomCreateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrRoomCreateReqBinding{
			RoomText:   b.RoomText,
			GameListID: b.GameListID,
			GameHardID: b.GameHardID,
			Capacity:   b.Capacity,
			Start:      b.Start.Time,
		}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	claimsID := claims["sub"].(string)

	ok, err := rU.roomService.CanInsert(claimsID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyExists{}
	}

	r, err := models.NewRoom(claimsID, b.RoomText, b.GameListID, b.GameHardID, b.Capacity, b.Start.Time)
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

func (rU roomUsecase) Delete(c *gin.Context) error {
	return nil
}

func (rU roomUsecase) Join(c *gin.Context) error {
	rid := c.Params.ByName("id")
	ok, err := rU.roomService.IsLock(rid)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyLock{RoomID: rid}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	uid := claims["sub"].(string)
	ok, err = rU.entryHistoryRepository.CheckEntry(rid, uid)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyEntry{RoomID: rid}
	}
	r, err := rU.roomRepository.FindByID(rid)
	cap := r.Capacity
	if r.UserID != uid {
		cap--
	}
	cnt, err := rU.entryHistoryRepository.CountEntryUser(rid)
	if cap < cnt {
		return errors.ErrRoomCapacityOver{RoomID: rid, Count: cnt}
	}
	h, err := models.NewEntryHistory(uid, rid)
	if err := rU.entryHistoryRepository.Insert(h); err != nil {
		return err
	}
	return nil
}

func (rU roomUsecase) Leave(c *gin.Context) error {
	rid := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	uid := claims["sub"].(string)
	ok, err := rU.entryHistoryRepository.CheckEntry(rid, uid)
	if err != nil {
		return err
	}
	if ok {
		return errors.ErrNotEntryRoom{RoomID: rid}
	}
	if err := rU.entryHistoryRepository.LeaveFlg(rid, uid); err != nil {
		return err
	}
	return nil
}
