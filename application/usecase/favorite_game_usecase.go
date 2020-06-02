package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// FavoriteGameUsecase :
type FavoriteGameUsecase interface {
	FindFavoriteGameByID(id string) ([]*models.FavoriteGame, error)
	InsertFavoriteGame(favgame *models.FavoriteGame) error
	DeleteFavoriteGame(uID, fID string) error
}

type favoriteGameUsecase struct {
	favoriteGameRepository repository.FavoriteGameRepository
}

// NewFavoriteGameUsecase :
func NewFavoriteGameUsecase(fR repository.FavoriteGameRepository) FavoriteGameUsecase {
	return &favoriteGameUsecase{
		favoriteGameRepository: fR,
	}
}

func (fU favoriteGameUsecase) FindFavoriteGameByID(id string) ([]*models.FavoriteGame, error) {
	favoriteGames, err := fU.favoriteGameRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return favoriteGames, nil
}

func (fU favoriteGameUsecase) InsertFavoriteGame(favgame *models.FavoriteGame) error {
	err := fU.favoriteGameRepository.Insert(favgame)
	if err != nil {
		return err
	}
	return nil
}

func (fU favoriteGameUsecase) DeleteFavoriteGame(uID, fID string) error {
	err := fU.favoriteGameRepository.Delete(uID, fID)
	if err != nil {
		return err
	}
	return nil
}
