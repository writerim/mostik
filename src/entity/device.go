package entity

type Device struct {
	Id             int `gorm:"primary_key"`
	Title          string
	PersonalAreaId int
	ModelId        int
	ParentId       int
	PlaceId        int
	Status         int
	Properties     string
	Location       string
}
