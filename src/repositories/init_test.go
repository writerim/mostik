package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

func _set_connect(t *testing.T) *gorm.DB {

	txdb.Register("mostik_test",
		"mysql",
		`newuser:111@/mostik_test?charset=utf8&parseTime=True&loc=Local`,
	)

	s, err := sql.Open("mostik_test", "m")
	db, err := gorm.Open("mysql", s)

	if err != nil {
		t.Fatal("Connect to database: ", err.Error())
		return nil
	}
	return db
}
