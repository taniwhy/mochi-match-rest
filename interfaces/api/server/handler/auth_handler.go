package handler

import (
	"fmt"
	"net/http"

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
func NewAuthHandler(uS service.IUserService) IAuthHandler {
	return &authHandler{
		userService: uS,
	}
}

func (h *authHandler) GetToken(c *gin.Context) {
	type Body struct {
		Token string `json:"token" binding:"required"`
	}
	body := &Body{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, "body error")
		return
	}
	claims, err := auth.GetTokenClaimsFromToken(body.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, "token claims error")
		return
	}
	claimsID := claims["sub"].(string)
	fmt.Println("id", claimsID)
	isAdmin := true
	fmt.Println("isAdmin", isAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	accessToken, refreshToken, exp, err := auth.TokenRefresh(body.Token, isAdmin)
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

func (h *authHandler) Refresh(c *gin.Context) {
	tokenReq := input.TokenReqBody{}
	c.Bind(&tokenReq)
	claims, err := auth.GetTokenClaimsFromToken(tokenReq.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	claimsID := claims["sub"].(string)
	isAdmin, err := h.userService.IsAdmin(claimsID)
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
