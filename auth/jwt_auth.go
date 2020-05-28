package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/taniwhy/css-study-game-WebAPI/domain/model"
)

var (
	signBytes []byte
	err       error
)

func init() {
	signBytes, err = ioutil.ReadFile("./config/key/authorize.rsa")
	if err != nil {
		fmt.Println(err)
	}
}

// GenerateAccessToken : トークンの生成
func GenerateAccessToken(userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := token.SignedString(signBytes)
	if err != nil {
		panic(err)
	}
	return accessToken
}

// GenerateRefreshToken : トークンの生成
func GenerateRefreshToken(userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	refreshToken, err := token.SignedString(signBytes)
	if err != nil {
		fmt.Println(err)
	}
	return refreshToken
}

// TokenAuth :
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(signBytes)
			return b, nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
			c.Abort()
		}
	}
}

// TokenRefresh :
func TokenRefresh(refreshToken string) (string, string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signBytes, nil
	})
	if err != nil {
		return "", "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	newAccessToken := GenerateAccessToken(claims["sub"].(string))
	newRefreshToken := GenerateRefreshToken(claims["sub"].(string))
	return newAccessToken, newRefreshToken, nil
}
