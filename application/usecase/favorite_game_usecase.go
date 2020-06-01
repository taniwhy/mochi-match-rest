package usecase

import (
<<<<<<< HEAD
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/response"
=======
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbres"
>>>>>>> a65084717e19b07f7c1f8955dec602840717615a
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// FavoriteGameUsecase :
type FavoriteGameUsecase interface {
<<<<<<< HEAD
	FindFavoriteGameByID(id string) ([]*response.FavoriteGamesRes, error)
	InsertFavoriteGame(favgame *models.FavoriteGame) error
=======
	FindFavoriteGameByID(id string) ([]*dbres.FavoriteGamesRes, error)
	InsertFavoriteGame(favgame *dbmodel.FavoriteGame) error
>>>>>>> a65084717e19b07f7c1f8955dec602840717615a
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

func (fU favoriteGameUsecase) InsertFavoriteGame(favgame *models.FavoriteGame) error {
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
