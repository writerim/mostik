package entity

type Role struct {
	Id             int `gorm:"primary_key"`
	Title          string
	PersonalAreaId int
}
