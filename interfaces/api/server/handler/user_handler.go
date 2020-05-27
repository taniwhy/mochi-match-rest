package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// UserHandler : インターフェース
type UserHandler interface {
	GetUser(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userHandler struct {
	userUsecase       usecase.UserUseCase
	userDetailUsecase usecase.UserDetailUseCase
}

// NewUserHandler : ユーザーのHandler生成
func NewUserHandler(uU usecase.UserUseCase, uDU usecase.UserDetailUseCase) UserHandler {
	return &userHandler{
		userUsecase:       uU,
		userDetailUsecase: uDU,
	}
}

func (uH userHandler) GetUser(c *gin.Context) {

}

func (uH userHandler) CreateUser(c *gin.Context) {
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	u := &models.User{
		UserID:     uid.String(),
		GoogleID:   sql.NullString{"", false},
		FacebookID: sql.NullString{"", false},
		TwitterID:  sql.NullString{"", false},
		IsAdmin:    false,
		IsFreeze:   false,
		IsDelete:   false,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
		DeleteAt:   sql.NullTime{time.Now(), false},
	}
	provider := c.Query("provider")
	switch {
	case provider == "google":
		goid, err := c.Cookie("goid")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Cockie error"})
			return
		}
		u.GoogleID = sql.NullString{goid, true}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Query error"})
		return
	}
	udid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	uD := &models.UserDetail{
		UserDetailID: udid.String(),
		UserID:       uid.String(),
		Icon:         1,
		UpdateAt:     time.Now(),
	}
	if err := c.BindJSON(&uD); err != nil {
		// todo : エラーメッセージを要修正
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if uD.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding error"})
		return
	}
	if err := uH.userUsecase.CreateUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if err := uH.userDetailUsecase.CreateUserDetail(uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": "testToken"})
}

func (uH userHandler) UpdateUser(c *gin.Context) {

}

func (uH userHandler) DeleteUser(c *gin.Context) {

}
