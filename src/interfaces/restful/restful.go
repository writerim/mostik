package restful

import (
    "github.com/hoisie/web"
    "fmt"
)

func New(port int) *web.Server{
	server := web.NewServer()
    server.Get("/user_get_all", user_get_all)
    go server.Run(fmt.Sprintf(`:%d` , port))
    return server
}