package entity

type Pool struct {
	Id      int `gorm:"primary_key"`
	PlaceId int
	Uid     string
}
