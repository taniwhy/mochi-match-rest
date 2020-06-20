package dao

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/config"

	// マイグレーションドライバ
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	migration()
	return conn
}

func migration() {
	driver, _ := postgres.WithInstance(conn.DB(), &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations", // マイグレーションファイルがあるディレクトリの指定
		"postgres", driver,
	)
	m.Steps(4)
}
