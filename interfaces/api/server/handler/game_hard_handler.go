package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"

	log "github.com/sirupsen/logrus"
)

// IGameHardHandler : インターフェース
type IGameHardHandler interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type gameHardHandler struct {
	gameHardUsecase usecase.IGameHardUseCase
}

// NewGameHardHandler : ゲームタイトルハンドラの生成
func NewGameHardHandler(gU usecase.IGameHardUseCase) IGameHardHandler {
	return &gameHardHandler{
		gameHardUsecase: gU,
	}
}

func (gH gameHardHandler) GetAll(c *gin.Context) {
	gameTitles, err := gH.gameHardUsecase.FindAll(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gameTitles)
}

func (gH gameHardHandler) Create(c *gin.Context) {
	err := gH.gameHardUsecase.Insert(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrGameHardCreateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Created gametitle"})
}

func (gH gameHardHandler) Update(c *gin.Context) {
	err := gH.gameHardUsecase.Update(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrGameHardUpdateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated gametitle"})
}

func (gH gameHardHandler) Delete(c *gin.Context) {
	err := gH.gameHardUsecase.Delete(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted gametitle"})
}
