package restful

import (
	"encoding/json"
	"entity"
	"repositories"
	"testing"
)

func TestAuthUserUndefined(t *testing.T) {

	db = _set_connect(t)
	defer db.Close()
	db.LogMode(false)

	res := auth(LoginPost{
		Username: "admin",
		Password: "admin",
	})

	r := LoginError{}
	json.Unmarshal(res, &r)

	if r.Error != USER_NOT_FOUND {
		t.Fatal("Не верный ответ")
	}
}

func TestAuthUserSuccess(t *testing.T) {

	db = _set_connect(t)
	defer db.Close()
	db.LogMode(false)

	rr := repositories.NewMysqlUserRepository(db)
	rr.Save(entity.User{
		Login:          "admin",
		Password:       "admin",
		PersonalAreaId: 1,
	})

	res := auth(LoginPost{
		Username: "admin",
		Password: "admin",
	})

	r := LoginSuccess{}
	json.Unmarshal(res, &r)

	if r.Token == "" {
		t.Fatal("Не верный ответ")
	}
}
