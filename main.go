package main

import (
	"github.com/taniwhy/mochi-match-rest/config"
)

func main() {
	conn := config.NewDB()
	defer conn.Close()

}
