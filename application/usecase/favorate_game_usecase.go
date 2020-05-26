package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// FavorateGameUsecase :
type FavorateGameUsecase interface {
	FindFavorateGameByID(id string) (*models.FavorateGame, error)
	InsertFavorateGame(favgame *models.FavorateGame) error
	DeleteFavorateGame(favgame *models.FavorateGame) error
}

type favorateGameUsecase struct {
	favorateGameRepository repository.FavorateGameRepository
}

// NewFavorateGameUsecase :
func NewFavorateGameUsecase(fR repository.FavorateGameRepository) FavorateGameUsecase {
	return &favorateGameUsecase{
		favorateGameRepository: fR,
	}
}

func (fU favorateGameUsecase) FindFavorateGameByID(id string) (*models.FavorateGame, error) {
	favorateGames, err := fU.favorateGameRepository.FindFavorateGameByID(id)
	if err != nil {
		return nil, err
	}
	return favorateGames, nil
}

func (fU favorateGameUsecase) InsertFavorateGame(favgame *models.FavorateGame) error {
	err := fU.favorateGameRepository.InsertFavorateGame(favgame)
	if err != nil {
		return err
	}
	return nil
}

func (fU favorateGameUsecase) DeleteFavorateGame(favgame *models.FavorateGame) error {
	err := fU.favorateGameRepository.DeleteFavorateGame(favgame)
	if err != nil {
		return err
	}
	return nil
}
