package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models/output"

	log "github.com/sirupsen/logrus"
)

// IChatPostHandler : インターフェース
type IChatPostHandler interface {
	GetChatPostByRoomID(*gin.Context)
	CreateChatPost(*gin.Context)
}

type chatPostHandler struct {
	chatPostUsecase usecase.IChatPostUseCase
}

// NewChatPostHandler : チャット投稿ハンドラの生成
func NewChatPostHandler(cU usecase.IChatPostUseCase) IChatPostHandler {
	return &chatPostHandler{
		chatPostUsecase: cU,
	}
}

var (
	messages []*output.ChatPostResBody
	err      error
)

func (h *chatPostHandler) GetChatPostByRoomID(c *gin.Context) {
	roomID := c.Params.ByName("id")
	limitStr := c.Query("limit")
	offset := c.Query("offset")
	switch {
	case limitStr == "" && offset == "":
		messages, err = h.chatPostUsecase.FindByRoomID(roomID)
	case limitStr != "" && offset == "":
		messages, err = h.chatPostUsecase.FindByRoomIDAndLimit(roomID, limitStr)
	case limitStr == "" && offset != "":
		messages, err = h.chatPostUsecase.FindByRoomIDAndOffset(roomID, offset)
	case limitStr != "" && offset != "":
		messages, err = h.chatPostUsecase.FindByRoomIDAndLimitAndOffset(roomID, limitStr, offset)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, messages)
}

func (h *chatPostHandler) CreateChatPost(c *gin.Context) {
	err := h.chatPostUsecase.Insert(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrChatPostCreateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"code": 2, "message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 3, "message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 4, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 99, "message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create chatpost"})
}
