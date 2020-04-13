package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/config"
	"golang.org/x/oauth2"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// InterfaceGoogleOAuthHandler : todo
type InterfaceGoogleOAuthHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type googleOAuthHandler struct {
	oauthConf *oauth2.Config
}

// NewGoogleOAuthHandler :
func NewGoogleOAuthHandler() InterfaceGoogleOAuthHandler {
	return &googleOAuthHandler{
		oauthConf: config.ConfigureOAuthClient(),
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

	/*
	   AuthCodeURLは、CSRF攻撃からユーザーを保護するトークンである状態を受け取ります。空でない文字列を常に提供する必要があります。
	   リダイレクトコールバックの状態クエリパラメータと一致することを確認します。
	*/
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

	client := gA.oauthConf.Client(oauth2.NoContext, tok)
	email, err := client.Get(oauthGoogleURLAPI)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()

	data, _ := ioutil.ReadAll(email.Body)
	c.JSON(http.StatusOK, string(data))
}
