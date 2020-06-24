package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/config"

	// Postgres ドライバ
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	conn *gorm.DB
	err  error
)

// NewDatabase :　データベースコネクションの確立
func NewDatabase() *gorm.DB {
	dsn := config.GetDatabaseConf()
	if err != nil {
		panic(err.Error())
	}
	conn, err = gorm.Open("postgres", dsn)
	fmt.Println(dsn)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
