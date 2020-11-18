package restful

import (
    "github.com/hoisie/web"
    "use_cases"
    "repositories"
)


func user_get_all(ctx *web.Context) string{

	repo := repositories.NewMysqlUserRepository(db)
	uc := use_cases.NewUserUseCase(repo)

	t,_ := uc.MoveTo(repo,10,15)

	fmt.Println(t)

	return ""
}