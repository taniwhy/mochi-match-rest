package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/models"
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
	gameTitles, err := gH.gameTitleUsecase.FindAllGameTitle()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gameTitles)
}

func (gH gameTitleHandler) CreateGameTitle(c *gin.Context) {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	g := &models.GameTitle{
		GameTitleID: id.String(),
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
	if err := c.BindJSON(&g); err != nil {
		// todo : エラーメッセージを要修正
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if g.GameTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding error"})
		return
	}
	if err := gH.gameTitleUsecase.InsertGameTitle(g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, g)
}

func (gH gameTitleHandler) UpdateGameTitle(c *gin.Context) {
	gameTitleID := c.Params.ByName("id")
	g := &models.GameTitle{
		GameTitleID: gameTitleID,
		UpdateAt:    time.Now(),
	}
	if err := c.BindJSON(&g); err != nil {
		// todo : エラーメッセージを要修正
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	// todo : エラーメッセージ要修正
	if g.GameTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding error"})
		return
	}
	if err := gH.gameTitleUsecase.UpdateGameTitle(g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.Status(http.StatusOK)
}

func (gH gameTitleHandler) DeleteGameTitle(c *gin.Context) {
	gameTitles, err := gH.gameTitleUsecase.FindAllGameTitle()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gameTitles)
}
