package entity

import (
	"time"
)

func (UserLog) TableName() string {
	return "user_log"
}

type UserLog struct {
	Id       int `gorm:"primary_key"`
	Action   string
	Object   string
	ObjectId int
	UserId   int
	AddTs    time.Time
}
