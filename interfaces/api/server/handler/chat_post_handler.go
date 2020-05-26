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

}

func (cH chatPostHandler) CreateChatPost(c *gin.Context) {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	room, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	m := models.ChatPost{
		ChatPostID: id.String(),
		Room:       room,
		UserID:     123,
		CreatedAt:  time.Now(),
	}
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	cH.chatPostUsecase.InsertChatPost(&m)

	res, _ := json.Marshal(m)
	_, err = cH.redis.Do("PUBLISH", "channel_1", string(res))
	if err != nil {
		panic(err)
	}
}
