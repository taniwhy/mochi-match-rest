package cors

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Write : レスポンスのヘッダーにCors設定を書き込むミドルウェア
func Write() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch os.Getenv("GO_ENV") {
		case "staging":
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
		case "prod":
			c.Writer.Header().Set("Access-Control-Allow-Origin", "https://mochi-match.work")
		default:
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
		}
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, Access-Control-Allow-Headers, Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
