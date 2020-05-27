package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

// GenerateToken : トークンの生成
func GenerateToken(userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userID
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signBytes)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

// TokenAuth :
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(signBytes)
			return b, nil
		})
		if err != nil && !token.Valid {
			c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
			c.Abort()
		}
	}
}
