package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// ChatPostUseCase :
type ChatPostUseCase interface {
	FindAllChatPost() ([]*models.ChatPost, error)
	FindChatPostByRoomID(id int64) ([]*models.ChatPost, error)
	InsertChatPost(room *models.ChatPost) error
	DeleteChatPost(room *models.ChatPost) error
}

type chatPostUsecase struct {
	chatPostRepository repository.ChatPostRepository
}

// NewChatPostUsecase :
func NewChatPostUsecase(rR repository.ChatPostRepository) ChatPostUseCase {
	return &chatPostUsecase{
		chatPostRepository: rR,
	}
}

func (cU chatPostUsecase) FindAllChatPost() ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindAllChatPost()
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomID(id int64) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomID(id)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) InsertChatPost(chatpost *models.ChatPost) error {
	err := cU.chatPostRepository.InsertChatPost(chatpost)
	if err != nil {
		return err
	}
	return nil
}

func (cU chatPostUsecase) DeleteChatPost(chatpost *models.ChatPost) error {
	err := cU.chatPostRepository.DeleteChatPost(chatpost)
	if err != nil {
		return err
	}
	return nil
}
