package usecase

import (
	"strconv"

	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// ChatPostUseCase :
type ChatPostUseCase interface {
	FindAllChatPost() ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomID(id string) ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomIDAndLimit(id, limit string) ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomIDAndOffset(id, offset string) ([]*dbmodel.ChatPost, error)
	FindChatPostByRoomIDAndLimitAndOffset(id, offset, limit string) ([]*dbmodel.ChatPost, error)
	InsertChatPost(room *dbmodel.ChatPost) error
	DeleteChatPost(room *dbmodel.ChatPost) error
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

func (cU chatPostUsecase) FindAllChatPost() ([]*dbmodel.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindAllChatPost()
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomID(id string) ([]*dbmodel.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomID(id)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomIDAndLimit(id, limitStr string) ([]*dbmodel.ChatPost, error) {
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

func (cU chatPostUsecase) FindChatPostByRoomIDAndOffset(id, offset string) ([]*dbmodel.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindChatPostByRoomIDAndOffset(id, offset)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindChatPostByRoomIDAndLimitAndOffset(id, limitStr, offset string) ([]*dbmodel.ChatPost, error) {
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

func (cU chatPostUsecase) InsertChatPost(chatpost *dbmodel.ChatPost) error {
	err := cU.chatPostRepository.InsertChatPost(chatpost)
	if err != nil {
		return err
	}
	return nil
}

func (cU chatPostUsecase) DeleteChatPost(chatpost *dbmodel.ChatPost) error {
	err := cU.chatPostRepository.DeleteChatPost(chatpost)
	if err != nil {
		return err
	}
	return nil
}
