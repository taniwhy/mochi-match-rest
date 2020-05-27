package dao

import (
	// postgres drive
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/taniwhy/mochi-match-rest/config"
)

// NewRedisConn : TODO
func NewRedisConn() redis.Conn {
	_, network, addr, _, _ := config.GetRedisConf()
	conn, err := redis.Dial(network, addr)
	if err != nil {
		panic(err)
	}
	return conn
}

// NewRedisStore : TODO
func NewRedisStore() redisStore.Store {
	size, network, addr, pass, key := config.GetRedisConf()
	store, err := redisStore.NewStore(size, network, addr, pass, []byte(key))
	if err != nil {
		panic(err.Error())
	}
	return store
}
