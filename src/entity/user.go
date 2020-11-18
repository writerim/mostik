// description entity
// bissnues logic

package entity

type User struct {
	Id             int
	Login          string
	Password       string
	PersonalAreaId int
	Email          string
	ApiToken       string
	Properties     string
	Location       string
	RoleId         int
}

type UserUseCases interface {
	GetByID(id int64) (User, error)
}

type UserRepository interface {
	GetByID(id int64) (User, error)
}

