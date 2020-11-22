package entity

type PersonalArea struct {
	Id       int `gorm:"primary_key"`
	Title    string
	UserId   int
	ParentId int
	Location string
}
