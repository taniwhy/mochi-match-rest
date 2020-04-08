package main

import (
	"net/http"
	"time"

	"github.com/taniwhy/mochi-match-rest/config"
	"github.com/taniwhy/mochi-match-rest/router"
)

func main() {
	conn := config.NewDB()
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
