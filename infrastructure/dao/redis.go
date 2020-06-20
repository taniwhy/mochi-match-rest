package dao

import (
	"github.com/gomodule/redigo/redis"
	"github.com/taniwhy/mochi-match-rest/config"

	// Redisドライバ
	redisDriver "github.com/gin-contrib/sessions/redis"
)

// NewRedisConn : Redisコネクションの確立
func NewRedisConn() redis.Conn {
	_, network, addr, _, _ := config.GetRedisConf()
	conn, err := redis.Dial(network, addr)
	if err != nil {
		panic(err)
	}
	return conn
}

// NewRedisStore : Redisストアの生成
func NewRedisStore() redisDriver.Store {
	size, network, addr, _, key := config.GetRedisConf()
	store, err := redisDriver.NewStore(size, network, addr, "", []byte(key))
	if err != nil {
		panic(err.Error())
	}
	return store
}
