//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"encoding/json"
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
	FindByRoomID(id string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimit(id, limit string) ([]*models.ChatPost, error)
	FindByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error)
	FindByRoomIDAndLimitAndOffset(id, offset, limit string) ([]*models.ChatPost, error)
	Insert(*gin.Context) error
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

func (cU chatPostUsecase) FindByRoomID(id string) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindByRoomID(id)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindByRoomIDAndLimit(id, limitStr string) ([]*models.ChatPost, error) {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
	chatposts, err := cU.chatPostRepository.FindByRoomIDAndLimit(id, limit)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindByRoomIDAndOffset(id, offset string) ([]*models.ChatPost, error) {
	chatposts, err := cU.chatPostRepository.FindByRoomIDAndOffset(id, offset)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) FindByRoomIDAndLimitAndOffset(id, limitStr, offset string) ([]*models.ChatPost, error) {
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil, err
	}
	chatposts, err := cU.chatPostRepository.FindByRoomIDAndLimitAndOffset(id, offset, limit)
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (cU chatPostUsecase) Insert(c *gin.Context) error {
	b := input.ChatPostCreateReqBody{}
	if err := c.BindJSON(&b); err != nil {
		return errors.ErrChatPostCreateReqBinding{
			Message: b.Message,
		}
	}
	rID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaimsFromRequest(c)
	if err != nil {
		return errors.ErrGetTokenClaims{Detail: err.Error()}
	}
	uID := claims["sub"].(string)
	cP, err := models.NewChatPost(rID, uID, b.Message)
	if err := cU.chatPostRepository.Insert(cP); err != nil {
		return err
	}
	res, _ := json.Marshal(cP)
	// todo : publishのチャンネル名がハードコーディングされているため要修正
	_, err = cU.redis.Do("PUBLISH", "channel_1", string(res))
	if err != nil {
		panic(err)
	}
	return nil
}
