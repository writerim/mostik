package repositories

import (
	"entity"
	"github.com/jinzhu/gorm"
)

type mysqlTimeZoneRepo struct {
	DB *gorm.DB
}

type TimeZoneRepo interface {
	Save(rule entity.TimeZone) (entity.TimeZone, error)
}

func NewMysqlTimeZoneRepository(db *gorm.DB) entity.TimeZoneRepo {
	return &mysqlTimeZoneRepo{
		DB: db,
	}
}

func (tz *mysqlTimeZoneRepo) Save(t entity.TimeZone) (entity.TimeZone, error) {
	tz.DB.Create(&t)
	return t, nil
}
