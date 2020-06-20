package usecase

import (
	"strconv"

	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// IChatPostUseCase : インターフェース
type IChatPostUseCase interface {
	FindAllChatPost() ([]*models.ChatPost, error)
	FindChatPostByRoomID(id string) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndLimit(id, limit string) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error)
	FindChatPostByRoomIDAndLimitAndOffset(id, offset, limit string) ([]*models.ChatPost, error)
	InsertChatPost(room *models.ChatPost) error
	DeleteChatPost(room *models.ChatPost) error
}

type chatPostUsecase struct {
	chatPostRepository repository.ChatPostRepository
}

// NewChatPostUsecase : ChatPostユースケースの生成
func NewChatPostUsecase(rR repository.ChatPostRepository) IChatPostUseCase {
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

func (cU chatPostUsecase) FindChatPostByRoomIDAndLimit(id, limitStr string) ([]*models.ChatPost, error) {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomIDAndLimit(id, limit)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomIDAndOffset(id, offset)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomIDAndLimitAndOffset(id, limitStr, offset string) ([]*models.ChatPost, error) {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
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
