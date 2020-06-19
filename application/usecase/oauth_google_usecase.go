package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"golang.org/x/oauth2"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="

// GoogleOAuthUsecase :
type GoogleOAuthUsecase interface {
	Login(c *gin.Context) (string, error)
	Callback(c *gin.Context) (bool, *models.GoogleUser, error)
}

type googleOAuthUsecase struct {
	oauthConf   *oauth2.Config
	userService service.IUserService
}

// NewGoogleOAuthUsecase :
func NewGoogleOAuthUsecase(uS service.IUserService) GoogleOAuthUsecase {
	return &googleOAuthUsecase{
		oauthConf:   config.ConfigureOAuthClient(),
		userService: uS,
	}
}

func (gA *googleOAuthUsecase) Login(c *gin.Context) (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", errors.ErrGenerateID{}
	}
	sessionID := u.String()
	session := sessions.Default(c)
	session.Set("state", sessionID)
	session.Set("name", "client")
	fmt.Println(sessionID)
	if err := session.Save(); err != nil {
		return "", errors.ErrSessionSave{}
	}

	url := gA.oauthConf.AuthCodeURL(sessionID)
	return url, nil
}

func (gA *googleOAuthUsecase) Callback(c *gin.Context) (bool, *models.GoogleUser, error) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	fmt.Println(retrievedState, c.Query("state"))
	fmt.Println(session.Get("name"))
	if retrievedState != c.Query("state") {
		return false, nil, errors.ErrInvalidSessionState{State: retrievedState}
	}

	tok, err := gA.oauthConf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		return false, nil, errors.ErrGoogleOAuthTokenExchange{}
	}

	if tok.Valid() == false {
		return false, nil, errors.ErrInvalidGoogleOAuthToken{}
	}

	client := gA.oauthConf.Client(oauth2.NoContext, tok)
	email, err := client.Get(oauthGoogleURLAPI)
	if err != nil {
		return false, nil, errors.ErrGoogleAPIRequest{}
	}
	defer email.Body.Close()

	data, err := ioutil.ReadAll(email.Body)
	if err != nil {
		return false, nil, errors.ErrReadGoogleAPIResponse{}
	}

	gU := models.GoogleUser{}
	err = json.Unmarshal(data, &gU)
	if err != nil {
		return false, nil, errors.ErrUnmarshalJSON{}
	}

	ok, err := gA.userService.IsExist("google", gU.ID)
	if err != nil {
		return false, nil, err
	}
	if ok {
		return true, &gU, nil
	}
	return false, &gU, nil
}
