package repositories

import (
	"entity"
	"fmt"
	"testing"
)

func TestGetById(t *testing.T) {

	db := _set_connect(t)
	defer db.Close()
	db.LogMode(false)

	ur := NewMysqlUserRepository(db)

	u := entity.User{
		Login:          "test",
		Password:       "123",
		PersonalAreaId: 1,
		Email:          "sdsdsd@sdsd",
		ApiToken:       "dsdsds",
		Properties:     "{}",
		Location:       "sdsds",
		RoleId:         232,
	}

	us, err := ur.Save(u)
	fmt.Println(us, err)

}
