package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// GameTitleHandler : インターフェース
type GameTitleHandler interface {
	GetAllGameTitle(*gin.Context)
	CreateGameTitle(*gin.Context)
	UpdateGameTitle(*gin.Context)
	DeleteGameTitle(*gin.Context)
}

type gameTitleHandler struct {
	gameTitleUsecase usecase.GameTitleUseCase
}

// NewGameTitleHandler : ユーザーのHandler生成
func NewGameTitleHandler(gU usecase.GameTitleUseCase) GameTitleHandler {
	return &gameTitleHandler{
		gameTitleUsecase: gU,
	}
}

func (gH gameTitleHandler) GetAllGameTitle(c *gin.Context) {

}

func (gH gameTitleHandler) CreateGameTitle(c *gin.Context) {

}

func (gH gameTitleHandler) UpdateGameTitle(c *gin.Context) {

}

func (gH gameTitleHandler) DeleteGameTitle(c *gin.Context) {

}
