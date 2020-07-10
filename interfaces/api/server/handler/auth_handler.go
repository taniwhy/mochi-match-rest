package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taniwhy/mochi-match-rest/domain/models/input"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// IAuthHandler : インターフェース
type IAuthHandler interface {
	GetToken(c *gin.Context)
	Refresh(c *gin.Context)
}

type authHandler struct {
	userService service.IUserService
}

// NewAuthHandler : 認証ハンドラの生成
func NewAuthHandler() IAuthHandler {
	return &authHandler{}
}

func (aH *authHandler) GetToken(c *gin.Context) {
	token, err := c.Cookie("token")
	fmt.Println("token", token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	expStr, err := c.Cookie("token_exp")
	fmt.Println("expStr", expStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	expAt, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if time.Now().Unix() > expAt {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	claims, err := auth.GetTokenClaimsFromToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	claimsID := claims["sub"].(string)
	fmt.Println("id", claimsID)
	// TODO
	isAdmin := true
	fmt.Println("isAdmin", isAdmin)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	accessToken, refreshToken, exp, err := auth.TokenRefresh(token, isAdmin)
	if err != nil {
		fmt.Println("err", err)
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
	tokenReq := input.TokenReqBody{}
	c.Bind(&tokenReq)
	claims, err := auth.GetTokenClaimsFromToken(tokenReq.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	claimsID := claims["sub"].(string)
	isAdmin, err := aH.userService.IsAdmin(claimsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	accessToken, refreshToken, exp, err := auth.TokenRefresh(tokenReq.RefreshToken, isAdmin)
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
