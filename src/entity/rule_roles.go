package entity

type RuleRoles struct {
	Id     int `gorm:"primary_key"`
	RuleId int
	RoleId int
}
