package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// IRoomBlacklistHandler : インターフェース
type IRoomBlacklistHandler interface {
	GetByRoomID(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
}

type roomBlacklistHandler struct {
	roomBlacklistUsecase usecase.IRoomBlacklistUseCase
}

// NewRoomBlacklistHandler : ルームブラックリストハンドラの生成
func NewRoomBlacklistHandler(rBU usecase.IRoomBlacklistUseCase) IRoomBlacklistHandler {
	return &roomBlacklistHandler{
		roomBlacklistUsecase: rBU,
	}
}

func (rH roomBlacklistHandler) GetByRoomID(c *gin.Context) {
	rB, err := rH.roomBlacklistUsecase.GetByRoomID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, rB)
}

func (rH roomBlacklistHandler) Create(c *gin.Context) {
	err := rH.roomBlacklistUsecase.Create(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create blacklist"})
}

func (rH roomBlacklistHandler) Delete(c *gin.Context) {
	err := rH.roomBlacklistUsecase.Delete(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete blacklist"})
}
