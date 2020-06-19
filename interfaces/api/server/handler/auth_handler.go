package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// AuthHandler : todo
type AuthHandler interface {
	GetToken(c *gin.Context)
	Refresh(c *gin.Context)
}

type authHandler struct {
}

type tokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

// NewAuthHandler :
func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (aH *authHandler) GetToken(c *gin.Context) {
	expStr, err := c.Cookie("session")
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	expStr, err = c.Cookie("token_exp")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(expStr)
	expAt, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if time.Now().Unix() > expAt {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	accessToken, refreshToken, exp, err := auth.TokenRefresh(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"expires_in":    exp,
	})
}

func (aH *authHandler) Refresh(c *gin.Context) {
	tokenReq := tokenReqBody{}
	c.Bind(&tokenReq)
	accessToken, refreshToken, exp, err := auth.TokenRefresh(tokenReq.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"expires_in":    exp,
	})
}
