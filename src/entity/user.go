// description entity
// bissnues logic

package entity

import (
	"encoding/json"
)

func (User) TableName() string {
	return "user"
}

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
	GetAll() ([]User, error)   // Получение всех пользоателей
	GetById(int) (User, error) // Получение по идентификатору
	AccessUser(User) error     // Проверка видимости одним пользователем второго
	GetRoleId() int
	GetPersonalAreaId() int
	GetLogin() string
	GetPassword() string
	GetProperties() string
	SetProperties(prop string) (User, error)
}

type UserRepository interface {
	GetById(id int) (User, error) // Получение по id
	// GetByPersonalArea(personal_area_id int) ([]User, error) // Получние всех в этом ЛК
	// GetByApiToken(tonen string) (User, error)               // Получение по токену
	// GetByLoginByPass(login, passworrd string) (User, error) // Получение по логину и паролю
	// Add(User) (User, error)                                 // Добавление
	// Update(User) (User, error)                              // Редактирование
	// Delete(User) error                                      // Удаление
	Save(User) (User, error)   // Сохранение
	GetCountAll() (int, error) // Общее кол-во людей в системе
}

func (u User) GetRoleId() int {
	return u.RoleId
}

func (u User) GetPersonalAreaId() int {
	return u.PersonalAreaId
}

func (u User) GetLogin() string {
	return u.Login
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetProperties() string {
	return u.Properties
}

func (u User) SetProperties(prop string) (User, error) {

	v := make(map[string]interface{})
	if err := json.Unmarshal([]byte(prop), &v); err != nil {
		return User{}, err
	}

	u.Properties = prop
	return u, nil
}

func (u User) SetRoleId(role_id int) (User, error) {
	u.RoleId = role_id
	return u, nil
}
