package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// ChatPostHandler : インターフェース
type ChatPostHandler interface {
	GetChatPostByRoomID(*gin.Context)
	CreateChatPost(*gin.Context)
}

type chatPostHandler struct {
	chatPostUsecase usecase.ChatPostUseCase
	redis           redis.Conn
}

// NewChatPostHandler : ユーザーのHandler生成
func NewChatPostHandler(cU usecase.ChatPostUseCase, rC redis.Conn) ChatPostHandler {
	return &chatPostHandler{
		chatPostUsecase: cU,
		redis:           rC,
	}
}

var (
	messages []*models.ChatPost
	err      error
)

func (cH chatPostHandler) GetChatPostByRoomID(c *gin.Context) {
	roomID := c.Params.ByName("id")
	limitStr := c.Query("limit")
	offset := c.Query("offset")
	switch {
	case limitStr == "" && offset == "":
		messages, err = cH.chatPostUsecase.FindChatPostByRoomID(roomID)
	case limitStr != "" && offset == "":
		messages, err = cH.chatPostUsecase.FindChatPostByRoomIDAndLimit(roomID, limitStr)
	case limitStr == "" && offset != "":
		messages, err = cH.chatPostUsecase.FindChatPostByRoomIDAndOffset(roomID, offset)
	case limitStr != "" && offset != "":
		messages, err = cH.chatPostUsecase.FindChatPostByRoomIDAndLimitAndOffset(roomID, limitStr, offset)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, messages)
}

func (cH chatPostHandler) CreateChatPost(c *gin.Context) {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	roomID := c.Params.ByName("id")
	// todo : テスト用に仮データを記述
	// idをトークンから取得できるように
	token := auth.GenerateAccessToken("a")
	fmt.Println(token)
	m := &models.ChatPost{
		ChatPostID: id.String(),
		RoomID:     roomID,
		UserID:     "id",
		CreatedAt:  time.Now(),
	}
	if err := c.BindJSON(&m); err != nil {
		// todo : エラーメッセージを要修正
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if m.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "message not found"})
		return
	}
	if err := cH.chatPostUsecase.InsertChatPost(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	res, _ := json.Marshal(m)
	// todo : publishのチャンネル名がハードコーディングされているため要修正
	_, err = cH.redis.Do("PUBLISH", "channel_1", string(res))
	if err != nil {
		panic(err)
	}
	c.Status(http.StatusOK)
}
