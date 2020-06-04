package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/infrastructure/persistence/datastore"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/auth"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/handler"
)

// InitRouter :　ルーティング
func InitRouter(dbConn *gorm.DB, redisConn redis.Conn) *gin.Engine {
	dbConn.LogMode(true)
	// DI
	userDatastore := datastore.NewUserDatastore(dbConn)
	userDetailDatastore := datastore.NewUserDetailDatastore(dbConn)
	roomDatastore := datastore.NewRoomDatastore(dbConn)
	roomBalacklistDatastore := datastore.NewRoomBlacklistDatastore(dbConn)
	roomReservationDatastore := datastore.NewRoomReservationDatastore(dbConn)
	chatPostDatastore := datastore.NewChatPostDatastore(dbConn)
	gameTitleDatastore := datastore.NewGameTitleDatastore(dbConn)
	favorateGameDatastore := datastore.NewFavoriteGameDatastore(dbConn)

	userService := service.NewUserService(userDatastore)
	roomService := service.NewRoomService(roomDatastore)

	userUsecase := usecase.NewUserUsecase(userDatastore, userDetailDatastore, userService, favorateGameDatastore)
	roomUsecase := usecase.NewRoomUsecase(roomDatastore, roomService)
	roomBlacklistUsecase := usecase.NewRoomBlacklistUsecase(roomBalacklistDatastore)
	roomReservationUsecase := usecase.NewRoomReservationUsecase(roomReservationDatastore)
	chatPostUsecase := usecase.NewChatPostUsecase(chatPostDatastore)
	gameTitleUsecase := usecase.NewGameTitleUsecase(gameTitleDatastore)

	userHandler := handler.NewUserHandler(userUsecase)
	roomHandler := handler.NewRoomHandler(userUsecase, roomUsecase, roomReservationUsecase)
	roomBlacklistHandler := handler.NewRoomBlacklistHandler(userUsecase, roomUsecase, roomBlacklistUsecase)
	chatPostHandler := handler.NewChatPostHandler(chatPostUsecase, redisConn)
	gameTitleHandler := handler.NewGameTitleHandler(gameTitleUsecase)
	googleAuthHandler := handler.NewGoogleOAuthHandler(userUsecase, userService)

	authHandler := handler.NewAuthHandler()

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
	authHandle := v1.Group("/auth")
	{
		authHandle.POST("/refresh", authHandler.Refresh)
	}
	google := authHandle.Group("/google")
	{
		google.GET("/login", googleAuthHandler.Login)
		google.GET("/callback", googleAuthHandler.Callback)
	}
	users := v1.Group("/users")
	users.POST("", userHandler.Create)
	users.Use(auth.TokenAuth())
	{
		users.GET("", userHandler.GetMe)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}
	room := v1.Group("/rooms")
	room.Use(auth.TokenAuth())
	{
		room.GET("", roomHandler.GetList)
		room.GET("/:id", roomHandler.GetByID)
		room.POST("", roomHandler.Create)
		room.PUT("/:id", roomHandler.Update)
		room.GET("/:id/messages", chatPostHandler.GetChatPostByRoomID)
		room.POST("/:id/messages", chatPostHandler.CreateChatPost)
		room.GET("/:id/report")
		room.POST("/:id/report")
		room.GET("/:id/blacklist", roomBlacklistHandler.GetByID)
		room.POST("/:id/blacklist", roomBlacklistHandler.Create)
		room.DELETE("/:id/blacklist", roomBlacklistHandler.Delete)
	}
	gamelist := v1.Group("/gamelist")
	{
		gamelist.GET("", gameTitleHandler.GetAllGameTitle)
		gamelist.POST("", gameTitleHandler.CreateGameTitle)
		gamelist.PUT("/:id", gameTitleHandler.UpdateGameTitle)
		gamelist.DELETE("/:id", gameTitleHandler.DeleteGameTitle)
	}
	return r
}
