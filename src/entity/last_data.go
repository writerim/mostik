package entity

type LastData struct {
	Id          int `gorm:"primary_key"`
	DeviceId    int
	ParameterId int
	DataId      int
}
