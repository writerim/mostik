package restful

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/web"
	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"repositories"
	"testing"
	"time"
)

var sessions *cache.Cache

func init() {
	sessions = cache.New(cache.NoExpiration, cache.NoExpiration)

	AddSession = func(user_id int) (string, error) {
		live := time.Duration(10) * time.Second
		sessions.Add("TEST", user_id, live)
		return "TEST", nil
	}
}

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
	repositories.Init(db)
	return db
}

func auth(login LoginPost) []byte {
	my_url := func(w http.ResponseWriter, r *http.Request) {
		ctx := &web.Context{
			Request:        r,
			Params:         make(map[string]string),
			ResponseWriter: w,
		}
		io.WriteString(w,
			handler_login(ctx),
		)
	}

	auth_str, _ := json.Marshal(login)

	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(auth_str))

	w := httptest.NewRecorder()
	my_url(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	return body
}
