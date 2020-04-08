package router

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitRouter :　ルーティングのセットアップ
func InitRouter(conn *gorm.DB) *gin.Engine {
	fmt.Print(conn)
	f, err := os.Create("./config/log/access.log")
	if err != nil {
		panic(err.Error())
	}

	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	return r
}
