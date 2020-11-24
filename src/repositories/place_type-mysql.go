package repositories

import (
	"entity"
	"github.com/jinzhu/gorm"
)

type mysqlPlaceTypeRepo struct {
	DB *gorm.DB
}

type PlaceTypeRepository interface {
	Save(e entity.PlaceType) (entity.PlaceType, error)
}

func NewMysqlPlaceTypeRepository(db *gorm.DB) entity.PlaceTypeRepository {
	return &mysqlPlaceTypeRepo{
		DB: db,
	}
}

func (m *mysqlPlaceTypeRepo) Save(e entity.PlaceType) (entity.PlaceType, error) {
	m.DB.Create(&e)
	return e, nil
}
