package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// RoomBlacklistHandler : インターフェース
type RoomBlacklistHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
}

type roomBlacklistHandler struct {
	userUsecase          usecase.UserUseCase
	roomUsecase          usecase.RoomUseCase
	roomBlacklistUsecase usecase.RoomBlacklistUseCase
}

// NewRoomBlacklistHandler : ユーザーのHandler生成
func NewRoomBlacklistHandler(
	uU usecase.UserUseCase,
	rU usecase.RoomUseCase,
	rBU usecase.RoomBlacklistUseCase) RoomBlacklistHandler {
	return &roomBlacklistHandler{
		userUsecase:          uU,
		roomUsecase:          rU,
		roomBlacklistUsecase: rBU,
	}
}

func (rH roomBlacklistHandler) GetByID(c *gin.Context) {

}

func (rH roomBlacklistHandler) Create(c *gin.Context) {

}

func (rH roomBlacklistHandler) Delete(c *gin.Context) {

}
