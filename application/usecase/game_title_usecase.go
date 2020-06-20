package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IGameTitleUseCase : インターフェース
type IGameTitleUseCase interface {
	FindAllGameTitle() ([]*models.GameTitle, error)
	InsertGameTitle(gameTitle *models.GameTitle) error
	UpdateGameTitle(gameTitle *models.GameTitle) error
	DeleteGameTitle(gameTitle *models.GameTitle) error
}

type gameTitleUsecase struct {
	gameTitleRepository repository.GameTitleRepository
}

// NewGameTitleUsecase : GameTitleユースケースの生成
func NewGameTitleUsecase(gR repository.GameTitleRepository) IGameTitleUseCase {
	return &gameTitleUsecase{
		gameTitleRepository: gR,
	}
}

func (gU gameTitleUsecase) FindAllGameTitle() ([]*models.GameTitle, error) {
	gameTitles, err := gU.gameTitleRepository.FindAllGameTitle()
	if err != nil {
		return nil, err
	}
	return gameTitles, nil
}

func (gU gameTitleUsecase) InsertGameTitle(gameTitle *models.GameTitle) error {
	err := gU.gameTitleRepository.InsertGameTitle(gameTitle)
	if err != nil {
		return err
	}
	return nil
}

func (gU gameTitleUsecase) UpdateGameTitle(gameTitle *models.GameTitle) error {
	err := gU.gameTitleRepository.UpdateGameTitle(gameTitle)
	if err != nil {
		return err
	}
	return nil
}

func (gU gameTitleUsecase) DeleteGameTitle(gameTitle *models.GameTitle) error {
	err := gU.gameTitleRepository.DeleteGameTitle(gameTitle)
	if err != nil {
		return err
	}
	return nil
}
