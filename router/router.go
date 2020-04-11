package router

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-contrib/cors"
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

	corsConf := cors.DefaultConfig()

	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders("authorization")

	r := gin.Default()
	// add middleware
	r.Use(cors.New(corsConf))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}
