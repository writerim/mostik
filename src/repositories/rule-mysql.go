package repositories

import (
	"entity"
	"github.com/jinzhu/gorm"
)

type mysqlRuleRepo struct {
	DB *gorm.DB
}

type RuleRepository interface {
	Save(rule entity.Rule) (entity.Rule, error)
	GetAll() ([]entity.Rule, error)
}

func NewMysqlRuleRepository(db *gorm.DB) entity.RuleRepository {
	return &mysqlRuleRepo{
		DB: db,
	}
}

func (m *mysqlRuleRepo) Save(rule entity.Rule) (entity.Rule, error) {
	m.DB.Create(&rule)
	return rule, nil
}

func (m *mysqlRuleRepo) GetAll() ([]entity.Rule, error) {
	return []entity.Rule{}, nil
}
