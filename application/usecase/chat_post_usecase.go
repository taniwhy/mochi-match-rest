//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"

	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IChatPostUseCase : インターフェース
type IChatPostUseCase interface {
	FindByRoomID(roomID string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimit(roomID, limit string) ([]*models.ChatPost, error)
	FindByRoomIDAndOffset(roomID, offset string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimitAndOffset(roomID, offset, limit string) ([]*models.ChatPost, error)
	Insert(c *gin.Context) error
}

type chatPostUsecase struct {
	chatPostRepository repository.IChatPostRepository
	redis              redis.Conn
}

// NewChatPostUsecase : ChatPostユースケースの生成
func NewChatPostUsecase(rR repository.IChatPostRepository, rC redis.Conn) IChatPostUseCase {
	return &chatPostUsecase{
		chatPostRepository: rR,
		redis:              rC,
	}
}

func (u *chatPostUsecase) FindByRoomID(id string) ([]*models.ChatPost, error) {
	chatposts, err := u.chatPostRepository.FindByRoomID(id)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (u *chatPostUsecase) FindByRoomIDAndLimit(id, limitStr string) ([]*models.ChatPost, error) {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
	chatposts, err := u.chatPostRepository.FindByRoomIDAndLimit(id, limit)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (u *chatPostUsecase) FindByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error) {
	chatposts, err := u.chatPostRepository.FindByRoomIDAndOffset(id, offset)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (u *chatPostUsecase) FindByRoomIDAndLimitAndOffset(id, limitStr, offset string) ([]*models.ChatPost, error) {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
	chatposts, err := u.chatPostRepository.FindByRoomIDAndLimitAndOffset(id, offset, limit)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (u *chatPostUsecase) Insert(c *gin.Context) error {
	body := input.ChatPostCreateReqBody{}
	if err := c.BindJSON(&body); err != nil {
		return errors.ErrChatPostCreateReqBinding{
			Message: body.Message,
		}
	}
	roomID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	userID := claims["sub"].(string)
	chatpost, err := models.NewChatPost(roomID, userID, body.Message)
	chatpostRes, err := u.chatPostRepository.Insert(chatpost)
	if err != nil {
		return err
	}
	res, _ := json.Marshal(chatpostRes)
	fmt.Println(string(res))
	// todo : publishのチャンネル名がハードコーディングされているため要修正
	_, err = u.redis.Do("PUBLISH", "create_message", string(res))
	if err != nil {
		panic(err)
	}
	return nil
}
