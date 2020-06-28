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

// IGameListUseCase : インターフェース
type IGameListUseCase interface {
	FindAll(c *gin.Context) ([]*models.GameList, error)
	Insert(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type gameListUsecase struct {
	gameListRepository repository.IGameListRepository
}

// NewGameListUsecase : GameTitleユースケースの生成
func NewGameListUsecase(gR repository.IGameListRepository) IGameListUseCase {
	return &gameListUsecase{
		gameListRepository: gR,
	}
}

func (gU gameListUsecase) FindAll(c *gin.Context) ([]*models.GameList, error) {
	gameTitles, err := gU.gameListRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return gameTitles, nil
}

func (gU gameListUsecase) Insert(c *gin.Context) error {
	b := input.GameListCreateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrGameListCreateReqBinding{GameTitle: b.GameTitle}
	}
	gL, err := models.NewGameList(b.GameTitle)
	if err != nil {
		return err
	}
	if err := gU.gameListRepository.Insert(gL); err != nil {
		return err
	}
	return nil
}

func (gU gameListUsecase) Update(c *gin.Context) error {
	gid := c.Params.ByName("id")
	b := input.GameListUpdateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrGameListUpdateReqBinding{GameTitle: b.GameTitle}
	}
	gL := &models.GameList{
		GameListID: gid,
		GameTitle:  b.GameTitle,
		UpdateAt:   time.Now(),
	}
	if err := gU.gameListRepository.Update(gL); err != nil {
		return err
	}
	return nil
}

func (gU gameListUsecase) Delete(c *gin.Context) error {
	id := c.Params.ByName("id")
	gL := &models.GameList{
		GameListID: id,
	}
	err := gU.gameListRepository.Delete(gL)
	if err != nil {
		return err
	}
	return nil
}
