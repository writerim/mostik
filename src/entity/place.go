package entity

type Place struct {
	Id             int `gorm:"primary_key"`
	Title          string
	PlaceId        int
	TypeId         int
	PayerId        int
	PersonalAreaId int
	IsDefault      int `gorm:"type:tinyint"`
	ApiKey         string
}
