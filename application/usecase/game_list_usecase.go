//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/util/clock"
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
	return &gameListUsecase{gameListRepository: gR}
}

func (u *gameListUsecase) FindAll(c *gin.Context) ([]*models.GameList, error) {
	gamelists, err := u.gameListRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return gamelists, nil
}

func (u *gameListUsecase) Insert(c *gin.Context) error {
	body := input.GameListCreateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrGameListCreateReqBinding{GameTitle: body.GameTitle}
	}
	gamelist, err := models.NewGameList(body.GameTitle)
	if err != nil {
		return err
	}
	if err := u.gameListRepository.Insert(gamelist); err != nil {
		return err
	}
	return nil
}

func (u *gameListUsecase) Update(c *gin.Context) error {
	gamelistID := c.Params.ByName("id")
	body := input.GameListUpdateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrGameListUpdateReqBinding{GameTitle: body.GameTitle}
	}
	gamelist := &models.GameList{
		GameListID: gamelistID,
		GameTitle:  body.GameTitle,
		UpdateAt:   clock.Now(),
	}
	if err := u.gameListRepository.Update(gamelist); err != nil {
		return err
	}
	return nil
}

func (u *gameListUsecase) Delete(c *gin.Context) error {
	gamelistID, _ := c.GetQueryArray("id")
	for _, id := range gamelistID {
		gamelist := &models.GameList{GameListID: id}
		err := u.gameListRepository.Delete(gamelist)
		if err != nil {
			return err
		}
	}
	return nil
}
