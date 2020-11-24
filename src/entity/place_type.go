package entity

func (PlaceType) TableName() string {
	return "place_type"
}

type PlaceType struct {
	Id    int `gorm:"primary_key"`
	Title string
	Ident string
}

type PlaceTypeRepository interface {
	Save(e PlaceType) (PlaceType, error)
}
