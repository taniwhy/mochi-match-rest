//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

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
	FindAll(c *gin.Context) ([]*models.GameHard, error)
	Insert(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type gameHardUsecase struct {
	gameHardRepository repository.IGameHardRepository
}

// NewGameHardUsecase : GameTitleユースケースの生成
func NewGameHardUsecase(gR repository.IGameHardRepository) IGameHardUseCase {
	return &gameHardUsecase{gameHardRepository: gR}
}

func (u *gameHardUsecase) FindAll(c *gin.Context) ([]*models.GameHard, error) {
	gamehards, err := u.gameHardRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return gamehards, nil
}

func (u *gameHardUsecase) Insert(c *gin.Context) error {
	body := input.GameHardCreateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrGameHardCreateReqBinding{HardName: body.HardName}
	}
	gamehard, err := models.NewGameHard(body.HardName)
	if err != nil {
		return err
	}
	if err := u.gameHardRepository.Insert(gamehard); err != nil {
		return err
	}
	return nil
}

func (u *gameHardUsecase) Update(c *gin.Context) error {
	gamehardID := c.Params.ByName("id")
	body := input.GameHardUpdateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrGameHardUpdateReqBinding{HardName: body.HardName}
	}
	gamehard := &models.GameHard{
		GameHardID: gamehardID,
		HardName:   body.HardName,
		UpdateAt:   time.Now(),
	}
	if err := u.gameHardRepository.Update(gamehard); err != nil {
		return err
	}
	return nil
}

func (u *gameHardUsecase) Delete(c *gin.Context) error {
	gamehardID, _ := c.GetQueryArray("id")
	for _, id := range gamehardID {
		gamehard := &models.GameHard{GameHardID: id}
		err := u.gameHardRepository.Delete(gamehard)
		if err != nil {
			return err
		}
	}
	return nil
}
