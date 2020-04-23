package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// RoomHandler : インターフェース
type RoomHandler interface {
	GetAllRoom(*gin.Context)
	CreateRoom(*gin.Context)
}

type roomHandler struct {
	roomUsecase usecase.RoomUseCase
}

// NewRoomHandler : ユーザーのHandler生成
func NewRoomHandler(uU usecase.UserUseCase, uDU usecase.UserDetailUseCase) UserHandler {
	return &userHandler{
		userUsecase:       uU,
		userDetailUsecase: uDU,
	}
}

func (uh userHandler) GetAllRoom(c *gin.Context) {

}

func (uh userHandler) CreateRoom(c *gin.Context) {

}
