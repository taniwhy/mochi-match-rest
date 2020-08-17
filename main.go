package main

import (
	"log"
	"net/http"
	"time"

	"github.com/taniwhy/mochi-match-rest/infrastructure/dao"
	"github.com/taniwhy/mochi-match-rest/interfaces/api/server/router"
	// logging driver
)

func main() {
	dbConn := dao.NewDatabase()
	defer dbConn.Close()

	redisConn := dao.NewRedisConn()
	defer redisConn.Close()

	routers := router.InitRouter(dbConn, redisConn)

	server := &http.Server{
		Addr:           ":8000",
		Handler:        routers,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Serve failed")
		panic(err)
	}
}
