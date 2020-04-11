package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	conn *gorm.DB
	err  error
)

//DBConn :　データベースコネクションの確立とマイグレーションの実行
func DBConn() *gorm.DB {
	InitConf()
	c := Config

	HOST := c.Database.Host
	PORT := c.Database.Port
	USER := c.Database.User
	PASS := c.Database.Password
	DBNAME := c.Database.DBName

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASS, DBNAME)
	conn, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
