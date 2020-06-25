package router

import (
	"io"
	"os"

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

	userUsecase := usecase.NewUserUsecase(userDatastore, userDetailDatastore, userService, favorateGameDatastore)
	roomUsecase := usecase.NewRoomUsecase(roomDatastore, entryHistoryDatastore, roomService)
	roomBlacklistUsecase := usecase.NewRoomBlacklistUsecase(roomBalacklistDatastore)
	chatPostUsecase := usecase.NewChatPostUsecase(chatPostDatastore)
	gameListUsecase := usecase.NewGameListUsecase(gameListDatastore)
	gameHardUsecase := usecase.NewGameHardUsecase(gameHardDatastore)
	googleAuthUsecase := usecase.NewGoogleOAuthUsecase(userService)

	userHandler := handler.NewUserHandler(userUsecase)
	roomHandler := handler.NewRoomHandler(userUsecase, roomUsecase)
	roomBlacklistHandler := handler.NewRoomBlacklistHandler(roomBlacklistUsecase)
	chatPostHandler := handler.NewChatPostHandler(chatPostUsecase, redisConn)
	gameListHandler := handler.NewGameListHandler(gameListUsecase)
	gameHardHandler := handler.NewGameHardHandler(gameHardUsecase)
	googleAuthHandler := handler.NewGoogleOAuthHandler(googleAuthUsecase, userUsecase, userService)

	authHandler := handler.NewAuthHandler()

	f, err := os.Create("./config/log/access.log")
	if err != nil {
		panic(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.Use(cors.Write())
	store := dao.NewRedisStore()
	r.Use(sessions.Sessions("session", store))

	v1 := r.Group("/v1")
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
	users := v1.Group("/users")
	users.Use(auth.TokenAuth())
	{
		users.GET("", userHandler.GetMe)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("", userHandler.Update)
		users.DELETE("", userHandler.Delete)
	}
	room := v1.Group("/rooms")
	room.Use(auth.TokenAuth())
	{
		room.GET("", roomHandler.GetList)
		room.GET("/:id", roomHandler.GetByID)
		room.POST("", roomHandler.Create)
		room.PUT("/:id", roomHandler.Update)
		room.DELETE("/:id", roomHandler.Delete)
		room.POST("/:id/join", roomHandler.Join)
		room.DELETE("/:id/leave", roomHandler.Leave)
	}
	messages := room.Group("/:id/messages")
	messages.Use(auth.TokenAuth())
	{
		messages.GET("", chatPostHandler.GetChatPostByRoomID)
		messages.POST("", chatPostHandler.CreateChatPost)
	}
	report := room.Group("/:id/report")
	messages.Use(auth.TokenAuth())
	{
		report.POST("")
	}
	blacklist := room.Group("/:id/blacklist")
	blacklist.Use(auth.TokenAuth())
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
		gamelist.DELETE("/:id", gameListHandler.Delete)
	}
	gamehard := v1.Group("/gamehard")
	{
		gamehard.GET("", gameHardHandler.GetAll)
	}
	gamehard.Use(auth.AdminAuth())
	{
		gamehard.POST("", gameHardHandler.Create)
		gamehard.PUT("/:id", gameHardHandler.Update)
		gamehard.DELETE("/:id", gameHardHandler.Delete)
	}
	return r
}
