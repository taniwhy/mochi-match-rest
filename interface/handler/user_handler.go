package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/css-study-game-WebAPI/usecase"
)

// UserHandler : インターフェース
type UserHandler interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// NewUserHandler : ユーザーのHandler生成
func NewUserHandler(uU usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uU,
	}
}

// GetUser : GET /user/me-> ユーザーのデータを返す
func (uh userHandler) GetUser(c *gin.Context) {

}
