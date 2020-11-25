// description entity
// bissnues logic

package entity

import (
	"crypto/sha1"
	"encoding/hex"
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
	// Авторизация
	Auth(string, string) (User, error)
}

type UserRepository interface {
	GetById(id int) (User, error)                                // Получение по id
	GetByLoginPasswordAny(login, passworrd string) (User, error) // Получение по логину и паролю
	Save(User) (User, error)                                     // Сохранение
	GetCountAll() (int, error)                                   // Общее кол-во людей в системе
}

// Поиск по всем пользотелям. Нужно для авторизации
func GetByLoginPasswordAny(r UserRepository, login, password string) (User, error) {
	return r.GetByLoginPasswordAny(login, CompressPass(password))
}

func CompressPass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
