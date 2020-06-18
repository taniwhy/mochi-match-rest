package main

import (
	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/router"

	// logging driver
	"github.com/taniwhy/mochi-match-rest/logging"
)

func main() {
	dbConn := dao.NewDatabase()
	defer dbConn.Close()

	redisConn := dao.NewRedisConn()
	defer redisConn.Close()

	dbConn.LogMode(true)
	dbConn.SetLogger(&logging.GormLogger{})

	routers := router.InitRouter(dbConn, redisConn)

	routers.Run(":8000")
}
