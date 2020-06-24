package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// GenerateAccessToken : アクセストークンの生成
func GenerateAccessToken(userID string, isAdmin bool) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["is_admin"] = isAdmin
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
	claims["iat"] = time.Now().Unix()
	claims["exp"] = exp
	refreshToken, err := token.SignedString(signBytes)
	if err != nil {
		panic(err)
	}
	return refreshToken, strconv.FormatInt(exp, 10)
}

// GetTokenClaimsFromToken : トークンからトークンClaimを取得
func GetTokenClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signBytes, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// GetTokenClaimsFromRequest : コンテキストからトークンClaimを取得
func GetTokenClaimsFromRequest(c *gin.Context) (jwt.MapClaims, error) {
	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return signBytes, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// TokenRefresh :　リフレッシュトークンを受取り、新しくトークンの発行を行う
func TokenRefresh(refreshToken string, isAdmin bool) (string, string, string, error) {
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
	newAccessToken := GenerateAccessToken(claims["sub"].(string), isAdmin)
	newRefreshToken, exp := GenerateRefreshToken(claims["sub"].(string))
	return newAccessToken, newRefreshToken, exp, nil
}
