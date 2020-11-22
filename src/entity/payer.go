package entity

type Payer struct {
	Id             int `gorm:"primary_key"`
	UserId         int
	PersonalAreaId int
	Fields         string `gorm:"type:json"`
}
