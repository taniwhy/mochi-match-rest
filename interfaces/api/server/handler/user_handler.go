package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
	"golang.org/x/sync/errgroup"
)

// UserHandler : インターフェース
type UserHandler interface {
	GetUser(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userHandler struct {
	userUsecase         usecase.UserUseCase
	userDetailUsecase   usecase.UserDetailUseCase
	favoriteGameUsecase usecase.FavoriteGameUsecase
}

type favoriteGameRecord struct {
	GameID string `json:"game_id" binding:"required"`
}

type updateReqBody struct {
	UserName      string               `json:"user_name" binding:"required"`
	Icon          int                  `json:"icon" binding:"required"`
	FavoriteGames []favoriteGameRecord `json:"favorite_games" binding:"required"`
}

type getUserResbody struct {
	UserName      string               `json:"user_name" binding:"required"`
	Icon          int                  `json:"icon" binding:"required"`
	FavoriteGames []favoriteGameRecord `json:"favorite_games" binding:"required"`
}

type signupReqBody struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

// NewUserHandler : UserHandler生成
func NewUserHandler(uU usecase.UserUseCase, uDU usecase.UserDetailUseCase, fGU usecase.FavoriteGameUsecase) UserHandler {
	return &userHandler{
		userUsecase:         uU,
		userDetailUsecase:   uDU,
		favoriteGameUsecase: fGU,
	}
}

func (uH userHandler) GetUser(c *gin.Context) {
	claims, err := auth.GetTokenClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	claimsID := claims["sub"].(string)
	_, err = uH.userDetailUsecase.FindUserDetailByID(claimsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}

func (uH userHandler) CreateUser(c *gin.Context) {
	signupReq := signupReqBody{}
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
	u := updateReqBody{}
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
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
	updateUserDetail := models.UserDetail{
		UserID:   claimsID,
		UserName: u.UserName,
		Icon:     u.Icon,
		UpdateAt: time.Now(),
	}
	if err := uH.userDetailUsecase.UpdateUserDetail(&updateUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//
	bfg, err := uH.favoriteGameUsecase.FindFavoriteGameByID(claimsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	beforeFavoriteGames := []favoriteGameRecord{}
	for _, v := range bfg {
		beforeFavoriteGames = append(beforeFavoriteGames, favoriteGameRecord{GameID: v.GameTitleID})
	}
	afterFavoriteGames := u.FavoriteGames
	var insertGames []favoriteGameRecord
	var deleteGames []favoriteGameRecord
	for _, a := range afterFavoriteGames {
		if !contains(beforeFavoriteGames, a.GameID) {
			insertGames = append(insertGames, a)
		}
	}
	for _, a := range beforeFavoriteGames {
		if !contains(afterFavoriteGames, a.GameID) {
			deleteGames = append(deleteGames, a)
		}
	}
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, i := range insertGames {
		id, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}
		f := models.FavoriteGame{
			FavoriteGameID: id.String(),
			UserID:         claimsID,
			GameTitleID:    i.GameID,
			CreatedAt:      time.Now(),
		}
		eg.Go(func() error {
			if err := uH.favoriteGameUsecase.InsertFavoriteGame(&f); err != nil {
				return err
			}
			return nil
		})
	}
	for _, g := range deleteGames {
		eg.Go(func() error {
			if err := uH.favoriteGameUsecase.DeleteFavoriteGame(claimsID, g.GameID); err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		cancel()
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated user"})
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

func contains(gr []favoriteGameRecord, id string) bool {
	for _, r := range gr {
		if id == r.GameID {
			return true
		}
	}
	return false
}
