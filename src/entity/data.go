package entity

type Data struct {
	Id          int `gorm:"primary_key"`
	DeviceId    int
	ParameterId int
	AddedTs     time.Time
	Value       int
	ValueTs     time.Time
}
