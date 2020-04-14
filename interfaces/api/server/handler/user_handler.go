package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// UserHandler : インターフェース
type UserHandler interface {
	Login(*gin.Context)
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

// Login : GET /user/me-> ユーザーのデータを返す
func (uh userHandler) Login(c *gin.Context) {

}
