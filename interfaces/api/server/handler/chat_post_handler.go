package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// ChatPostHandler : インターフェース
type ChatPostHandler interface {
	GetChatPostByRoomID(*gin.Context)
	CreateChatPost(*gin.Context)
}

type chatPostHandler struct {
	chatPostUsecase usecase.ChatPostUseCase
}

// NewChatPostHandler : ユーザーのHandler生成
func NewChatPostHandler(cU usecase.ChatPostUseCase) ChatPostHandler {
	return &chatPostHandler{chatPostUsecase: cU}
}

func (cH chatPostHandler) GetChatPostByRoomID(c *gin.Context) {

}

func (cH chatPostHandler) CreateChatPost(c *gin.Context) {

}
