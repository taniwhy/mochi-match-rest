package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/router"
)

func init() {
	config.InitConf()
	fmt.Println("1", config.Config.Database.DBName)
	fmt.Println("2", config.Config.GoogleOAuth.ClientID)
}

func main() {
	conn := config.DBConn()
	defer conn.Close()

	routers := router.InitRouter(conn)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routers,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
