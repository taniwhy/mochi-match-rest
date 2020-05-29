package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
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

// SignupReqBody : リクエストボディのマッピングに使用
type SignupReqBody struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

// NewUserHandler : UserHandler生成
func NewUserHandler(uU usecase.UserUseCase, uDU usecase.UserDetailUseCase) UserHandler {
	return &userHandler{
		userUsecase:       uU,
		userDetailUsecase: uDU,
	}
}

func (uH userHandler) GetUser(c *gin.Context) {

}

func (uH userHandler) CreateUser(c *gin.Context) {
	signupReq := SignupReqBody{}
	if err := c.Bind(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding error"})
		return
	}
	if signupReq.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user_name is required"})
		return
	}
	if signupReq.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email is required"})
		return
	}
	pid, err := c.Cookie("pid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cockie error"})
		return
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	u := &models.User{
		UserID:     uid.String(),
		GoogleID:   sql.NullString{String: "", Valid: false},
		FacebookID: sql.NullString{String: "", Valid: false},
		TwitterID:  sql.NullString{String: "", Valid: false},
		Email:      signupReq.Email,
		IsAdmin:    false,
		IsFreeze:   false,
		IsDelete:   false,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
		DeleteAt:   sql.NullTime{Time: time.Now(), Valid: false},
	}
	provider := c.Query("provider")
	switch {
	case provider == "google":
		res, _ := uH.userUsecase.FindUserByProviderID("google", pid)
		if res != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("ID alreday exists: %v", pid)})
			return
		}
		u.GoogleID = sql.NullString{String: pid, Valid: true}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Unexpected query provider: %v", provider)})
		return
	}
	udid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	uD := &models.UserDetail{
		UserDetailID: udid.String(),
		UserID:       uid.String(),
		UserName:     signupReq.UserName,
		Icon:         1,
		UpdateAt:     time.Now(),
	}
	if err := uH.userUsecase.CreateUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uH.userDetailUsecase.CreateUserDetail(uD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	accessToekn := auth.GenerateAccessToken(uD.UserID)
	refleshToken, exp := auth.GenerateRefreshToken(uD.UserID)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToekn,
		"refresh_token": refleshToken,
		"expires_in":    exp,
	})
}

func (uH userHandler) UpdateUser(c *gin.Context) {

}

func (uH userHandler) DeleteUser(c *gin.Context) {
	userID := c.Params.ByName("id")
	claims, err := auth.GetTokenClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	claimsID := claims["sub"].(string)
	if userID != claimsID {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Params error: %v", userID)})
		return
	}
	if err := uH.userUsecase.DeleteUser(claimsID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}
