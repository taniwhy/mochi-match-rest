package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/models"
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

func (cH chatPostHandler) GetChatPostByRoomID(c *gin.Context) {
	limitStr := c.Query("limit")
	offset := c.Query("offset")
	roomID := c.Params.ByName("id")
	// todo : limitのint変換処理の位置を要修正
	switch {
	case limitStr == "" && offset == "":
		chats, err := cH.chatPostUsecase.FindChatPostByRoomID(roomID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, chats)
	case offset == "":
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		chats, err := cH.chatPostUsecase.FindChatPostByRoomIDAndLimit(roomID, limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, chats)
	default:
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		chats, err := cH.chatPostUsecase.FindChatPostByRoomIDAndLimitAndOffset(roomID, offset, limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, chats)
	}
}

func (cH chatPostHandler) CreateChatPost(c *gin.Context) {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	roomID := c.Params.ByName("id")
	// todo : テスト用に仮データを記述
	// idをトークンから取得できるように
	m := models.ChatPost{
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
	if err := cH.chatPostUsecase.InsertChatPost(&m); err != nil {
		c.JSON(http.StatusBadRequest, err)
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
