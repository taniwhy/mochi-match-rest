package dao

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	//
	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/config"

	//
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// postgres driver
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
	if err != nil {
		panic(err.Error())
	}
	driver, err := postgres.WithInstance(conn.DB(), &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations", // マイグレーションファイルがあるディレクトリの指定
		"mysql",
		driver,
	)
	m.Steps(2)
	return conn
}
