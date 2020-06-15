package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"golang.org/x/oauth2"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// GoogleOAuthHandler : todo
type GoogleOAuthHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type googleOAuthHandler struct {
	oauthConf   *oauth2.Config
	userUsecase usecase.UserUseCase
	userService service.IUserService
}

// NewGoogleOAuthHandler :
func NewGoogleOAuthHandler(uU usecase.UserUseCase, uS service.IUserService) GoogleOAuthHandler {
	return &googleOAuthHandler{
		oauthConf:   config.ConfigureOAuthClient(),
		userUsecase: uU,
		userService: uS,
	}
}

func (gA *googleOAuthHandler) Login(c *gin.Context) {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err.Error())
	}
	sessionID := u.String()

	session := sessions.Default(c)
	session.Set("state", sessionID)
	session.Save()

	url := gA.oauthConf.AuthCodeURL(sessionID)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (gA *googleOAuthHandler) Callback(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}
	tok, err := gA.oauthConf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if tok.Valid() == false {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	client := gA.oauthConf.Client(oauth2.NoContext, tok)
	email, err := client.Get(oauthGoogleURLAPI)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()

	data, err := ioutil.ReadAll(email.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gU := models.GoogleUser{}
	err = json.Unmarshal(data, &gU)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// todo : errのpanic処理
	ok, err := gA.userService.IsExist("google", gU.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	if ok {
		c.SetCookie("pid", gU.ID, 0, "/", "", false, true)
		c.SetCookie("name", gU.Name, 0, "/", "localhost", false, true)
		c.SetCookie("mail", gU.Email, 0, "/", "localhost", false, true)
		// ユーザー登録ページにリダイレクト
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4000/sign-up")
		return
	}
	// ログインしリダイレクト
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4000/profilesetting")
}
