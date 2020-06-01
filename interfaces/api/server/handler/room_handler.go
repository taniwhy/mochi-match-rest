package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// RoomHandler : インターフェース
type RoomHandler interface {
	GetRoomList(*gin.Context)
	GetRoomByID(*gin.Context)
	CreateRoom(*gin.Context)
	GetBlacklist(*gin.Context)
	CreateBlacklist(*gin.Context)
}

type roomHandler struct {
	userUsecase            usecase.UserUseCase
	roomUsecase            usecase.RoomUseCase
	roomBlacklistUseCase   usecase.RoomBlacklistUseCase
	roomReservationUseCase usecase.RoomReservationUseCase
}

// NewRoomHandler : ユーザーのHandler生成
func NewRoomHandler(
	uU usecase.UserUseCase,
	rU usecase.RoomUseCase,
	rBU usecase.RoomBlacklistUseCase,
	rRU usecase.RoomReservationUseCase) RoomHandler {
	return &roomHandler{
		userUsecase:            uU,
		roomUsecase:            rU,
		roomBlacklistUseCase:   rBU,
		roomReservationUseCase: rRU,
	}
}

func (rH roomHandler) GetRoomList(c *gin.Context) {

}

func (rH roomHandler) GetRoomByID(c *gin.Context) {

}

func (rH roomHandler) CreateRoom(c *gin.Context) {

}

func (rH roomHandler) GetBlacklist(c *gin.Context) {

}

func (rH roomHandler) CreateBlacklist(c *gin.Context) {

}
