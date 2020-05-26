package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/infrastructure/persistence/datastore"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/handler"
)

// InitRouter :　ルーティング
func InitRouter(dbConn *gorm.DB, redisConn redis.Conn) *gin.Engine {
	// DI
	userDatastore := datastore.NewUserDatastore(dbConn)
	roomDatastore := datastore.NewRoomDatastore(dbConn)
	roomBalacklistDatastore := datastore.NewRoomBlacklistDatastore(dbConn)
	roomReservationDatastore := datastore.NewRoomReservationDatastore(dbConn)
	chatPostDatastore := datastore.NewChatPostDatastore(dbConn)

	userUsecase := usecase.NewUserUsecase(userDatastore)
	roomUsecase := usecase.NewRoomUsecase(roomDatastore)
	roomBlacklistUsecase := usecase.NewRoomBlacklistUsecase(roomBalacklistDatastore)
	roomReservationUsecase := usecase.NewRoomReservationUsecase(roomReservationDatastore)
	chatPostUsecase := usecase.NewChatPostUsecase(chatPostDatastore)

	roomHandler := handler.NewRoomHandler(userUsecase, roomUsecase, roomBlacklistUsecase, roomReservationUsecase)
	chatPostHandler := handler.NewChatPostHandler(chatPostUsecase, redisConn)
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
	room := v1.Group("/room")
	{
		room.GET("/list", roomHandler.GetRoom)
		room.POST("/:id/chat", chatPostHandler.CreateChatPost)
	}

	return r
}
