package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// IGameListHandler : インターフェース
type IGameListHandler interface {
	GetAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type gameListHandler struct {
	gameListUsecase usecase.IGameListUseCase
}

// NewGameListHandler : ゲームタイトルハンドラの生成
func NewGameListHandler(gU usecase.IGameListUseCase) IGameListHandler {
	return &gameListHandler{
		gameListUsecase: gU,
	}
}

func (gH gameListHandler) GetAll(c *gin.Context) {
	gameTitles, err := gH.gameListUsecase.FindAll(c)
	if err != nil {
		switch err := err.(type) {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gameTitles)
}

func (gH gameListHandler) Create(c *gin.Context) {
	err := gH.gameListUsecase.Insert(c)
	if err != nil {
		switch err := err.(type) {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Created gametitle"})
}

func (gH gameListHandler) Update(c *gin.Context) {
	err := gH.gameListUsecase.Update(c)
	if err != nil {
		switch err := err.(type) {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated gametitle"})
}

func (gH gameListHandler) Delete(c *gin.Context) {
	err := gH.gameListUsecase.Delete(c)
	if err != nil {
		switch err := err.(type) {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted gametitle"})
}
