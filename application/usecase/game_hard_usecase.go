package usecase

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IGameHardUseCase : インターフェース
type IGameHardUseCase interface {
	FindAll(*gin.Context) ([]*models.GameHard, error)
	Insert(*gin.Context) error
	Update(*gin.Context) error
	Delete(*gin.Context) error
}

type gameHardUsecase struct {
	gameHardRepository repository.GameHardRepository
}

// NewGameHardUsecase : GameTitleユースケースの生成
func NewGameHardUsecase(gR repository.GameHardRepository) IGameHardUseCase {
	return &gameHardUsecase{
		gameHardRepository: gR,
	}
}

func (gU gameHardUsecase) FindAll(c *gin.Context) ([]*models.GameHard, error) {
	gameHards, err := gU.gameHardRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return gameHards, nil
}

func (gU gameHardUsecase) Insert(c *gin.Context) error {
	b := input.GameHardCreateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrGameHardCreateReqBinding{HardName: b.HardName}
	}
	gH, err := models.NewGameHard(b.HardName)
	if err != nil {
		return err
	}
	if err := gU.gameHardRepository.Insert(gH); err != nil {
		return err
	}
	return nil
}

func (gU gameHardUsecase) Update(c *gin.Context) error {
	gid := c.Params.ByName("id")
	b := input.GameHardUpdateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrGameHardUpdateReqBinding{HardName: b.HardName}
	}
	gH := &models.GameHard{
		GameHardID: gid,
		HardName:   b.HardName,
		UpdateAt:   time.Now(),
	}
	if err := gU.gameHardRepository.Update(gH); err != nil {
		return err
	}
	return nil
}

func (gU gameHardUsecase) Delete(c *gin.Context) error {
	id := c.Params.ByName("id")
	gH := &models.GameHard{
		GameHardID: id,
	}
	err := gU.gameHardRepository.Delete(gH)
	if err != nil {
		return err
	}
	return nil
}
