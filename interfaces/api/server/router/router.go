package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	"github.com/taniwhy/mochi-match-rest/application/usecase"
	"github.com/taniwhy/mochi-match-rest/domain/service"
	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/infrastructure/persistence/datastore"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/handler"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/auth"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/middleware/cors"
)

// InitRouter :　ルーティングセットアップ
func InitRouter(dbConn *gorm.DB, redisConn redis.Conn) *gin.Engine {
	// DI
	userDatastore := datastore.NewUserDatastore(dbConn)
	userDetailDatastore := datastore.NewUserDetailDatastore(dbConn)
	roomDatastore := datastore.NewRoomDatastore(dbConn)
	roomBalacklistDatastore := datastore.NewRoomBlacklistDatastore(dbConn)
	entryHistoryDatastore := datastore.NewEntryHistoryDatastore(dbConn)
	chatPostDatastore := datastore.NewChatPostDatastore(dbConn)
	gameListDatastore := datastore.NewGameListDatastore(dbConn)
	gameHardDatastore := datastore.NewGameHardDatastore(dbConn)
	favorateGameDatastore := datastore.NewFavoriteGameDatastore(dbConn)

	userService := service.NewUserService(userDatastore)
	roomService := service.NewRoomService(roomDatastore)
	entryHistoryService := service.NewEntryHistoryService(entryHistoryDatastore)

	userUsecase := usecase.NewUserUsecase(userDatastore, userDetailDatastore, userService, favorateGameDatastore)
	roomUsecase := usecase.NewRoomUsecase(roomDatastore, entryHistoryDatastore, roomService, entryHistoryService)
	roomBlacklistUsecase := usecase.NewRoomBlacklistUsecase(roomBalacklistDatastore, roomService)
	chatPostUsecase := usecase.NewChatPostUsecase(chatPostDatastore, redisConn)
	gameListUsecase := usecase.NewGameListUsecase(gameListDatastore)
	gameHardUsecase := usecase.NewGameHardUsecase(gameHardDatastore)
	googleAuthUsecase := usecase.NewGoogleOAuthUsecase(userService)

	userHandler := handler.NewUserHandler(userUsecase)
	roomHandler := handler.NewRoomHandler(userUsecase, roomUsecase)
	roomBlacklistHandler := handler.NewRoomBlacklistHandler(roomBlacklistUsecase)
	chatPostHandler := handler.NewChatPostHandler(chatPostUsecase)
	gameListHandler := handler.NewGameListHandler(gameListUsecase)
	gameHardHandler := handler.NewGameHardHandler(gameHardUsecase)
	googleAuthHandler := handler.NewGoogleOAuthHandler(googleAuthUsecase, userUsecase, userService)
	authHandler := handler.NewAuthHandler(userService)

	r := gin.Default()
	r.Use(cors.Write())

	store := dao.NewRedisStore()
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.Write())
	r.LoadHTMLGlob("public/*")

	v1 := r.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"content-type": "text/html"})
	})
	authHandle := v1.Group("/auth")
	{
		authHandle.GET("/get", authHandler.GetToken)
		authHandle.POST("/refresh", authHandler.Refresh)
	}
	google := authHandle.Group("/google")
	{
		google.GET("/login", googleAuthHandler.Login)
		google.GET("/callback", googleAuthHandler.Callback)
	}
	check := v1.Group("/check")
	{
		check.GET("/entry", roomHandler.CheckEntry)
	}
	hot := v1.Group("/hot")
	{
		hot.GET("/games")
	}
	users := v1.Group("/users")
	users.Use(auth.TokenAuth())
	{
		users.GET("", userHandler.GetMe)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("", userHandler.Update)
		users.DELETE("", userHandler.Delete)
	}
	room := v1.Group("/rooms")
	room.GET("", roomHandler.GetList)
	room.Use(auth.TokenAuth())
	{
		room.GET("/:id", roomHandler.GetByID)
		room.POST("", roomHandler.Create)
		room.PUT("/:id", roomHandler.Update)
		room.DELETE("/:id", roomHandler.Delete)
		room.POST("/:id/join", roomHandler.Join)
		room.DELETE("/:id/leave", roomHandler.Leave)
	}
	messages := room.Group("/:id/messages")
	{
		messages.GET("", chatPostHandler.GetChatPostByRoomID)
		messages.POST("", chatPostHandler.CreateChatPost)
	}
	report := room.Group("/:id/report")
	{
		report.POST("")
	}
	blacklist := room.Group("/:id/blacklist")
	{
		blacklist.GET("", roomBlacklistHandler.GetByRoomID)
		blacklist.POST("", roomBlacklistHandler.Create)
		blacklist.DELETE("", roomBlacklistHandler.Delete)
	}
	gamelist := v1.Group("/gamelist")
	{
		gamelist.GET("", gameListHandler.GetAll)
	}
	gamelist.Use(auth.AdminAuth())
	{
		gamelist.POST("", gameListHandler.Create)
		gamelist.PUT("/:id", gameListHandler.Update)
		gamelist.DELETE("/", gameListHandler.Delete)
	}
	gamehard := v1.Group("/gamehard")
	{
		gamehard.GET("", gameHardHandler.GetAll)
	}
	gamehard.Use(auth.AdminAuth())
	{
		gamehard.POST("", gameHardHandler.Create)
		gamehard.PUT("/:id", gameHardHandler.Update)
		gamehard.DELETE("/", gameHardHandler.Delete)
	}
	return r
}
