package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"

	log "github.com/sirupsen/logrus"
)

// IRoomHandler : インターフェース
type IRoomHandler interface {
	GetList(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Join(*gin.Context)
	Leave(*gin.Context)
	CheckEntry(*gin.Context)
}

type roomHandler struct {
	userUsecase usecase.IUserUseCase
	roomUsecase usecase.IRoomUseCase
}

// NewRoomHandler : ルームハンドラの生成
func NewRoomHandler(
	uU usecase.IUserUseCase,
	rU usecase.IRoomUseCase) IRoomHandler {
	return &roomHandler{
		userUsecase: uU,
		roomUsecase: rU,
	}
}

func (h *roomHandler) GetList(c *gin.Context) {
	r, err := h.roomUsecase.GetList(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h *roomHandler) GetByID(c *gin.Context) {
	roomDetail, err := h.roomUsecase.GetByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, roomDetail)
}

func (h *roomHandler) Create(c *gin.Context) {
	err := h.roomUsecase.Create(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrRoomCreateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrRoomAlreadyExists:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrRoomAlreadyEntry:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create room"})
}

func (h *roomHandler) Update(c *gin.Context) {

}

func (h *roomHandler) Delete(c *gin.Context) {
	err := h.roomUsecase.Delete(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrNotRoomOwner:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete room"})
}

func (h *roomHandler) Join(c *gin.Context) {
	err := h.roomUsecase.Join(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": err.Error()})
			return
		case errors.ErrNotRoomOwner:
			c.JSON(http.StatusBadRequest, gin.H{"code": 2, "message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"code": 3, "message": err.Error()})
			return
		case errors.ErrRoomAlreadyEntry:
			c.JSON(http.StatusBadRequest, gin.H{"code": 4, "message": err.Error()})
			return
		case errors.ErrRoomCapacityOver:
			c.JSON(http.StatusBadRequest, gin.H{"code": 5, "message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 6, "message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 7, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"code": 99, "message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Join room"})
}

func (h *roomHandler) Leave(c *gin.Context) {
	err := h.roomUsecase.Leave(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrNotEntryRoom:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Leave room"})
}

func (h *roomHandler) CheckEntry(c *gin.Context) {
	isEntry, room, err := h.roomUsecase.CheckEntry(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrNotEntryRoom:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrRecordNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			log.Warn("Unexpected error")
			panic(err)
		}
	}
	if isEntry {
		c.JSON(http.StatusOK, gin.H{"message": "Already entry", "room": room})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Not entry"})
}
