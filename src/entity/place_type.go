package entity

type PlaceType struct {
	Id    int `gorm:"primary_key"`
	Title string
	Ident string
}
