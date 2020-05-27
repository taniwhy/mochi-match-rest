package datastore

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestFindAllUser(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	test := NewUserDatastore(db)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{""}))

	_, err = test.FindAllUser()
	if err != nil {
		t.Fatal(err)
	}
}
