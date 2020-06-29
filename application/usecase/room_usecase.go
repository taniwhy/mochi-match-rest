//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IRoomUseCase : インターフェース
type IRoomUseCase interface {
	GetList(c *gin.Context) ([]*output.RoomResBody, error)
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	Join(c *gin.Context) error
	Leave(c *gin.Context) error
}

type roomUsecase struct {
	roomRepository         repository.IRoomRepository
	entryHistoryRepository repository.IEntryHistoryRepository
	roomService            service.IRoomService
}

// NewRoomUsecase : Roomユースケースの生成
func NewRoomUsecase(
	rR repository.IRoomRepository,
	eHR repository.IEntryHistoryRepository,
	rS service.IRoomService) IRoomUseCase {
	return &roomUsecase{
		roomRepository:         rR,
		entryHistoryRepository: eHR,
		roomService:            rS,
	}
}

func (u roomUsecase) GetList(c *gin.Context) ([]*output.RoomResBody, error) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.ErrParams{Need: "page=`int`", Got: pageStr}
	}
	if page < 1 {
		return nil, errors.ErrParams{Need: "page=`int`", Got: pageStr}
	}
	limit := 8
	offset := 8 * (page - 1)
	if page == 1 {
		offset = 0
	}
	rooms, err := u.roomRepository.FindByLimitAndOffset(limit, offset)
	return rooms, nil
}

func (u roomUsecase) Create(c *gin.Context) error {
	body := input.RoomCreateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrRoomCreateReqBinding{
			RoomText:   body.RoomText,
			GameListID: body.GameListID,
			GameHardID: body.GameHardID,
			Capacity:   body.Capacity,
			Start:      body.Start.Time,
		}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err := u.roomService.CanInsert(userID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyExists{}
	}
	room, err := models.NewRoom(userID, body.RoomText, body.GameListID, body.GameHardID, body.Capacity, body.Start.Time)
	if err != nil {
		return err
	}
	if err := u.roomRepository.Insert(room); err != nil {
		return err
	}
	history, err := models.NewEntryHistory(userID, room.RoomID)
	if err := u.entryHistoryRepository.Insert(history); err != nil {
		return err
	}
	return nil
}

func (u roomUsecase) Update(c *gin.Context) error {
	return nil
}

func (u roomUsecase) Delete(c *gin.Context) error {
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
	if err := u.roomRepository.LockFlg(userID, roomID); err != nil {
		return err
	}
	return nil
}

func (u roomUsecase) Join(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	ok, err := u.roomService.IsLock(roomID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyLock{RoomID: roomID}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err = u.entryHistoryRepository.CheckEntry(roomID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.ErrRoomAlreadyEntry{RoomID: roomID}
	}
	room, err := u.roomRepository.FindByID(roomID)
	roomCap := room.Capacity
	userCnt, err := u.entryHistoryRepository.CountEntryUser(roomID)
	if roomCap <= userCnt {
		return errors.ErrRoomCapacityOver{RoomID: roomID, Count: userCnt}
	}
	history, err := models.NewEntryHistory(userID, roomID)
	if err := u.entryHistoryRepository.Insert(history); err != nil {
		return err
	}
	return nil
}

func (u roomUsecase) Leave(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err := u.entryHistoryRepository.CheckEntry(roomID, userID)
	if err != nil {
		return err
	}
	if ok {
		return errors.ErrNotEntryRoom{RoomID: roomID}
	}
	if err := u.entryHistoryRepository.LeaveFlg(roomID, userID); err != nil {
		return err
	}
	return nil
}
