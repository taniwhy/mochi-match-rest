package usecase

import (
<<<<<<< HEAD
	"github.com/taniwhy/mochi-match-rest/domain/models"
<<<<<<< HEAD
=======
	"github.com/taniwhy/mochi-match-rest/domain/models/response"
=======
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbres"
>>>>>>> a65084717e19b07f7c1f8955dec602840717615a
>>>>>>> 9a70d396a44f4ff618d89b4dafe030a585b76c6f
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// FavoriteGameUsecase :
type FavoriteGameUsecase interface {
<<<<<<< HEAD
	FindFavoriteGameByID(id string) ([]*models.FavoriteGame, error)
=======
<<<<<<< HEAD
	FindFavoriteGameByID(id string) ([]*response.FavoriteGamesRes, error)
>>>>>>> 9a70d396a44f4ff618d89b4dafe030a585b76c6f
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

<<<<<<< HEAD
func (fU favoriteGameUsecase) FindFavoriteGameByID(id string) ([]*models.FavoriteGame, error) {
	favoriteGames, err := fU.favoriteGameRepository.FindByID(id)
=======
func (fU favoriteGameUsecase) FindFavoriteGameByID(id string) ([]*dbres.FavoriteGamesRes, error) {
	favoriteGames, err := fU.favoriteGameRepository.FindFavoriteGameByID(id)
>>>>>>> 9a70d396a44f4ff618d89b4dafe030a585b76c6f
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
