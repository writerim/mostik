package entity

type UserLog struct {
	Id       int `gorm:"primary_key"`
	Action   string
	Object   string
	ObjectId int
	UserId   int
	AddTs    time.Time
}
