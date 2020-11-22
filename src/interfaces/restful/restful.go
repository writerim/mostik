package restful

import (
	"fmt"
	"github.com/hoisie/web"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var sessions *cache.Cache

func init() {
	sessions = cache.New(cache.NoExpiration, cache.NoExpiration)
}

func Init(port int) *web.Server {
	server := web.NewServer()

	init_router(server)

	go server.Run(fmt.Sprintf(`:%d`, port))

	logrus.Infof("Запуск на :%d", port)
	return server
}
