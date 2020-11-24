package repositories

import (
	"entity"
	"github.com/jinzhu/gorm"
)

type mysqlPersonlAreaRepo struct {
	DB *gorm.DB
}

type PersonalAreaRepository interface {
	Save(e entity.PersonalArea) (entity.PersonalArea, error)
}

func NewMysqlPersonalAreaRepository(db *gorm.DB) entity.PersonalAreaRepository {
	return &mysqlPersonlAreaRepo{
		DB: db,
	}
}

func (m *mysqlPersonlAreaRepo) Save(e entity.PersonalArea) (entity.PersonalArea, error) {
	m.DB.Create(&e)
	return e, nil
}
