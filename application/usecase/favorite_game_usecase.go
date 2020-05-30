package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbres"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// FavoriteGameUsecase :
type FavoriteGameUsecase interface {
	FindFavoriteGameByID(id string) ([]*dbres.FavoriteGamesRes, error)
	InsertFavoriteGame(favgame *dbmodel.FavoriteGame) error
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

func (fU favoriteGameUsecase) FindFavoriteGameByID(id string) ([]*dbres.FavoriteGamesRes, error) {
	favoriteGames, err := fU.favoriteGameRepository.FindFavoriteGameByID(id)
	if err != nil {
		return nil, err
	}
	return favoriteGames, nil
}

func (fU favoriteGameUsecase) InsertFavoriteGame(favgame *dbmodel.FavoriteGame) error {
	err := fU.favoriteGameRepository.InsertFavoriteGame(favgame)
	if err != nil {
		return err
	}
	return nil
}

func (fU favoriteGameUsecase) DeleteFavoriteGame(uID, fID string) error {
	err := fU.favoriteGameRepository.DeleteFavoriteGame(uID, fID)
	if err != nil {
		return err
	}
	return nil
}
