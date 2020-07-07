package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"golang.org/x/oauth2"

	log "github.com/sirupsen/logrus"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// IGoogleOAuthHandler : インターフェース
type IGoogleOAuthHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type googleOAuthHandler struct {
	oauthConf          *oauth2.Config
	googleOAuthUsecase usecase.IGoogleOAuthUsecase
	userUsecase        usecase.IUserUseCase
	userService        service.IUserService
}

// NewGoogleOAuthHandler : GoogleOAuth認証ハンドラの生成
func NewGoogleOAuthHandler(
	gU usecase.IGoogleOAuthUsecase,
	uU usecase.IUserUseCase,
	uS service.IUserService) IGoogleOAuthHandler {
	return &googleOAuthHandler{
		oauthConf:          config.GetOAuthClientConf(),
		googleOAuthUsecase: gU,
		userUsecase:        uU,
		userService:        uS,
	}
}

func (gA *googleOAuthHandler) Login(c *gin.Context) {
	url, err := gA.googleOAuthUsecase.Login(c)
	if err != nil {

	}
	c.JSON(http.StatusOK, url)
}

func (gA *googleOAuthHandler) Callback(c *gin.Context) {
	var u *models.User
	ok, gU, err := gA.googleOAuthUsecase.Callback(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Error("StatusBadRequest")
		return
	}
	if ok {
		b := input.UserCreateReqBody{
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
				log.Error("StatusBadRequest")
				return
			case errors.ErrCoockie:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				log.Error("StatusBadRequest")
				return
			case errors.ErrUnexpectedQueryProvider:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				log.Error("StatusBadRequest")
				return
			case errors.ErrIDAlreadyExists:
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				log.Error("StatusBadRequest")
				return
			case errors.ErrDataBase:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				log.Error("StatusInternalServerError")
				return
			case errors.ErrGenerateID:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				log.Error("StatusInternalServerError")
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
				log.Error("StatusBadRequest")
				return
			case errors.ErrDataBase:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				log.Error("StatusInternalServerError")
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				log.Warnf("Unexpected error: %s", err.Error())
				panic(err)
			}
		}
	}
	refleshToken, exp := auth.GenerateRefreshToken(u.UserID)
	c.SetCookie("token", refleshToken, 0, "/", "", false, true)
	c.SetCookie("token_exp", exp, 0, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4000/login-done")
}
