package handler

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
)

// RoomHandler : インターフェース
type RoomHandler interface {
	GetRoom(*gin.Context)
	GetRoomByID(*gin.Context)
	CreateRoom(*gin.Context)
	GetBlacklist(*gin.Context)
	CreateBlacklist(*gin.Context)
}

type roomHandler struct {
	userUsecase            usecase.UserUseCase
	roomUsecase            usecase.RoomUseCase
	roomBlacklistUseCase   usecase.RoomBlacklistUseCase
	roomReservationUseCase usecase.RoomReservationUseCase
}

// NewRoomHandler : ユーザーのHandler生成
func NewRoomHandler(
	uU usecase.UserUseCase,
	rU usecase.RoomUseCase,
	rBU usecase.RoomBlacklistUseCase,
	rRU usecase.RoomReservationUseCase) RoomHandler {
	return &roomHandler{
		userUsecase:            uU,
		roomUsecase:            rU,
		roomBlacklistUseCase:   rBU,
		roomReservationUseCase: rRU,
	}
}

// GenerateToken : aa
func GenerateToken() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(10000)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("keyData"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (rH roomHandler) GetRoom(c *gin.Context) {
	token := GenerateToken()
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (rH roomHandler) GetRoomByID(c *gin.Context) {

}

func (rH roomHandler) CreateRoom(c *gin.Context) {

}

func (rH roomHandler) GetBlacklist(c *gin.Context) {

}

func (rH roomHandler) CreateBlacklist(c *gin.Context) {

}
