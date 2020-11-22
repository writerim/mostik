package entity

type ModelParameter struct {
	Id          int `gorm:"primary_key"`
	ModelId     int
	ParameterId int
	RangeHours  int
	IsSystem    int
}
