package dao

import (
	// postgres drive
	"github.com/gin-contrib/sessions/redis"
	"github.com/taniwhy/mochi-match-rest/config"
)

// NewRedisStore : TODO
func NewRedisStore() redis.Store {
	size, network, addr, pass, key := config.GetRedisConf()
	store, err := redis.NewStore(size, network, addr, pass, []byte(key))
	if err != nil {
		panic(err.Error())
	}
	return store
}
