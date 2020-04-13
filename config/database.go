package config

import (
	"fmt"

	"github.com/gin-contrib/sessions/redis"
	"github.com/jinzhu/gorm"

	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	conn *gorm.DB
	err  error
)

//DBConn :　データベースコネクションの確立
func DBConn() *gorm.DB {
	c := Config
	HOST := c.Database.Host
	PORT := c.Database.Port
	USER := c.Database.User
	PASS := c.Database.Pass
	DBNAME := c.Database.DBName

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASS, DBNAME)
	conn, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

// NewRedisStore : TODO
func NewRedisStore() redis.Store {
	c := Config
	SIZE := c.Redis.Size
	NETWORK := c.Redis.Network
	ADDR := c.Redis.Addr
	PASS := c.Redis.Pass
	KEY := c.Redis.Key

	store, err := redis.NewStore(SIZE, NETWORK, ADDR, PASS, []byte(KEY))
	if err != nil {
		panic(err.Error())
	}
	return store
}
