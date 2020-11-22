package entity

type Device2Device struct {
	Id                   int `gorm:"primary_key"`
	DeviceDonorId        int
	ParameterDonorId     int
	DeviceRecipientId    int
	ParameterRecipientId int
	StartValue           int
	Weight               int
}
