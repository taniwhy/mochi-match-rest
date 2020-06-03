package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
)

// UserHandler : インターフェース
type UserHandler interface {
	GetMe(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUseCase
}

// NewUserHandler : UserHandler生成
func NewUserHandler(uU usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUsecase: uU,
	}
}

func (uH userHandler) GetMe(c *gin.Context) {
	u, err := uH.userUsecase.GetMe(c)
	switch err := err.(type) {
	case errors.ErrGetTokenClaims:
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case errors.ErrRecordNotFound:
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case errors.ErrDataBase:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (uH userHandler) GetByID(c *gin.Context) {
	u, err := uH.userUsecase.GetByID(c)
	switch err := err.(type) {
	case errors.ErrGetTokenClaims:
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case errors.ErrRecordNotFound:
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	case errors.ErrDataBase:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (uH userHandler) Create(c *gin.Context) {
	uD, err := uH.userUsecase.Create(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrCreateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrCoockie:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrIDAlreadyExists:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrDataBase:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		case errors.ErrGenerateID:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			panic(err)
		}
	}
	accessToken := auth.GenerateAccessToken(uD.UserID)
	refleshToken, exp := auth.GenerateRefreshToken(uD.UserID)
	c.JSON(http.StatusOK, gin.H{
		"id":            uD.UserID,
		"access_token":  accessToken,
		"refresh_token": refleshToken,
		"expires_in":    exp,
	})
}

func (uH userHandler) Update(c *gin.Context) {
	err := uH.userUsecase.Update(c)
	if err != nil {
		switch err := err.(type) {
		case errors.ErrUpdateReqBinding:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrGetTokenClaims:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		case errors.ErrParams:
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
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated user"})
}

func (uH userHandler) Delete(c *gin.Context) {
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
	if err := uH.userUsecase.Delete(claimsID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
