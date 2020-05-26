package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// ChatPostUseCase :
type ChatPostUseCase interface {
	FindAllChatPost() ([]*models.ChatPost, error)
	FindChatPostByRoomID(id string) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndLimit(id string, limit int) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*models.ChatPost, error)
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

func (cU chatPostUsecase) FindChatPostByRoomID(id string) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomID(id)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomIDAndLimit(id string, limit int) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomIDAndLimit(id, limit)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomIDAndLimitAndOffset(id, offset string, limit int) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomIDAndLimitAndOffset(id, offset, limit)
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
