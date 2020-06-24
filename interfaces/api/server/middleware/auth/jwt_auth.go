package auth

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var (
	signBytes []byte
)

func init() {
	signBytes = []byte(os.Getenv("AUTHORIZE_RSA"))
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
