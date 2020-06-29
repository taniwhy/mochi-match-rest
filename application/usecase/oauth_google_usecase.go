//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="

// IGoogleOAuthUsecase : インターフェース
type IGoogleOAuthUsecase interface {
	Login(c *gin.Context) (string, error)
	Callback(c *gin.Context) (bool, *models.GoogleUser, error)
}

type googleOAuthUsecase struct {
	oauthConf   *oauth2.Config
	userService service.IUserService
}

// NewGoogleOAuthUsecase : GoogleOAuthユースケースの生成
func NewGoogleOAuthUsecase(uS service.IUserService) IGoogleOAuthUsecase {
	return &googleOAuthUsecase{
		oauthConf:   config.GetOAuthClientConf(),
		userService: uS,
	}
}

func (u *googleOAuthUsecase) Login(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	option := sessions.Options{Path: "/", Domain: "", MaxAge: 300, Secure: false, HttpOnly: true}
	session.Options(option)

	sessionID := uuid.UuID()
	session.Set("state", sessionID)
	session.Save()

	url := u.oauthConf.AuthCodeURL(sessionID)
	return url, nil
}

func (u *googleOAuthUsecase) Callback(c *gin.Context) (bool, *models.GoogleUser, error) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		return false, nil, errors.ErrInvalidSessionState{State: retrievedState}
	}
	token, err := u.oauthConf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		return false, nil, errors.ErrGoogleOAuthTokenExchange{}
	}

	if token.Valid() == false {
		return false, nil, errors.ErrInvalidGoogleOAuthToken{}
	}

	client := u.oauthConf.Client(oauth2.NoContext, token)
	response, err := client.Get(oauthGoogleURLAPI)
	if err != nil {
		return false, nil, errors.ErrGoogleAPIRequest{}
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, nil, errors.ErrReadGoogleAPIResponse{}
	}
	googleUser := models.GoogleUser{}
	err = json.Unmarshal(data, &googleUser)
	if err != nil {
		return false, nil, errors.ErrUnmarshalJSON{}
	}

	ok, err := u.userService.IsExist("google", googleUser.ID)
	if err != nil {
		return false, nil, err
	}
	if ok {
		return true, &googleUser, nil
	}
	return false, &googleUser, nil
}
