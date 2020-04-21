package router

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/infrastructure/persistence/datastore"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
)

// InitRouter :　ルーティング
func InitRouter(conn *gorm.DB) *gin.Engine {
	// DI
	userStore := datastore.NewUserDatastore(conn)
	userUsecase := usecase.NewUserUsecase(userStore)
	googleAuthHandler := auth.NewGoogleOAuthHandler(userUsecase)

	store := dao.NewRedisStore()
	//f, err := os.Create("./config/log/access.log")
	//if err != nil {
	//	panic(err.Error())
	//}
	//gin.DefaultWriter = io.MultiWriter(f)

	corsConf := cors.DefaultConfig()

	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders("authorization")

	r := gin.Default()
	// add middleware
	r.Use(cors.New(corsConf))
	r.Use(sessions.Sessions("session", store))

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	google := auth.Group("/google")
	{
		google.GET("/login", googleAuthHandler.Login)
		google.GET("/callback", googleAuthHandler.Callback)
	}
	v1.Use(sessionCheck())
	{
		v1.GET("/", SigninFormRoute)
	}

	return r
}

func SigninFormRoute(g *gin.Context) {
	g.String(200, "hello")
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		retrievedState := session.Get("state")

		// セッションがない場合、ログインフォームをだす
		if retrievedState == nil {
			log.Println("ログインしていません")
			c.String(200, "ログインしてません")
			c.Abort() // これがないと続けて処理されてしまう
		} else {
			c.String(200, "ログインしてます")
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}
