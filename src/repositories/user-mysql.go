package repositories

import (
	"encoding/json"
	"entity"
	"errors"
	"github.com/jinzhu/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

const (
	EMPTY_LOGIN          = "empty login"
	EMPTY_PASSWORD       = "empty password"
	INVALID_PERSONALAREA = "invalid personal area"
	INVALID_PROPERTIES   = "invalid properties"
)

type UserRepository interface {
	GetById(id int) (entity.User, error)
	Save(user entity.User) (entity.User, error)
	GetCountAll() (int, error)
	GetByLoginPasswordAny(login, password string)
}

func NewMysqlUserRepository(db *gorm.DB) entity.UserRepository {
	return &mysqlUserRepo{
		DB: db,
	}
}

func (m *mysqlUserRepo) GetById(id int) (entity.User, error) {
	// select parse return
	return entity.User{}, nil
}

func (m *mysqlUserRepo) GetByLoginPasswordAny(login, password string) (entity.User, error) {
	record := entity.User{}
	m.DB.Where("login = ? and password = ?", login, password).First(&record)
	return record, nil
}

func (m *mysqlUserRepo) GetCountAll() (int, error) {

	records := []entity.User{}
	count := 0
	m.DB.Find(&records).
		Count(&count)

	return count, nil
}

func (m *mysqlUserRepo) Save(user entity.User) (entity.User, error) {

	if err := m.validate(&user); err != nil {
		return entity.User{}, err
	}

	m.DB.Create(&user)
	return user, nil
}

// Валидация данных
func (m *mysqlUserRepo) validate(user *entity.User) error {
	if user.Login == "" {
		return errors.New(EMPTY_LOGIN)
	}
	if user.Password == "" {
		return errors.New(EMPTY_PASSWORD)
	}
	if user.PersonalAreaId == 0 {
		return errors.New(INVALID_PERSONALAREA)
	}
	if user.Properties == "" {
		user.Properties = `{}`
	}

	v := make(map[string]interface{})
	if err := json.Unmarshal([]byte(user.Properties), &v); err != nil {
		return err
	}

	// Get role by default
	if user.RoleId == 0 {
		role := NewMysqlRoleRepository(m.DB)
		default_role, err := role.GetDefaultByPersonalAreaId(user.PersonalAreaId)
		if err != nil {
			return err
		}
		user.RoleId = default_role.Id
	}

	user.Password = entity.CompressPass(user.Password)

	return nil
}
