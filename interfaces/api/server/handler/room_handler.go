package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
)

// RoomHandler : インターフェース
type RoomHandler interface {
	GetList(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type roomHandler struct {
	userUsecase            usecase.UserUseCase
	roomUsecase            usecase.RoomUseCase
	roomReservationUseCase usecase.RoomReservationUseCase
}

// NewRoomHandler : ユーザーのHandler生成
func NewRoomHandler(
	uU usecase.UserUseCase,
	rU usecase.RoomUseCase,
	rRU usecase.RoomReservationUseCase) RoomHandler {
	return &roomHandler{
		userUsecase:            uU,
		roomUsecase:            rU,
		roomReservationUseCase: rRU,
	}
}

func (rH roomHandler) GetList(c *gin.Context) {
	r, err := rH.roomUsecase.GetList(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

func (rH roomHandler) GetByID(c *gin.Context) {

}

func (rH roomHandler) Create(c *gin.Context) {
	err := rH.roomUsecase.Create(c)
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
	c.JSON(http.StatusOK, gin.H{"message": "created room"})
}

func (rH roomHandler) Update(c *gin.Context) {

}

func (rH roomHandler) Delete(c *gin.Context) {

}
