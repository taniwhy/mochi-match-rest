package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"golang.org/x/oauth2"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// GoogleOAuthHandler : todo
type GoogleOAuthHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type googleOAuthHandler struct {
	oauthConf          *oauth2.Config
	googleOAuthUsecase usecase.GoogleOAuthUsecase
	userUsecase        usecase.UserUseCase
	userService        service.IUserService
}

// NewGoogleOAuthHandler :
func NewGoogleOAuthHandler(
	gU usecase.GoogleOAuthUsecase,
	uU usecase.UserUseCase,
	uS service.IUserService) GoogleOAuthHandler {
	return &googleOAuthHandler{
		oauthConf:          config.ConfigureOAuthClient(),
		googleOAuthUsecase: gU,
		userUsecase:        uU,
		userService:        uS,
	}
}

func (gA *googleOAuthHandler) Login(c *gin.Context) {
	url, err := gA.googleOAuthUsecase.Login(c)
	if err != nil {

	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (gA *googleOAuthHandler) Callback(c *gin.Context) {
	var u *models.User
	ok, gU, err := gA.googleOAuthUsecase.Callback(c)
	if err != nil {

	}
	// ない
	if ok {
		b := input.UserCreateBody{
			Provider:   "google",
			ProviderID: gU.ID,
			UserName:   gU.Name,
			Email:      gU.Email,
		}
		u, err = gA.userUsecase.Create(c, b)
		if err != nil {
			switch err := err.(type) {
			case errors.ErrUserCreateReqBinding:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			case errors.ErrCoockie:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			case errors.ErrUnexpectedQueryProvider:
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
				log.Warn("Unexpected error")
				panic(err)
			}
		}
	} else {
		u, err = gA.userUsecase.GetByProviderID("google", gU.ID)
		if err != nil {
			switch err := err.(type) {
			case errors.ErrUnexpectedQueryProvider:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				return
			case errors.ErrDataBase:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				log.Warnf("Unexpected error: %s", err.Error())
				panic(err)
			}
		}
	}
	accessToken := auth.GenerateAccessToken(u.UserID)
	refleshToken, exp := auth.GenerateRefreshToken(u.UserID)

	session := sessions.Default(c)
	session.Set("access_token", accessToken)
	session.Set("refresh_token", refleshToken)
	session.Set("exp", exp)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"id":            u.UserID,
		"access_token":  accessToken,
		"refresh_token": refleshToken,
		"expires_in":    exp,
	})
}
