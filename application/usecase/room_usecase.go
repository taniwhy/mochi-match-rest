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
	GetList(c *gin.Context) ([]*output.RoomResBody, *int, error)
	GetByID(c *gin.Context) (*output.RoomDetailResBody, error)
	GetHotGame(c *gin.Context) ([]*output.HotGameRes, error)
	Create(c *gin.Context) (*output.RoomDetailResBody, error)
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	Join(c *gin.Context) error
	Leave(c *gin.Context) error
	CheckEntry(c *gin.Context) (bool, *output.RoomDetailResBody, error)
}

type roomUsecase struct {
	roomRepository         repository.IRoomRepository
	entryHistoryRepository repository.IEntryHistoryRepository
	roomService            service.IRoomService
	entryHistoryService    service.IEntryHistoryService
}

// NewRoomUsecase : Roomユースケースの生成
func NewRoomUsecase(
	rR repository.IRoomRepository,
	eHR repository.IEntryHistoryRepository,
	rS service.IRoomService,
	hS service.IEntryHistoryService) IRoomUseCase {
	return &roomUsecase{
		roomRepository:         rR,
		entryHistoryRepository: eHR,
		roomService:            rS,
		entryHistoryService:    hS,
	}
}

func (u *roomUsecase) GetList(c *gin.Context) ([]*output.RoomResBody, *int, error) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, nil, errors.ErrParams{Need: "page=`int`", Got: pageStr}
	}
	if page < 1 {
		return nil, nil, errors.ErrParams{Need: "page=`int`", Got: pageStr}
	}
	limit := 12
	offset := 12 * (page - 1)
	if page == 1 {
		offset = 0
	}
	rooms, err := u.roomRepository.FindByLimitAndOffset(limit, offset)
	roomCnt, err := u.roomRepository.FindUnlockCountByID()
	return rooms, roomCnt, nil
}

func (u *roomUsecase) GetByID(c *gin.Context) (*output.RoomDetailResBody, error) {
	roomID := c.Params.ByName("id")
	room, err := u.roomRepository.FindByID(roomID)
	if err != nil {
		return nil, err
	}
	joinUsers, err := u.entryHistoryRepository.FindNotLeaveListByRoomID(roomID)

	resBody := &output.RoomDetailResBody{
		RoomID:    roomID,
		OwnerID:   room.UserID,
		HardName:  room.HardName,
		GameTitle: room.GameTitle,
		Capacity:  room.Capacity,
		Count:     room.Count,
		RoomText:  room.RoomText,
	}
	for _, g := range joinUsers {
		r := output.JoinUserRes{
			UserID:   g.UserID,
			UserName: g.UserName,
			Icon:     g.Icon,
		}
		resBody.JoinUsers = append(resBody.JoinUsers, r)
	}
	return resBody, nil
}

func (u *roomUsecase) GetHotGame(c *gin.Context) ([]*output.HotGameRes, error) {
	return nil, nil
}

func (u *roomUsecase) Create(c *gin.Context) (*output.RoomDetailResBody, error) {
	body := input.RoomCreateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return nil, errors.ErrRoomCreateReqBinding{
			RoomText:   body.RoomText,
			GameListID: body.GameListID,
			GameHardID: body.GameHardID,
			Capacity:   body.Capacity,
			Start:      body.Start.Time,
		}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return nil, errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err := u.roomService.CanInsert(userID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.ErrRoomAlreadyExists{}
	}
	ok, err = u.entryHistoryService.CanJoin(userID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.ErrRoomAlreadyEntry{}
	}
	newRoom, err := models.NewRoom(userID, body.RoomText, body.GameListID, body.GameHardID, body.Capacity, body.Start.Time)
	if err != nil {
		return nil, err
	}
	if err := u.roomRepository.Insert(newRoom); err != nil {
		return nil, err
	}
	history, err := models.NewEntryHistory(userID, newRoom.RoomID)
	if err := u.entryHistoryRepository.Insert(history); err != nil {
		return nil, err
	}

	room, err := u.roomRepository.FindByID(newRoom.RoomID)
	if err != nil {
		return nil, err
	}
	joinUsers, err := u.entryHistoryRepository.FindNotLeaveListByRoomID(newRoom.RoomID)

	resBody := &output.RoomDetailResBody{
		RoomID:    newRoom.RoomID,
		OwnerID:   room.UserID,
		HardName:  room.HardName,
		GameTitle: room.GameTitle,
		Capacity:  room.Capacity,
		Count:     room.Count,
		RoomText:  room.RoomText,
	}
	for _, g := range joinUsers {
		r := output.JoinUserRes{
			UserID:   g.UserID,
			UserName: g.UserName,
			Icon:     g.Icon,
		}
		resBody.JoinUsers = append(resBody.JoinUsers, r)
	}
	return resBody, nil
}

func (u *roomUsecase) Update(c *gin.Context) error {
	return nil
}

func (u *roomUsecase) Delete(c *gin.Context) error {
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

func (u *roomUsecase) Join(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	if roomID == "" {
		return errors.ErrParams{Need: "id", Got: roomID}
	}
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
	ok, err = u.entryHistoryService.CanJoin(userID)
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

func (u *roomUsecase) Leave(c *gin.Context) error {
	roomID := c.Params.ByName("id")
	if roomID == "" {
		return errors.ErrParams{Need: "id", Got: roomID}
	}
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	ok, err := u.entryHistoryService.CheckJoin(roomID, userID)
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

func (u *roomUsecase) CheckEntry(c *gin.Context) (bool, *output.RoomDetailResBody, error) {
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return false, nil, errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	history, err := u.entryHistoryRepository.FindNotLeave(userID)
	// 入室していない
	if history == nil {
		return false, nil, nil
	}
	// 入室している
	roomID := history.RoomID
	room, err := u.roomRepository.FindByID(roomID)
	if err != nil {
		return false, nil, err
	}
	joinUsers, err := u.entryHistoryRepository.FindNotLeaveListByRoomID(roomID)

	resBody := &output.RoomDetailResBody{
		RoomID:    roomID,
		OwnerID:   room.UserID,
		HardName:  room.HardName,
		GameTitle: room.GameTitle,
		Capacity:  room.Capacity,
		Count:     room.Count,
		RoomText:  room.RoomText,
	}
	for _, g := range joinUsers {
		r := output.JoinUserRes{
			UserID:   g.UserID,
			UserName: g.UserName,
			Icon:     g.Icon,
		}
		resBody.JoinUsers = append(resBody.JoinUsers, r)
	}
	return true, resBody, nil
}
