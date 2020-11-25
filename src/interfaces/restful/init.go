package restful

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hoisie/web"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	NOT_INIT_SESSION = "session not init"
)

// Определяем дефолтную запись в сессию
var AddSession = func(user_id int) (string, error) {
	return "", errors.New(NOT_INIT_SESSION)
}

var db *gorm.DB

func Init(port int, db *gorm.DB) *web.Server {
	server := web.NewServer()

	init_router(server)

	go server.Run(fmt.Sprintf(`:%d`, port))

	logrus.Infof("Запуск на :%d", port)
	return server
}

func toJSON(out interface{}) string {
	e, err := json.Marshal(out)
	if err != nil {
		logrus.Fatal("toJSON ", err.Error())
		return `{}`
	}
	return string(e)
}
