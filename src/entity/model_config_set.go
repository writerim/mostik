package entity

type ModelConfigSet struct {
	Id            int `gorm:"primary_key"`
	DeviceId      int
	ModelConfigId int
	Value         string
}
