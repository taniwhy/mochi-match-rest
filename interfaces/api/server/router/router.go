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

	userUsecase := usecase.NewUserUsecase(userDatastore, userDetailDatastore, userService, favorateGameDatastore)
	roomUsecase := usecase.NewRoomUsecase(roomDatastore)
	roomBlacklistUsecase := usecase.NewRoomBlacklistUsecase(roomBalacklistDatastore)
	roomReservationUsecase := usecase.NewRoomReservationUsecase(roomReservationDatastore)
	chatPostUsecase := usecase.NewChatPostUsecase(chatPostDatastore)
	gameTitleUsecase := usecase.NewGameTitleUsecase(gameTitleDatastore)
	favorateGameUsecase := usecase.NewFavoriteGameUsecase(favorateGameDatastore)

	userHandler := handler.NewUserHandler(userUsecase, favorateGameUsecase)
	roomHandler := handler.NewRoomHandler(userUsecase, roomUsecase, roomBlacklistUsecase, roomReservationUsecase)
	chatPostHandler := handler.NewChatPostHandler(chatPostUsecase, redisConn)
	gameTitleHandler := handler.NewGameTitleHandler(gameTitleUsecase)
	authHandler := handler.NewAuthHandler()
	googleAuthHandler := handler.NewGoogleOAuthHandler(userUsecase, userService)

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
	games := users.Group("/:id/favorate/games")
	{
		games.GET("")
		games.POST("")
	}
	room := v1.Group("/rooms")
	room.Use(auth.TokenAuth())
	{
		room.GET("")
		room.GET("/:id/messages", chatPostHandler.GetChatPostByRoomID)
		room.POST("", roomHandler.CreateRoom)
		room.POST("/:id/messages", chatPostHandler.CreateChatPost)
		room.GET("/:id/report")
		room.POST("/:id/report")
	}
	gamelist := v1.Group("/gamelist")
	{
		gamelist.GET("", gameTitleHandler.GetAllGameTitle)
		gamelist.POST("", gameTitleHandler.CreateGameTitle)
		gamelist.PUT("/:id", gameTitleHandler.UpdateGameTitle)
		gamelist.DELETE("/:id", gameTitleHandler.DeleteGameTitle)
	}
	report := v1.Group("/report")
	{
		report.GET("")
	}

	return r
}
