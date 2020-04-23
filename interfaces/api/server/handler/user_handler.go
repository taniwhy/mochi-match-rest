package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// UserHandler : インターフェース
type UserHandler interface {
	GetUser(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userHandler struct {
	userUsecase       usecase.UserUseCase
	userDetailUsecase usecase.UserDetailUseCase
}

// NewUserHandler : ユーザーのHandler生成
func NewUserHandler(uU usecase.UserUseCase, uDU usecase.UserDetailUseCase) UserHandler {
	return &userHandler{
		userUsecase:       uU,
		userDetailUsecase: uDU,
	}
}

func (uh userHandler) GetUser(c *gin.Context) {

}

func (uh userHandler) CreateUser(c *gin.Context) {

}

func (uh userHandler) UpdateUser(c *gin.Context) {

}

func (uh userHandler) DeleteUser(c *gin.Context) {

}
