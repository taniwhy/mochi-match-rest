package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// IRoomBlacklistHandler : インターフェース
type IRoomBlacklistHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
}

type roomBlacklistHandler struct {
	userUsecase          usecase.IUserUseCase
	roomUsecase          usecase.IRoomUseCase
	roomBlacklistUsecase usecase.IRoomBlacklistUseCase
}

// NewRoomBlacklistHandler : ルームブラックリストハンドラの生成
func NewRoomBlacklistHandler(
	uU usecase.IUserUseCase,
	rU usecase.IRoomUseCase,
	rBU usecase.IRoomBlacklistUseCase) IRoomBlacklistHandler {
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
