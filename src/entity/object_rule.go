package entity

type ObjectRule struct {
	Id       int `gorm:"primary_key"`
	Object   string
	ObjectId int
	UserId   int
	RoleId   int
}
