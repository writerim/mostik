package restful

import (
	"encoding/json"
	"entity"
	"github.com/hoisie/web"
	"io/ioutil"
	"repositories"
)

const (
	USER_NOT_FOUND = "user not found"
)

type LoginPost struct {
	Username string `json:"username" example:"admin"` // логин
	Password string `json:"password" example:"admin"` // пароль
}

type LoginError struct {
	Error string `json:"error"`
}

type LoginSuccess struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

// @Tags Сессии
// @Summary Авторизация клиента
// @ID auth
// @Accept  json
// @Produce  json
// @Param json body LoginPost true " "
// @Success 200 {object} LoginSuccess
// @Failure 500 {object} LoginError
// @Router /api/auth [post]
func handler_login(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	body, err_body := ioutil.ReadAll(ctx.Request.Body)

	if err_body != nil {
		ctx.ResponseWriter.WriteHeader(500)
		return toJSON(LoginError{Error: err_body.Error()})
	}

	req := LoginPost{}

	err := json.Unmarshal(body, &req)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(500)
		return toJSON(LoginError{Error: err.Error()})
	}

	ur := repositories.NewMysqlUserRepository(db)

	user, err := entity.GetByLoginPasswordAny(ur, req.Username, req.Password)

	if err != nil {
		ctx.ResponseWriter.WriteHeader(500)
		return toJSON(LoginError{Error: err.Error()})
	}

	if user.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return toJSON(LoginError{Error: USER_NOT_FOUND})
	}

	token, err := AddSession(user.Id)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(500)
		return toJSON(LoginError{Error: err.Error()})
	}

	return toJSON(LoginSuccess{Token: token, Id: user.Id})
}

/*

// LoginPost godoc
// @Tags Сессии
// @Security BasicAuth
// @ID Информация о себе. auth
// @Accept  json
// @Produce  json
// @Success 200 {object} LoginGet
// @Failure 401 {object} HttpErr "Не авторизован"
// @Router /api/auth [get]
func handler_get_info_by_token(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, user_ := is_auth(ctx)

	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	out_str_g, _ := uuid.NewUUID()
	out_str := out_str_g.String()

	for ident, _ := range sessions.Items() {
		id_user, _ := sessions.Get(ident)
		if id_user.(int) == user_.Id {
			sessions.Delete(ident)
		}
	}

	sessions.Add(out_str, user_.Id, cache.NoExpiration)

	return toJSON(LoginGet{
		Token:    out_str,
		Username: user_.Login,
		Email:    user_.Email,
		Access:   user_.GetAccess(),
	})

}

// @Security BasicAuth
// @Tags Сессии
// @ID Выход из системы. auth
// @Accept  json
// @Produce  json
// @Success 200 {object} ApiStatus
// @Failure 401 {object} HttpErr "Не авторизован"
// @Router /api/auth [delete]
func handler_logout(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, _ := is_auth(ctx)

	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	sessions.Delete(ctx.Request.Header["Token"][0])

	return `{"status": "nologged"}`

}

*/
