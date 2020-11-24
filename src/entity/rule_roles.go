package entity

func (RuleRoles) TableName() string {
	return "rule_roles"
}

type RuleRoles struct {
	Id     int `gorm:"primary_key"`
	RuleId int
	RoleId int
}

type RuleRoleRepository interface {
	Save(rule RuleRoles) (RuleRoles, error)
}
