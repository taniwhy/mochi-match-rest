package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var (
	signBytes []byte
)

func init() {
	var err error
	signBytes, err = ioutil.ReadFile("./config/key/authorize.rsa")
	if err != nil {
		panic(err)
	}
}

// GenerateAccessToken : アクセストークンの生成
func GenerateAccessToken(userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := token.SignedString(signBytes)
	if err != nil {
		panic(err)
	}
	return accessToken
}

// GenerateRefreshToken : リフレッシュトークンの生成
func GenerateRefreshToken(userID string) (string, string) {
	exp := time.Now().Add(time.Hour * 72).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = exp
	refreshToken, err := token.SignedString(signBytes)
	if err != nil {
		panic(err)
	}
	return refreshToken, strconv.FormatInt(exp, 10)
}

// GetTokenClaims : コンテキストからトークンClaimを取得
func GetTokenClaims(c *gin.Context) (jwt.MapClaims, error) {
	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return signBytes, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// TokenAuth :　トークンで認証を行う
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(signBytes)
			return b, nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
		}
	}
}

// TokenRefresh :　リフレッシュトークンを受取り、新しくトークンの発行を行う
func TokenRefresh(refreshToken string) (string, string, string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signBytes, nil
	})
	if err != nil {
		return "", "", "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	newAccessToken := GenerateAccessToken(claims["sub"].(string))
	newRefreshToken, exp := GenerateRefreshToken(claims["sub"].(string))
	return newAccessToken, newRefreshToken, exp, nil
}
