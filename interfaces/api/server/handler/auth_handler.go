package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
)

// AuthHandler : todo
type AuthHandler interface {
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
