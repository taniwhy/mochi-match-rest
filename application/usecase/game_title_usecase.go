package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// GameTitleUseCase :
type GameTitleUseCase interface {
	FindAllGameTitle() ([]*dbmodel.GameTitle, error)
	InsertGameTitle(gameTitle *dbmodel.GameTitle) error
	UpdateGameTitle(gameTitle *dbmodel.GameTitle) error
	DeleteGameTitle(gameTitle *dbmodel.GameTitle) error
}

type gameTitleUsecase struct {
	gameTitleRepository repository.GameTitleRepository
}

// NewGameTitleUsecase :
func NewGameTitleUsecase(gR repository.GameTitleRepository) GameTitleUseCase {
	return &gameTitleUsecase{
		gameTitleRepository: gR,
	}
}

func (gU gameTitleUsecase) FindAllGameTitle() ([]*dbmodel.GameTitle, error) {
	gameTitles, err := gU.gameTitleRepository.FindAllGameTitle()
	if err != nil {
		return nil, err
	}
	return gameTitles, nil
}

func (gU gameTitleUsecase) InsertGameTitle(gameTitle *dbmodel.GameTitle) error {
	err := gU.gameTitleRepository.InsertGameTitle(gameTitle)
	if err != nil {
		return err
	}
	return nil
}

func (gU gameTitleUsecase) UpdateGameTitle(gameTitle *dbmodel.GameTitle) error {
	err := gU.gameTitleRepository.UpdateGameTitle(gameTitle)
	if err != nil {
		return err
	}
	return nil
}

func (gU gameTitleUsecase) DeleteGameTitle(gameTitle *dbmodel.GameTitle) error {
	err := gU.gameTitleRepository.DeleteGameTitle(gameTitle)
	if err != nil {
		return err
	}
	return nil
}
