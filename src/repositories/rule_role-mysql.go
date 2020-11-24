package repositories

import (
	"entity"
	"github.com/jinzhu/gorm"
)

type mysqlRuleRoleRepo struct {
	DB *gorm.DB
}

type RuleRoleRepository interface {
	Save(rule entity.RuleRoles) (entity.RuleRoles, error)
}

func NewMysqlRuleRoleRepository(db *gorm.DB) entity.RuleRoleRepository {
	return &mysqlRuleRoleRepo{
		DB: db,
	}
}

func (m *mysqlRuleRoleRepo) Save(rr entity.RuleRoles) (entity.RuleRoles, error) {
	// select parse return
	return entity.RuleRoles{}, nil
}
